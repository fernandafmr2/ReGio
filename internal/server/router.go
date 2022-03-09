package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	// create default gin router
	router := gin.Default()

	// Enables automatic redirection if the current route can't be matched but a
  	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// create API route
	api := router.Group("/api")
	{
		api.POST("/signup", signUp)
		api.POST("/signin", signIn)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
