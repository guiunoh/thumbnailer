package engine

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func NewGin() *gin.Engine {
	engine := gin.Default()
	engine.Use(requestid.New())
	engine.MaxMultipartMemory = 4 << 20 // 4 MiB
	return engine
}

func Serve(engine *gin.Engine, port int) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}

	go func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}(srv)
	return srv
}

func Shutdown(server *http.Server, timeout time.Duration) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	server.Shutdown(ctx)
}

func Ping(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet && strings.HasSuffix(c.Request.URL.Path, path) {
			c.String(http.StatusOK, "pong")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Monitor(r gin.IRouter, port int, path string) {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath(path)
	metrics.UseWithoutExposingEndpoint(r)
	metrics.Expose(router)

	Serve(router, port)
}
