package main

import (
	"flag"

	"github.com/gsxhnd/go-api-template/src/di"
)

func main() {
	filePath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	app, err := di.InitApp(filePath)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
