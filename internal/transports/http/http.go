package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	defaultPort = 8000
	defaultHost = "0.0.0.0"
)

type Options struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Mode string `yaml:"mode"`
}

func NewOptions(v *viper.Viper, logger *log.Logger) (*Options, error) {
	options := new(Options)
	if err := v.UnmarshalKey("http", options); err != nil {
		return nil, errors.Wrap(err, "unmarshal http options error")
	}
	logger.Println("router options loaded successfully", options)
	return options, nil
}

type Server struct {
	o          *Options
	host       string
	port       int
	logger     *log.Logger
	router     *gin.Engine
	httpServer http.Server
}

func New(options *Options, logger *log.Logger, router *gin.Engine) (*Server, error) {
	logger.SetPrefix("[server] ")
	s := &Server{
		logger: logger,
		router: router,
		o:      options,
	}

	return s, nil
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port <= 0 {
		s.port = defaultPort
	}

	s.host = s.o.Host

	if s.host == "" {
		s.host = defaultHost
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Println("http server starting ...")
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("start http server err %s", err.Error())
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	s.logger.Println("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

var ProviderSet = wire.NewSet(New, NewRouter, NewOptions)
