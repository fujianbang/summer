package middleware

import (
	"log"

	"github.com/fujianbang/summer"
)

func Recovery() summer.ControllerHandler {
	return func(c *summer.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
			log.Println("[Recovery] recover execute")
		}()
		c.Next()

		return nil
	}
}
