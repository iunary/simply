package prom

import (
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricsPath = "/metrics"
	faviconPath = "/favicon.ico"
)

var httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "http_server",
	Name:      "requests_seconds",
	Help:      "Histogram of response latency (seconds) of http handlers.",
}, []string{"method", "code", "uri"})

func init() {
	prometheus.MustRegister(httpHistogram)
}

type handlerPath struct {
	sync.Map
}

func (hp *handlerPath) get(handler string) string {
	v, ok := hp.Load(handler)
	if !ok {
		return ""
	}
	return v.(string)
}

func (hp *handlerPath) set(ri gin.RouteInfo) {
	hp.Store(ri.Handler, ri.Path)
}

// GinPrometheus  struct
type Prometheus struct {
	engine  *gin.Engine
	ignored map[string]bool
	pathMap *handlerPath
	updated bool
}

// Option
type Option func(*Prometheus)

// Ignore
func Ignore(path ...string) Option {
	return func(gp *Prometheus) {
		for _, p := range path {
			gp.ignored[p] = true
		}
	}
}

// New
func New(e *gin.Engine, options ...Option) *Prometheus {
	if e == nil {
		return nil
	}
	gp := &Prometheus{
		engine: e,
		ignored: map[string]bool{
			metricsPath: true,
			faviconPath: true,
		},
		pathMap: &handlerPath{},
	}
	for _, o := range options {
		o(gp)
	}
	return gp
}

func (gp *Prometheus) updatePath() {
	gp.updated = true
	for _, ri := range gp.engine.Routes() {
		gp.pathMap.set(ri)
	}
}

// Middleware
func (gp *Prometheus) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !gp.updated {
			gp.updatePath()
		}

		if gp.ignored[c.Request.URL.String()] {
			c.Next()
			return
		}
		start := time.Now()

		c.Next()

		httpHistogram.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
			gp.pathMap.get(c.HandlerName()),
		).Observe(time.Since(start).Seconds())
	}
}
