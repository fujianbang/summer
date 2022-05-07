package summer

import "net/http"

// Core is the core of the summer framework.
type Core struct {
	router map[string]ControllerHandler
}

// NewCore creates a new Core
func NewCore() *Core {
	return &Core{
		router: make(map[string]ControllerHandler),
	}
}

func (core *Core) Get(url string, handler ControllerHandler) {
	core.router[url] = handler
}

// ServeHTTP implements the http.Handler interface.
func (core *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r, w)

	router := core.router["foo"]
	if router == nil {
		return
	}

	router(ctx)
}
