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
	r.POST("/usuario", controladores.CadastrarUsuario)                //Cadastrar Usuario
	r.PUT("/usuario/:id", controladores.AtualizarUsuario)             //Atualizar Usuario
	r.PUT("/atualizarsenha/:id", controladores.AtualizarSenhaUsuario) //Atualizar senha do Usuario
	r.DELETE("/usuario/:id", controladores.DeletarUsuario)            //Deletar Usuario
	//Buscar
	r.GET("/usuario/:id", controladores.BuscarUsuario)
	r.GET("/usuario", controladores.BuscarUsuarios)
	//Cadastrar Extras
	r.PUT("/cadastrarTelefone/:id", controladores.AtualizarTelefone)
	r.PUT("/cadastrarEndereco", controladores.AtualizarEndereco)

	//Carrinho & Favoritos
	r.POST("/adicionarProdutoFavoritos", controladores.AdicionarProdutoFavorito) //recebo id usuario, id produto por JSON
	r.POST("/adicionarProdutoCarrinho", controladores.AdicionarProdutoCarrinho)  //recebo id usuario, id produto e quantidade por JSON
	//Buscar
	//r.GET("/BuscarProdutocarrinho", controladores.BuscarCarrinho)    //Buscar produto x no carrinho
	r.GET("/BuscarProdutoscarrinho/:id", controladores.BuscarCarrinho) //id do usuario
	r.GET("/buscarFavorito", controladores.BuscarFavorito)             //recebo id usuario, id produto por JSON
	r.GET("buscarFavoritos/:id", controladores.BuscarFavoritos)        //id do usuario

	//Fornecedores
	r.POST("/fornecedor", controladores.CadastrarFornecedor)    //Cadastrar Fornecedor
	r.PUT("/fornecedor/:id", controladores.AtualizarFornecedor) //Atualizar Fornecedor
	//Buscar
	r.GET("/fornecedor/:id", controladores.BuscarFornecedor)
	r.GET("/fornecedor", controladores.BuscarFornecedores)

	//Prudotos
	r.POST("/produto", controladores.CadastrarProduto)    //Cadastrar Produto
	r.PUT("/produto/:id", controladores.AtualizarProduto) //Atualizar Produto
	//Buscar
	r.GET("/produto/:id", controladores.BuscarProduto)
	r.GET("/produto", controladores.BuscarProdutos)

	//Login
	r.POST("/usuario/login", controladores.Login)

	r.Run()
}
