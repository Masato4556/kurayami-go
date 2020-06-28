package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Masato4556/kurayami-go/db"
	"github.com/asuforce/gin-gorm-tutorial/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	db.Init()
	server.Init()
}
