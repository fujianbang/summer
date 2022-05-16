package main

import (
	"log"
	"net/http"

	"github.com/fujianbang/summer"
	"github.com/fujianbang/summer/middleware"
)

func main() {
	core := summer.NewCore()
	core.Use(middleware.Recovery())

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
	core.Get("/panic", func(c *summer.Context) error {
		log.Println("panic")
		panic("panic")
		return nil
	})
}
