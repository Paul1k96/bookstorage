package run

import (
	"context"
	"fmt"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/responder"
	"log"
	"net/http"
	"os"

	"github.com/Paul1k96/bookstorage/config"
	"github.com/Paul1k96/bookstorage/internal/db"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/errors"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/router"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/server"
	"github.com/Paul1k96/bookstorage/internal/modules"
	bs "github.com/Paul1k96/bookstorage/internal/modules/books/service"
	"github.com/Paul1k96/bookstorage/pkg/grpc/books"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Application interface {
	Runner
	Bootstraper
}

type Runner interface {
	Run() int
}

type Bootstraper interface {
	Bootstrap() Runner
}

type App struct {
	conf     config.AppConf
	srv      server.Server
	gRPC     server.Server
	Sig      chan os.Signal
	Storages *modules.Storages
	Services *modules.Services
}

func NewApp(conf config.AppConf) *App {
	return &App{conf: conf, Sig: make(chan os.Signal, 1)}
}

func (a *App) Run() int {
	ctx, cancel := context.WithCancel(context.Background())

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		sigInt := <-a.Sig
		log.Println("signal interrupt recieved", "os_signal", sigInt)
		cancel()
		return nil
	})

	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && err != http.ErrServerClosed {
			log.Println("app: server error", err)
			return err
		}
		return nil
	})

	errGroup.Go(func() error {
		err := a.gRPC.Serve(ctx)
		if err != nil {
			log.Println("app: gRPC server error", err)
			return err
		}
		return err
	})

	if err := errGroup.Wait(); err != nil {
		return errors.GeneralError
	}

	return errors.NoError
}
func (a *App) Bootstrap() Runner {
	dbx, err := db.NewSqlDB(a.conf.DB)
	if err != nil {
		log.Fatal("error init db", err)
	}

	newResponder := responder.NewResponder()

	storages := modules.NewStorages(dbx)
	a.Storages = storages
	services := modules.NewServices(storages)
	a.Services = services

	controllers := modules.NewControllers(services, newResponder)

	var r *chi.Mux
	r = router.NewRouter(controllers)

	gRPCServer := grpc.NewServer()

	gRPCBooks := bs.NewBooksServiceGRPC(services.Book)
	books.RegisterBookServiceGRPCServer(gRPCServer, gRPCBooks)

	a.gRPC = server.NewGRPC(a.conf.RPCServer, gRPCServer)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: r,
	}
	a.srv = server.NewHttpServer(a.conf.Server, srv)
	return a
}
