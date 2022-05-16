package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fujianbang/summer"
)

func main() {
	core := summer.NewCore()
	core.Use(TestMiddleware1(), TestMiddleware2())

	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func registerRouter(core *summer.Core) {

	core.Get("/subject", func(c *summer.Context) error {
		log.Println("do nothing")
		return nil
	})
}

func TestMiddleware1() summer.ControllerHandler {
	return func(c *summer.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}
func TestMiddleware2() summer.ControllerHandler {
	return func(c *summer.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}
