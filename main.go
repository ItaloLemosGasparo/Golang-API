package main

import (
	"fmt"
	"projeto_def/inicializadores"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Funcionou meu chapa")
	inicializadores.CarregarVariaveisDeAmbiente()
	inicializadores.ConectarAoBD()
}

func main() {
	r := gin.Default()

	//r.POST("/usuario", controladores.CadastrarUsuario)
	//r.GET("/usuario", controladores.BuscarUsuarios)
	//r.GET("/usuario/:id", controladores.BuscarUsuario)
	//r.PUT("/usuario/:id", controladores.AtualizarUsuario)
	r.DELETE("/deletarUsuario/:id")

	r.Run()
}
