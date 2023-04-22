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
	r.GET("/usuario/:id", controladores.BuscarUsuario)
	r.GET("/usuario", controladores.BuscarUsuarios)
	r.PUT("/usuario/:id", controladores.AtualizarUsuario)
	r.PUT("/atualizarsenha/:id", controladores.AtualizarSenhaUsuario)
	r.DELETE("/usuario/:id", controladores.DeletarUsuario)
	r.PUT("/cadastrarTelefone/:id", controladores.AtualizarTelefone)
	r.POST("/cadastrarEndereco", controladores.CadastrarEndereco)

	//Login
	r.POST("/usuario/login", controladores.Login)

	//Fornecedores
	r.POST("/fornecedor", controladores.CadastrarFornecedor)
	r.GET("/fornecedor/:id", controladores.BuscarFornecedor)
	r.GET("/fornecedor", controladores.BuscarFornecedores)
	//Editar

	//Prudotos
	r.POST("/produto", controladores.CadastrarProduto)
	r.POST("/produto/:idU/:idP", controladores.AdicionarProdutoFavorito)      //idUsuario /: idProduto
	r.POST("/produto/:idU/:idP/:qtd", controladores.AdicionarProdutoCarrinho) //idUsuario /: idProduto /: Quantidade
	r.GET("/produto/:id", controladores.BuscarProduto)
	r.GET("/produto", controladores.BuscarProdutos)
	//Editar
	//Buscar Carrinho
	//Buscar Favoritos
	//Buscar x nos favoritos

	r.Run()
}
