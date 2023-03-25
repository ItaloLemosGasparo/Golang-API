package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Funcionou meu chapa")
}

func main() {
	r := gin.Default()

	//r.POST("/usuario", controladores.CriarPost)
	//r.GET("/usuario", controladores.BuscarPosts)
	//r.GET("/usuario/:id", controladores.BuscarPost)
	//r.PUT("/usuario/:id", controladores.AtualizarPost)
	//r.DELETE("/usuario/:id", controladores.DeletarPost)

	r.Run()
}
