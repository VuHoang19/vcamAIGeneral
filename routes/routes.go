package routes

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

//Register the route configuration of the app
func Include(opts ...Option) {
	options = append(options, opts...)
}

//Initialize
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
