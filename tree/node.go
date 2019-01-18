package tree

import	(
	"os"
	"fmt"
	"sync"
	"time"
	"strings"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/security"
)

func NewNode(config *Config) *Node {
	return &Node{
		Config: config,
		Headers: map[string]interface{}{},
		Routes:	map[string]*Node{},
		Methods: map[string]*Handler{},
		RequestParams: map[string]interface{}{},
		Modules: []*Module{},
		Validations: []*validation.Config{},
	}
}

type Node struct {
	Config *Config
	Parent *Node
	Path string
	Parameter *Node
	Headers map[string]interface{}
	RequestParams map[string]interface{}
	Routes map[string]*Node
	Methods map[string]*Handler
	Module *Module
	Modules []*Module
	SecurityModule security.SecurityModule
	Validation *validation.Config
	Validations []*validation.Config
	spec interface{}
	sync.RWMutex
}

func (node *Node) new(path string) *Node {
	n := NewNode(node.Config)
	n.Parent = node
	n.Modules = node.Modules
	n.Path = path
	n.SecurityModule = node.SecurityModule
	n.Validations = node.Validations
	// create a new map inheriting the values from the parant node
	for k, v := range node.Headers {
		n.Headers[k] = v
	}
	return n
}

func (node *Node) SetHeaders(headers map[string]interface{}) *Node {
	node.Headers = headers
	return node
}

// Returns the node's full path string
func (node *Node) FullPath() string {
	var path string
	if node.Parent != nil {
		path = fmt.Sprintf("%s/%s", node.Parent.FullPath(), node.Path)
	}
	return path
}

// Adds a new node to the tree
func (node *Node) Add(path string, pathKeys ...string) *Node {

	path = strings.TrimSpace(
		strings.Replace(path, "/", "", -1),
	)

	node.RLock()
		existing := node.Routes[path]
	node.RUnlock()
	if existing != nil {
		return existing
	}

	n := node.new(path)
	node.Lock()
		node.Routes[path] = n
	node.Unlock()

	if len(pathKeys) > 0 {
		n.Lock()
			for _, key := range pathKeys {
				n.RequestParams[key] = path
			}
		n.Unlock()
	}

	return n
}

// Adds a new param-node
func (node *Node) Param(vc *validation.Config, keys ...string) *Node {

	if len(keys) == 0 { panic("NO KEYS SUPPLIED FOR NEW PARAMETER") }

	node.RLock()
		p := node.Parameter
	node.RUnlock()
	if p != nil {
		return p
	}

	n := node.new("{" + keys[0] + "}")

	vc.Keys = keys

	n.Lock()
		n.Validation = vc
		n.Validations = append(n.Validations, vc)
	n.Unlock()

	node.Lock()
		node.Parameter = n
	node.Unlock()

	return n
}

/*
// Adds a new param-node
func (node *Node) AuthBearer() *Node {

	node.security = &security.Auth_HTTP{
		Scheme: "bearer",
		BearerFormat: "JWT",
	}

	return node
}
*/
func (node *Node) newModule(function ModuleFunction, arg interface{}) *Module {

	//if node.Config.ModuleRegistry == nil { panic("Config has no ModuleRegistry setting!") }

	return &Module{
		config:					node.Config,
		function:				function,
		arg:					arg,
	}
}

// Adds a module that will be executed at the point it is added to the route
func (node *Node) Init(function ModuleFunction, arg interface{}) *Node {
	module := node.newModule(function, arg)
	node.Lock()
		node.Module = module
	node.Unlock()
	return node
}

// Adds a module that will be executed upon reaching a handler
func (node *Node) Mod(function ModuleFunction, arg interface{}) *Node {

	if function == nil { panic("INVALID MODULE FUNC") }

	module := node.newModule(function, arg)

	node.Lock()
		node.Modules = append(node.Modules, module)
	node.Unlock()

	return node
}

