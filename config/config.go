package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	serverPort         = "PORT"
	appName            = "APP_NAME"
	envShutdownTimeout = "SHUTDOWN_TIMEOUT"

	parseDbTimeoutErr            = "config: parse db timeout err"
	parseDbMaxConnErr            = "config: parse db max connection err"
	parseShutdownTimeoutError    = "config: parse server shutdown timeout error"
	parseRpcShutdownTimeoutError = "config: parse rpc server shutdown timeout error"
)

// AppConf структура описания конфигурации приложения
type AppConf struct {
	AppName   string
	Server    Server
	DB        DB
	RPCServer RPCServer
}

// RPCServer структура конфигурации RPC сервера
type RPCServer struct {
	Port         string
	ShutdownTime time.Duration
}

// DB структура конфигурации базы данных
type DB struct {
	Driver   string
	Net      string
	Name     string
	User     string
	Password string
	Host     string
	MaxConn  int
	Port     string
	Timeout  int
}

// Server структура сервера приложения
type Server struct {
	Port            string
	ShutdownTimeout time.Duration
}

// NewAppConf конструктор конфигурации приложения
func NewAppConf() AppConf {
	return AppConf{
		AppName: os.Getenv(appName),
		Server: Server{
			Port: os.Getenv(serverPort),
		},
		DB: DB{
			Net:      os.Getenv("DB_NET"),
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}

// Init инициализация конфигурации приложения
func (a *AppConf) Init() {
	// получение таймаута сервера из .env
	shutdownTimeOut, err := strconv.Atoi(os.Getenv(envShutdownTimeout))
	if err != nil {
		log.Fatal(parseShutdownTimeoutError)
	}
	shutdownTimeout := time.Duration(shutdownTimeOut) * time.Second
	if err != nil {
		log.Fatal(parseRpcShutdownTimeoutError)
	}

	// получение таймаута базы данных из .env
	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		log.Fatal(parseDbTimeoutErr)
	}

	// получение ограничения на подключение к базе данных
	dbMaxConn, err := strconv.Atoi(os.Getenv("MAX_CONN"))
	if err != nil {
		log.Fatal(parseDbMaxConnErr)
	}
	a.DB.Timeout = dbTimeout
	a.DB.MaxConn = dbMaxConn

	a.RPCServer.Port = os.Getenv("RPC_PORT")

	a.Server.ShutdownTimeout = shutdownTimeout
}
