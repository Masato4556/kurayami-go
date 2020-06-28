package server

import (
    "github.com/gin-gonic/gin"

    "github.com/Masato4556/kurayami-go/controller"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    u := r.Group("/todos")
    {
        ctrl := todo.Controller{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.Show)
        u.POST("", ctrl.Create)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

    return r
}