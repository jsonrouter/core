package core

import (
	"strings"
	"strconv"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
	//"github.com/jsonrouter/core/metrics"
	//"fmt"
)

const	(
	ROBOTS_TXT = "User-agent: *\nDisallow: /api/"
)

type Router interface{
  Serve(int)
}

type Headers map[string]string


// MainHandler is the main handler function for incoming requests
func MainHandler(req http.Request, node *tree.Node, fullPath string) (status *http.Status) {

	// enforce https-only if required
	if node.Config.ForcedTLS {
		if !req.IsTLS() {
			status = req.Respond(502, "PLEASE UPGRADE TO HTTPS")
			return
		}
	}

	if node.Config.UseMetrics {
		met := node.Config.Metrics
		met.Timers["requestTime"].Start()
		defer func() {

			if status == nil {
				status = req.Respond(200, "OK")
			} else {
				status.Respond(req)
			}

			met.MultiCounters["requestMethods"].Increment(req.Method())
			met.MultiCounters["requestMethods"].Update(met.SetResults)

			met.MultiCounters["responseCodes"].Increment(strconv.Itoa(status.Code))
			met.MultiCounters["responseCodes"].Update(met.SetResults)

			met.Counters["requestCount"].Increment()
			met.Counters["requestCount"].Update(met.SetResults)

			met.Timers["requestTime"].Stop()
			met.Timers["requestTime"].Update(met.SetResults)
		}()
	} else {
		defer func(){

			if status == nil {
				status = req.Respond(200, "OK")
			} else {
				status.Respond(req)
			}

		}()
	}

	switch fullPath {

		case "/robots.txt":
			status = req.Respond([]byte(ROBOTS_TXT))
			return

	}

	segments := strings.Split(fullPath, "/")[1:]

	next := node

	for _, segment := range segments {

		if len(segment) == 0 { break }

		var n *tree.Node
		n, status = next.Next(req, segment)
		if status != nil {
			return
		}

		if n != nil {
			/*
			for k, v := range n.RequestParams {
				req.SetParam(k, v)
			}
			*/
			next = n
			continue
		}
		req.Respond()
		//req.HttpError("NO ROUTE FOUND AT " + next.FullPath() + "/" + segment, 404)
		status = req.Respond(404, "NO ROUTE FOUND")
		return
	}

	// resolve handler
	handler := next.Handler(req)
	if handler == nil {
		//req.HttpError("NO CONTROLLER FOUND AT " + next.FullPath(), 500)
		status = req.Respond(404, "NO CONTROLLER FOUND")
		return
	}

	// set CORS headers
	for k, v := range handler.Headers {
		req.SetResponseHeader(k, v.(string))
	}

	// return if preflight request
	if req.Method() == "OPTIONS" {
		status = req.Respond(200, "OK")
		return
	}

	// read the request body and unmarshal into specified schema
	status = handler.ReadPayload(req)
	if status != nil {
		return
	}

	// execute modules
	status = handler.Node.RunModules(req)
	if status != nil {
		return
	}

	if handler.File != nil {

		status = handler.DetectContentType(req, handler.File.Path)
		if status != nil {
			return
		}

		req.SetResponseHeader("Content-Type", handler.File.MimeType)

		status = req.Respond(handler.File.Cache)
		return
	}

	if handler.Function == nil {
		req.Log().Panic("FAILED TO GET FUNCTION TO SERVE HANDLE OPERATION!")
	}

	// execute the handler
	status = handler.Function(req)
	return
}
