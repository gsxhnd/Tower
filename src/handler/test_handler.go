package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestHandler interface {
	HelloHandler(ctx *gin.Context)
}

type testHandle struct{}

func NewTestHandle() TestHandler {
	return &testHandle{}
}

func (h *testHandle) HelloHandler(ctx *gin.Context) {
	fmt.Println("hello")
}
