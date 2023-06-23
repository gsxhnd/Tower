package main

import (
	"github.com/gsxhnd/go-api-template/src/di"
)

func main() {
	task, err := di.InitTask()
	if err != nil {
		panic(err)
	}
	if err := task.Run(); err != nil {
		panic(err)
	}
}
