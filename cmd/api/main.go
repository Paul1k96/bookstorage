package main

import (
	"log"
	"os"

	"github.com/Paul1k96/bookstorage/config"
	"github.com/Paul1k96/bookstorage/run"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	conf := config.NewAppConf()
	conf.Init()

	app := run.NewApp(conf)

	exitCode := app.Bootstrap().Run()

	os.Exit(exitCode)

}
