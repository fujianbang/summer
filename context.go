package summer

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// Context is the custom context of `summer` framework.
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	writeMux       *sync.Mutex
	hasTimeout     bool                // timeout flag
	handlers       []ControllerHandler // handlers
	index          int                 // current handler index
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		writeMux:       &sync.Mutex{},
		hasTimeout:     false,
		handlers:       make([]ControllerHandler, 0),
		index:          -1,
	}
}

func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writeMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.request.Context().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) QueryInt(key string, def int) {
	// TODO
	panic("implement me")
}

func (ctx *Context) QueryString(key string, def string) {
	// TODO
	panic("implement me")
}

func (ctx *Context) QueryArray(key string, def bool) {
	// TODO
	panic("implement me")
}

func (ctx *Context) QueryAll() {
	// TODO
	panic("implement me")
}

func (ctx *Context) FromInt(key string, def int) {
	// TODO
	panic("implement me")
}

func (ctx *Context) FromString(key string, def string) {
	// TODO
	panic("implement me")
}

func (ctx *Context) FromArray(key string, def bool) {
	// TODO
	panic("implement me")
}

func (ctx *Context) FromAll() {
	// TODO
	panic("implement me")
}

func (ctx *Context) BindJson(obj interface{}) {
	// TODO
	panic("implement me")
}

func (ctx *Context) Json(status int, data interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}

	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = ctx.responseWriter.Write(bytes)
	if err != nil {
		return err
	}

	// TODO

	return nil
}

func (ctx *Context) HTML(status int, data interface{}, template string) error {
	// TODO
	panic("implement me")
}

func (ctx *Context) Text(status int, data string) error {
	// TODO
	panic("implement me")
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// Next calls the next handler in the chain.
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}
