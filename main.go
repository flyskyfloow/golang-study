package main

import (
	"github.com/flyskyfloow/golang-study/todo"
	"github.com/gin-gonic/gin"
)

func main() {
	store := todo.NewStore()
	h := todo.NewHandler(store)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	h.RegisterRoutes(v1)

	r.Run(":8080")
}
