package main

import (
	"projeto/inicializadores"
	"projeto/modelos"
)

func init() {
	inicializadores.CarregarVariaveisDeAmbiente()
	inicializadores.ConectarAoBD()
}

func main() {
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
