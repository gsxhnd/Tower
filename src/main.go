package main

import "github.com/gsxhnd/tower/src/wire"

func main() {
	app := wire.InitApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
