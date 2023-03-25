package main

import (
	"fmt"
	"projeto/controladores"
	"projeto/inicializadores"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Funcionando")
	inicializadores.CarregarVariaveisDeAmbiente()
	inicializadores.ConectarAoBD()
}

func main() {
	r := gin.Default()

	r.POST("/usuario", controladores.CadastrarUsuario)
	r.GET("/usuario", controladores.BuscarUsuarios)
	r.GET("/usuario/:id", controladores.BuscarUsuario)
	r.PUT("/usuario/:id", controladores.AtualizarUsuario)
	r.PUT("/atualizarsenha/:id", controladores.AtualizarSenhaUsuario)
	r.DELETE("/usuario/:id", controladores.DeletarUsuario)

	r.Run()
}
