package summer

import (
	"net/http"
	"strings"
)

// Core is the core of the summer framework.
type Core struct {
	router map[string]map[string]ControllerHandler
}

// NewCore creates a new Core
func NewCore() *Core {
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	router := map[string]map[string]ControllerHandler{
		"GET":    getRouter,
		"POST":   postRouter,
		"PUT":    putRouter,
		"DELETE": deleteRouter,
	}

	return &Core{
		router: router,
	}
}

func (core *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	core.router["GET"][upperUrl] = handler
}

func (core *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	core.router["POST"][upperUrl] = handler
}

func (core *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	core.router["PUT"][upperUrl] = handler
}

func (core *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	core.router["DELETE"][upperUrl] = handler
}

// ServeHTTP implements the http.Handler interface.
func (core *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r, w)

	router := core.findRouteByRequest(r)
	if router == nil {
		ctx.Json(404, "Not found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.Json(500, "Inner server error")
		return
	}
}

// findRouteByRequest finds the matched route by looping.
func (core *Core) findRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)

	if methodHandlers, ok := core.router[upperMethod]; ok {
		if handle, ok := methodHandlers[upperUri]; ok {
			return handle
		}
	}

	return nil
}
