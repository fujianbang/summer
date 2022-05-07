package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fujianbang/summer"
)

func main() {
	core := summer.NewCore()
	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}

func registerRouter(core *summer.Core) {
	core.Get("foo", FooControllerHandler)
}

func FooControllerHandler(ctx *summer.Context) error {
	log.Println("Got a request to /foo")
	// set context
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), 10*time.Second)
	defer cancel()

	// create new goroutine to handle business logic
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// TODO business logic
		time.Sleep(10 * time.Second)

		ctx.Json(http.StatusOK, map[string]interface{}{
			"code": 0,
		})

		// notice parent goroutine when finish
		finish <- struct{}{}
	}()

	// handle event
	select {
	case <-finish:
		log.Println("finish")
	case p := <-panicChan:
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()

		log.Println("panic", p)
	case <-durationCtx.Done():
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()

		ctx.Json(http.StatusInternalServerError, "timeout")
		ctx.SetHasTimeout()
	}
	return nil
}
