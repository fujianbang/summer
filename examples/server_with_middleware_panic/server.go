package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fujianbang/summer"
	"github.com/fujianbang/summer/middleware"
)

func main() {
	core := summer.NewCore()
	core.Use(middleware.Duration())

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
	core.Get("/duration", func(c *summer.Context) error {
		log.Println("request /duration")
		time.Sleep(time.Millisecond * 123)
		return nil
	})
}
