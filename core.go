package summer

import (
	"log"
	"net/http"
	"strings"
)

// Core is the core of the summer framework.
type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

// NewCore creates a new Core
func NewCore() *Core {
	router := map[string]*Tree{
		"GET":    NewTree(),
		"POST":   NewTree(),
		"PUT":    NewTree(),
		"DELETE": NewTree(),
	}

	return &Core{
		router: router,
	}
}

// Use register middleware
func (core *Core) Use(middleware ...ControllerHandler) {
	core.middlewares = append(core.middlewares, middleware...)
}

func (core *Core) Get(url string, handler ControllerHandler) {
	allHandlers := append(core.middlewares, handler)
	if err := core.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatalln("add router error: ", err)
	}
}

func (core *Core) Post(url string, handler ControllerHandler) {
	allHandlers := append(core.middlewares, handler)
	if err := core.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatalln("add router error: ", err)
	}
}

func (core *Core) Put(url string, handler ControllerHandler) {
	allHandlers := append(core.middlewares, handler)
	if err := core.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatalln("add router error: ", err)
	}
}

func (core *Core) Delete(url string, handler ControllerHandler) {
	allHandlers := append(core.middlewares, handler)
	if err := core.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatalln("add router error: ", err)
	}
}

func (core *Core) Group(prefix string) IGroup {
	return NewGroup(core, prefix)
}

// findRouteByRequest finds the matched route by looping.
func (core *Core) findRouteByRequest(request *http.Request) []ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := core.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}

	return nil
}

// ServeHTTP implements the http.Handler interface.
func (core *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r, w)

	// find router
	handlers := core.findRouteByRequest(r)
	if handlers == nil {
		// not find then printing log
		ctx.Json(404, "Not found")
		return
	}

	// set handlers
	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(500, "Inner server error")
		return
	}
}