// execute init function added with .Init(...)
func (node *Node) RunModule(req http.Request) *http.Status {

	node.RLock()
		module := node.Module
	node.RUnlock()

	if module != nil {

		status := module.Run(req); if status != nil { return status }
	}

	return nil
}

// execute all module functions added with .Mod(...)
func (node *Node) RunModules(req http.Request) *http.Status {

	for _, module := range node.Modules {
		status := module.Run(req)
		if status != nil {
			return status
		}
	}

	return nil
}

// traversal

// finds next node according to supplied URL path segment
func (node *Node) Next(req http.Request, pathSegment string) (*Node, *http.Status) {

	// execute any init module(s)

	if status := node.RunModule(req); status != nil {
		return nil, status
	}

	// check for child routes

	next := node.Routes[pathSegment]

	if next != nil { return next, nil }

	// check for path param

	next = node.Parameter
	if next == nil {
		return nil, nil
	}

	if next.Validation != nil {

		status, value := next.Validation.PathFunction(req, pathSegment)
		if status != nil {

			status.Message = fmt.Sprintf("%s KEY(%v)", status.MessageString(), pathSegment)

			//return nil, &http.Status{nil, 400, fmt.Sprintf("UNEXPECTED VALUE  %v, %v", pathSegment, next.Validation.Expecting())}
			return nil, status
		}

		// write route params into request object

		for _, key := range next.Validation.Keys { req.SetParam(key, value) }

	}

	return next, nil
}

// Returns the handler assciated with the HTTP request method.
func (node *Node) Handler(req http.Request) *Handler {
	node.RLock()
	defer node.RUnlock()
	return node.Methods[req.Method()]
}

// Adds a file to be served from the specified path.
func (node *Node) File(path string) *Node {
	node.addHandler(
		"GET",
		&Handler{
			File: &File{
				Path: path,
			},
		},
	)
	return node
}

// Walks through the specified folder to mirror the file structure for files containing all filters
func (node *Node) Folder(directoryPath string, filters ...string) *Node {

	// remove trailing slash from the directory path if existing
	directoryPath = strings.TrimSuffix(directoryPath, "/")

	go func () {

		for {
			f, err := os.Open(directoryPath); if err != nil { panic(err) }

			names, err := f.Readdirnames(-1)
			f.Close()
			if err != nil { panic(err) }

			for _, name := range names {

				path := strings.Replace(directoryPath + "/" + name, "//", "/", -1)

				node.checkFile(name, path, filters)

			}

			time.Sleep(5 * time.Second)
		}

	}()

	return node
}

// Walks through the specified folder to mirror the file structure for files containing all filters
func (node *Node) StaticFolder(directoryPath string, filters ...string) *Node {

	// remove trailing slash from the directory path if existing
	directoryPath = strings.TrimSuffix(directoryPath, "/")

	f, err := os.Open(directoryPath); if err != nil { panic(err) }

	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil { panic(err) }

	for _, name := range names {

		path := strings.Replace(directoryPath + "/" + name, "//", "/", -1)

		node.checkFile(name, path, filters)

	}

	return node
}

// Checks if file or folder, adding any files
func (node *Node) checkFile(name, path string, filters []string) {

	info, err := os.Lstat(path)
	if err != nil {
		panic(err)
	}
	// check if the file is a directory
	if info.Mode() & os.ModeSymlink != 0 {
		linked, err := os.Readlink(path)
		if err != nil {
			panic(err)
		}
		s := strings.Split(path, "/")
		path = strings.Join(
			append(
				s[:len(s)-1],
				linked,
			),
			"/",
		)
		info, err = os.Lstat(path)
		if err != nil {
			panic(err)
		}
	}
	// check if path is a directory
	if info.IsDir() {
		node.Add(name).StaticFolder(path)
		return
	}

	for _, filter := range filters {
		if !strings.Contains(name, filter) { return }
	}

	node.Add(name).File(path)

}
