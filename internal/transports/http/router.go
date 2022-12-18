package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iunary/simply/internal/transports/http/middlewares/prom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handlers func(r *gin.Engine)

func NewRouter(o *Options, logger *log.Logger, setup Handlers) *gin.Engine {
	gin.SetMode(o.Mode)
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.Use(prom.New(r).Middleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// setup handlers
	setup(r)
	return r
}
