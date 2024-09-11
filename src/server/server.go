package server

import (
	"context"
	"crypto/tls"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"strconv"
	"teacherScheduler/src/config"
	"teacherScheduler/stub"
)

type Server struct {
	tlsCfg          *tls.Config
	port            int
	handler         http.Handler
	serveLoops      []func(context.Context) error
	runHealthzCheck bool
}

func Start() {
	conf := config.Load()

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	controller, err := NewController()
	if err != nil {
		e.Logger.Fatal(err)
	}

	baseUrl := "/v1/scheduler"

	stub.RegisterHandlersWithBaseURL(e, controller, baseUrl)

	s := NewServer(e, conf.Server.Port)
	if err := s.Run(); err != nil {
		logrus.Fatal("Failure occurred while running server")
	}

}

func NewServer(handler http.Handler, port int) *Server {
	s := &Server{
		handler:    handler,
		port:       port,
		serveLoops: make([]func(context.Context) error, 0),
	}

	return s
}

func (s *Server) Run() error {
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(s.serveHTTP(ctx))
	//for _, fn := range s.serveLoops {
	//	eg.Go(s.handleFunc(ctx, fn))
	//}
	return eg.Wait()
}

func (s *Server) serveHTTP(ctx context.Context) func() error {
	return func() error {
		eg, childCtx := errgroup.WithContext(ctx)
		server := http.Server{
			Addr:    ":" + strconv.Itoa(s.port),
			Handler: s.handler,
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
		}

		eg.Go(func() error {
			if s.tlsCfg != nil {
				server.TLSConfig = s.tlsCfg
				return server.ListenAndServeTLS("", "")
			} else {
				return server.ListenAndServe()
			}
		})

		eg.Go(func() error {
			<-childCtx.Done()
			return server.Shutdown(childCtx)
		})

		return eg.Wait()
	}
}
