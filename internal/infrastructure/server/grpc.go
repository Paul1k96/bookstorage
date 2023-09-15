package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Paul1k96/bookstorage/config"
	"google.golang.org/grpc"
)

type ServerGRPC struct {
	conf config.RPCServer
	srv  *grpc.Server
}

func NewGRPC(conf config.RPCServer, srv *grpc.Server) Server {
	return &ServerGRPC{conf: conf, srv: srv}
}

func (s *ServerGRPC) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)
	go func() {
		var lis net.Listener
		lis, err = net.Listen("tcp", fmt.Sprintf(":%s", s.conf.Port))
		if err != nil {
			log.Println("grpc server register error: ", err)
			chErr <- err
		}

		log.Println("grpc server started", s.conf.Port)

		if err = s.srv.Serve(lis); err != nil {
			chErr <- err
		}
	}()

	select {
	case <-chErr:
		return err
	case <-ctx.Done():
		s.srv.GracefulStop()
	}

	return err
}
