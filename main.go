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

	//Usuarios
	r.POST("/usuario", controladores.CadastrarUsuario)
	r.PUT("/usuario/:id", controladores.AtualizarUsuario)
	r.DELETE("/usuario/:id", controladores.DeletarUsuario)
	r.GET("/usuario/:id", controladores.BuscarUsuario)
	r.GET("/usuario", controladores.BuscarUsuarios)

	r.PUT("/atualizarsenha/:id", controladores.AtualizarSenhaUsuario)
	r.PUT("/cadastrarTelefone/:id", controladores.AtualizarTelefone)
	r.PUT("/cadastrarEndereco", controladores.AtualizarEndereco)

	r.GET("/produto/:id", controladores.BuscarCarrinho)       //id do usuario
	r.GET("/produto/:idU/:idP", controladores.BuscarFavorito) //id do usuario / id do produto
	r.GET("/produto/:id", controladores.BuscarFavoritos)      //id do usuario

	//Login
	r.POST("/usuario/login", controladores.Login)

	//Fornecedores
	r.POST("/fornecedor", controladores.CadastrarFornecedor)
	r.PUT("/fornecedor/:id", controladores.AtualizarFornecedor)
	r.GET("/fornecedor/:id", controladores.BuscarFornecedor)
	r.GET("/fornecedor", controladores.BuscarFornecedores)

	//Prudotos
	r.POST("/produto", controladores.CadastrarProduto)
	r.POST("/produto/:idU/:idP", controladores.AdicionarProdutoFavorito)      //idUsuario /: idProduto
	r.POST("/produto/:idU/:idP/:qtd", controladores.AdicionarProdutoCarrinho) //idUsuario /: idProduto /: Quantidade
	r.GET("/produto/:id", controladores.BuscarProduto)
	r.GET("/produto", controladores.BuscarProdutos)
	r.PUT("/produto/:id", controladores.AtualizarProduto)
	//Buscar x nos favoritos

	r.Run()
}
