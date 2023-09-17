package run

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Paul1k96/bookstorage/internal/infrastructure/responder"
	bs "github.com/Paul1k96/bookstorage/internal/modules/books/service/book_grpc"

	"github.com/Paul1k96/bookstorage/config"
	"github.com/Paul1k96/bookstorage/internal/db"
	apierr "github.com/Paul1k96/bookstorage/internal/infrastructure/errors"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/router"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/server"
	"github.com/Paul1k96/bookstorage/internal/modules"
	"github.com/Paul1k96/bookstorage/pkg/grpc/books"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Application интерфейс приложения
type Application interface {
	Runner
	Bootstraper
}

// Runner интерфейс для запуска приложения
type Runner interface {
	Run() int
}

// Bootstraper интерфейс для инициализации приложения
type Bootstraper interface {
	Bootstrap() Runner
}

// App структура приложения
type App struct {
	conf     config.AppConf
	srv      server.Server
	gRPC     server.Server
	Sig      chan os.Signal
	Storages *modules.Storages
	Services *modules.Services
}

// NewApp конструктор приложения
func NewApp(conf config.AppConf) *App {
	return &App{conf: conf, Sig: make(chan os.Signal, 1)}
}

// Run запуск приложения
func (a *App) Run() int {
	// создание контекста для завершения приложения
	ctx, cancel := context.WithCancel(context.Background())

	// создание errorGroup по контексту для завершения всех запущеных горутин
	errGroup, ctx := errgroup.WithContext(ctx)

	// запуск горутины для graceful shutdown
	// при получении сигнала ОС вызывается завершение контекста
	errGroup.Go(func() error {
		sigInt := <-a.Sig
		log.Println("signal interrupt recieved", "os_signal", sigInt)
		cancel()
		return nil
	})

	// запуск http сервера приложения
	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("app: server error", err)
			return err
		}
		return nil
	})

	// запуск gRPC сервера приложения
	errGroup.Go(func() error {
		err := a.gRPC.Serve(ctx)
		if err != nil {
			log.Println("app: gRPC server error", err)
			return err
		}
		return err
	})

	// ожидание завершения горутин
	if err := errGroup.Wait(); err != nil {
		return apierr.GeneralError
	}

	return apierr.NoError
}

// Bootstrap сборка приложения
func (a *App) Bootstrap() Runner {
	// инициализация базы данных
	dbx, err := db.NewSqlDB(a.conf.DB)
	if err != nil {
		log.Fatal("error init db", err)
	}

	// инициализация менеджера ответа сервера
	newResponder := responder.NewResponder()

	// инициализация слоёв storage,service и controller приложения
	storages := modules.NewStorages(dbx)
	a.Storages = storages
	services := modules.NewServices(storages)
	a.Services = services
	controllers := modules.NewControllers(services, newResponder)

	// инициализация роутера
	var r *chi.Mux
	r = router.NewRouter(controllers)

	// инициализация gRPC сервера сервиса books
	gRPCServer := grpc.NewServer()

	gRPCBooks := bs.NewBooksServiceGRPC(services.Book)
	books.RegisterBookServiceGRPCServer(gRPCServer, gRPCBooks)

	a.gRPC = server.NewGRPC(a.conf.RPCServer, gRPCServer)

	// инициализация сервера
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: r,
	}
	a.srv = server.NewHttpServer(a.conf.Server, srv)
	// возвращение приложения
	return a
}
