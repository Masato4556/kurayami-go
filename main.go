package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Masato4556/kurayami-go/db"
	"github.com/Masato4556/kurayami-go/server"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	db.Init()
	server.Init()
}
