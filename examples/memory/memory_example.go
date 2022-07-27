package main

import (
	cache "github.com/hare85/gin-cache"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hare85/gin-cache/persist"
)

func main() {
	app := gin.New()

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	app.GET("/hello",
		cache.CacheByRequestURI(memoryStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
