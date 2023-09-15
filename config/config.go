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

type AppConf struct {
	AppName   string
	Server    Server
	DB        DB
	RPCServer RPCServer
}

type RPCServer struct {
	Port         string
	ShutdownTime time.Duration
}

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

type Server struct {
	Port            string
	ShutdownTimeout time.Duration
}

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

func (a *AppConf) Init() {
	shutdownTimeOut, err := strconv.Atoi(os.Getenv(envShutdownTimeout))
	if err != nil {
		log.Fatal(parseShutdownTimeoutError)
	}
	shutdownTimeout := time.Duration(shutdownTimeOut) * time.Second
	if err != nil {
		log.Fatal(parseRpcShutdownTimeoutError)
	}

	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		log.Fatal(parseDbTimeoutErr)
	}
	dbMaxConn, err := strconv.Atoi(os.Getenv("MAX_CONN"))
	if err != nil {
		log.Fatal(parseDbMaxConnErr)
	}
	a.DB.Timeout = dbTimeout
	a.DB.MaxConn = dbMaxConn

	a.RPCServer.Port = os.Getenv("RPC_PORT")

	a.Server.ShutdownTimeout = shutdownTimeout
}
