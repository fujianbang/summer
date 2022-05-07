package main

import (
	"net/http"

	"github.com/fujianbang/summer"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: summer.NewCore(),
	}

	server.ListenAndServe()
}
