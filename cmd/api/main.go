package main

import (
	"log"
	"os"

	"github.com/Paul1k96/bookstorage/config"
	"github.com/Paul1k96/bookstorage/run"
	"github.com/joho/godotenv"
)

func main() {
	// инициализация godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// создание конфигурации приложения
	conf := config.NewAppConf()
	// сборка и инициализация конфигурации приложения
	conf.Init()

	// создание приложения
	app := run.NewApp(conf)

	// сборка и зпуск приложения
	exitCode := app.Bootstrap().Run()

	os.Exit(exitCode)
}
