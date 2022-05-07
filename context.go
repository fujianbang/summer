package summer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Context is the custom context of `summer` framework.
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
	}
}

func (ctx *Context) WriteMux() {
	// TODO
}

func (ctx *Context) GetRequest() {
	// TODO
}

func (ctx *Context) GetResponse() {
	// TODO
}

func (ctx *Context) SetHasTimeout() {
	// TODO
	panic("implement me")
}

func (ctx *Context) HasTimeout() {
	// TODO
	panic("implement me")
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() {
	// TODO
	panic("implement me")
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	// TODO
	panic("implement me")
}

func (ctx *Context) Value(key interface{}) {
	// TODO
	panic("implement me")
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
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// write http response header
	ctx.responseWriter.WriteHeader(status)

	// write http response body
	_, err = ctx.responseWriter.Write(bytes)
	if err != nil {
		return err
	}

	// TODO

	return nil
}

func (ctx *Context) HTML(status int, data interface{}) error {
	// TODO
	panic("implement me")
}

func (ctx *Context) Text(status int, data string) error {
	// TODO
	panic("implement me")
}
