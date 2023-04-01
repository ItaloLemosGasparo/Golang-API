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
}
