package modelos

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome       string
	Telefone   string
	TelefoneB  string
	Email      string
	CPF        string
	Privilegio string
}

type Endereco struct {
	Id_Usuario int
	Logradouro string
	Numero     int
	Bairro     string
	Cidade     string
	Uf         string
	Cep        string
}

type Senhas struct {
	Id_Usuario int
	SenhaA     string
}

type Fornecedor struct {
	gorm.Model
	Nome      string
	Email     string
	Telefone  string
	TelefoneB string
	Cpf       string
	Cnpj      string
}

type Produto struct {
	gorm.Model
	Id_Fornecedor int
	Nome          string
	Descricao     string
	Preco         float32
}

type Carrinho struct {
	Id         int
	Id_Usuario int
}

type Items_Carrinho struct {
	Id          int
	Id_Carrinho int
	ID_Produto  int
	Quantidade  float32
}

type Pedido struct {
	Id          int
	Id_Carrinho int
	Id_Usuario  int
	Situacao    string
}

type Favoritos struct {
	Id_Usuario int
	Id_Produto int
}
