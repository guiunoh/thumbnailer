package main

import (
	"thumbnailer/cmd/thumbnailer/bootstrap"
	_gin "thumbnailer/infrastructure/gin"
	_redis "thumbnailer/infrastructure/redis"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.MaxMultipartMemory = 4 << 20 // 4 MiB

	rdb := _redis.NewClient(bootstrap.Config.RDB)
	defer rdb.Close()

	engine.Use(requestid.New())
	engine.Use(_gin.MiddlewarePing("/ping"))

	bootstrap.Setup(engine, rdb)
	engine.Run()
}
