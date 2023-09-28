package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() http.Handler {
	r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	apiV1 := r.Group("/api/v1")
	NewPublicRoute(apiV1.Group("/public"))

	return r
}
