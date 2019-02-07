package core

import (
	"strings"
	"strconv"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
	//"github.com/chrysmore/metrics"
)

const	(
	ROBOTS_TXT = "User-agent: *\nDisallow: /api/"
)

type Router interface{
  Serve(int)
}

type Headers map[string]string


// main handler
func MainHandler(req http.Request, node *tree.Node, fullPath string) (status *http.Status) {
	met := node.Config.Metrics
	node.Config.Log.Debug("Hello Request")

	met.Counters["requestCount"].Increment()
	met.Counters["requestCount"].Update(&node.Config.MetResults)

	met.Timers["requestTime"].Start()	
	defer met.Timers["requestTime"].Update(&node.Config.MetResults)
	defer met.Timers["requestTime"].Stop()

	status = &http.Status{}
	defer met.MultiCounters["responseCodes"].Update(&node.Config.MetResults)
	
	defer func(){
		met.MultiCounters["responseCodes"].Increment(strconv.Itoa(status.Code))
	}()

	// enforce https-only if required

	if node.Config.ForcedTLS {
		if !req.IsTLS() {
			status = req.Respond(502, "PLEASE UPGRADE TO HTTPS")
			status.Respond(req)
			return
		}
	}

	switch fullPath {

		case "/robots.txt":

			req.Write([]byte(ROBOTS_TXT))
			return

	}

	segments := strings.Split(fullPath, "/")[1:]

	next := node

	for _, segment := range segments {

		if len(segment) == 0 { break }

		var n *tree.Node
		n, status = next.Next(req, segment)
		if status != nil {
			status.Respond(req)
			return
		}

		if n != nil {
			for k, v := range n.RequestParams {
				req.SetParam(k, v)
			}
			next = n
			continue
		}

		req.HttpError("NO ROUTE FOUND AT " + next.FullPath() + "/" + segment, 404)
		return
	}

	// resolve handler
	handler := next.Handler(req)
	if handler == nil {
		req.HttpError("NO CONTROLLER FOUND AT " + next.FullPath(), 500)
		return
	}
/*
	// set CORS headers
	for k, v := range handler.Headers {
		req.SetHeader(k, v)
	}
*/
	// return if preflight request
	if req.Method() == "OPTIONS" { return }

	// read the request body and unmarshal into specified schema
	status = handler.ReadPayload(req)
	if status != nil {
		status.Respond(req)
		return
	}

	// execute modules
	status = handler.Node.RunModules(req)
	if status != nil {
		status.Respond(req)
		return
	}

	if handler.File != nil {

		status = handler.DetectContentType(req, handler.File.Path)
		if status != nil {
			status.Respond(req)
			return
		}

		req.SetResponseHeader("Content-Type", handler.File.MimeType)

		status = req.Respond(handler.File.Cache)
		status.Respond(req)
		return
	}

	if handler.Function == nil {
		req.Log().Panic("FAILED TO GET FUNCTION TO SERVE HANDLE OPERATION!")
	}

	// execute the handler
	status = handler.Function(req)
	status.Respond(
		req,
	)


	
	return
}
