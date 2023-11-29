package engine

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Engine interface {
	Serve(port int)
	Shutdown(timeout time.Duration)
	Ping(path string)
	Monitor(port int, path string)
	Group(path string) gin.IRouter
}
