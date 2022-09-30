package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.GET("/sleep/:time", func(context *gin.Context) {
		t := context.Param("time")
		t1, _ := strconv.Atoi(t)
		time.Sleep(time.Duration(t1) * time.Second)
		context.String(http.StatusOK, "sleep "+t)
	})
	server.Run(":8888")
}
