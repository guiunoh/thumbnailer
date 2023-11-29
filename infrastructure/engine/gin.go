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

func NewGin() Engine {
	engine := gin.Default()
	engine.Use(requestid.New())
	engine.MaxMultipartMemory = 4 << 20 // 4 MiB
	return &ginEngine{
		app:    engine,
		server: nil,
	}
}

type ginEngine struct {
	app    *gin.Engine
	server *http.Server
}

// Group implements Engine.
func (e *ginEngine) Group(path string) gin.IRouter {
	return e.app.Group(path)
}

// Monitor implements Engine.
func (e *ginEngine) Monitor(port int, path string) {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath(path)
	metrics.UseWithoutExposingEndpoint(e.app)
	metrics.Expose(router)

	serve(router, port)
}

// Ping implements Engine.
func (e *ginEngine) Ping(path string) {
	e.app.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet && strings.HasSuffix(c.Request.URL.Path, path) {
			c.String(http.StatusOK, "pong")
			c.Abort()
			return
		}
		c.Next()
	})
}

// Serve implements Engine.
func (e *ginEngine) Serve(port int) {
	e.server = serve(e.app, port)
}

// Shutdown implements Engine.
func (e *ginEngine) Shutdown(timeout time.Duration) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	e.server.Shutdown(ctx)
}

func serve(engine *gin.Engine, port int) *http.Server {
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
