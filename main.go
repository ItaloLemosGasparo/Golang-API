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

	inicializadores.BD.AutoMigrate(&modelos.Usuario{})
	inicializadores.BD.AutoMigrate(&modelos.Senhas{})
	inicializadores.BD.AutoMigrate(&modelos.Endereco{})
	inicializadores.BD.AutoMigrate(&modelos.Fornecedor{})
	inicializadores.BD.AutoMigrate(&modelos.Produto{})
	inicializadores.BD.AutoMigrate(&modelos.Carrinho{})
	inicializadores.BD.AutoMigrate(&modelos.Items_Carrinho{})
	inicializadores.BD.AutoMigrate(&modelos.Pedido{})
	inicializadores.BD.AutoMigrate(&modelos.Favoritos{})
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

	r.GET("/carrinho/:id", controladores.BuscarCarrinho)        //id do usuario
	r.GET("/favoritos/:idU/:idP", controladores.BuscarFavorito) //id do usuario / id do produto
	r.GET("/favoritos_/:id", controladores.BuscarFavoritos)     //id do usuario

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
