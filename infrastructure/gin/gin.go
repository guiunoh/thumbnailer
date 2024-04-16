package _gin

import (
	"context"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func Serve(r *gin.Engine, addr string) *http.Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return srv
}

func Shutdown(server *http.Server) {
	const timeout = 5 * time.Second

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
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

func Monitor(r *gin.Engine, addr string) {
	p := ginprometheus.NewPrometheus("gin")
	p.SetListenAddress(addr)
	p.Use(r)
}
