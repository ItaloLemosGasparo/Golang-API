package main

import (
	"fmt"
	"projeto/controladores"
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Funcionando")
	inicializadores.CarregarVariaveisDeAmbiente()
	inicializadores.ConectarAoBD()
	inicializadores.BD.AutoMigrate(&modelos.Endereco{})
}

func main() {
	r := gin.Default()

	r.POST("/usuario", controladores.CadastrarUsuario)
	r.GET("/usuario", controladores.BuscarUsuarios)
	r.GET("/usuario/:id", controladores.BuscarUsuario)
	r.PUT("/usuario/:id", controladores.AtualizarUsuario)
	r.PUT("/atualizarsenha/:id", controladores.AtualizarSenhaUsuario)
	r.DELETE("/usuario/:id", controladores.DeletarUsuario)
	r.POST("/usuario/login", controladores.Login)
	r.POST("/endereco", controladores.CadastrarEndereco)

	r.Run()
}
