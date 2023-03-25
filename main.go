package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Print("Funcionou meu chapa")
}

func main() {
	r := gin.Default()

	//r.POST("/posts", controladores.CriarPost)
	//r.GET("/posts", controladores.BuscarPosts)
	//r.GET("/posts/:id", controladores.BuscarPost)
	//r.PUT("/posts/:id", controladores.AtualizarPost)
	//r.DELETE("/posts/:id", controladores.DeletarPost)

	r.Run()
}
