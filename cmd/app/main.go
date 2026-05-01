package main

import (
	"github.com/FaustaAkbar/BE_Golang/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.Hello)

	r.Run(":8080")
}
