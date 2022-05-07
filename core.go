package summer

import "net/http"

// Core is the core of the summer framework.
type Core struct {
}

// NewCore creates a new Core
func NewCore() *Core {
	return &Core{}
}

// ServeHTTP implements the http.Handler interface.
func (core *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO
}
