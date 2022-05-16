package middleware

import (
	"log"
	"time"

	"github.com/fujianbang/summer"
)

func Duration() summer.ControllerHandler {
	return func(c *summer.Context) error {
		startT := time.Now()
		c.Next()

		duration := time.Since(startT)

		log.Printf("request [%s]%s Duration: %s\n", c.GetRequest().Method, c.GetRequest().URL.Path, duration)
		return nil
	}
}
