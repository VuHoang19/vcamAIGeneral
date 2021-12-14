package routes

import "github.com/gin-gonic/gin"

type Option func(*gin.RouterGroup)

var options = []Option{}

//Register the route configuration of the app
func Include(opts ...Option) {
	options = append(options, opts...)
}

//Initialize
func Init() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/api")
	for _, opt := range options {
		opt(v1)
	}
	return r
}
