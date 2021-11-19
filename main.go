package main

import (
	"pollserver/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/poll/polls", func(c *gin.Context) {
		handlers.GetPolls(c)
	})

	r.POST("/api/v1/polls/add", func(c *gin.Context) {
		handlers.CreatePoll(c)
	})

	r.GET("/:id", func(c *gin.Context) {
		handlers.Find(c)
	})

	r.Run()
}
