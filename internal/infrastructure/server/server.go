package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Paul1k96/bookstorage/config"
)

type Server interface {
	Serve(ctx context.Context) error
}

type HttpServer struct {
	conf config.Server
	srv  *http.Server
}

func NewHttpServer(conf config.Server, server *http.Server) Server {
	return &HttpServer{conf: conf, srv: server}
}

func (s *HttpServer) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)
	go func() {
		log.Printf("server started, port %s", s.conf.Port)
		if err = s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("http listen and serve error %v", err)
			chErr <- err
		}
	}()

	select {
	case <-chErr:
		return err
	case <-ctx.Done():
	}

	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.conf.ShutdownTimeout*time.Second)
	defer cancel()
	err = s.srv.Shutdown(ctxShutdown)

	return err
}
