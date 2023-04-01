package modelos

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome       string `json:"nome" gorm:"not null;type:varchar(50)"`
	Telefone   string `json:"telefone" gorm:"type:varchar(14)"`
	TelefoneB  string `json:"telefone_b" gorm:"type:varchar(14)"`
	Email      string `json:"email" gorm:"not null;uniqueIndex;type:varchar(128)"`
	CPF        string `json:"cpf" gorm:"uniqueIndex;type:varchar(14)"`
	Privilegio string `json:"privilegio"`
}
type Endereco struct {
	Id_Usuario int    `json:"id_usuario" gorm:"primaryKey;not null;foreignKey:id"`
	Logradouro string `json:"logradouro" gorm:"not null;type:varchar(100)"`
	Numero     int    `json:"numero" gorm:"not null"`
	Bairro     string `json:"bairro" gorm:"not null;type:varchar(50)"`
	Cidade     string `json:"cidade" gorm:"not null;type:varchar(50)"`
	Uf         string `json:"uf" gorm:"not null;type:varchar(2)"`
	Cep        string `json:"cep" gorm:"not null;type:varchar(9)"`
}

type Senhas struct {
	Id_Usuario int    `json:"id_usuario" gorm:"primaryKey;not null;foreignKey:id"`
	SenhaA     string `json:"SenhaA" gorm:"not null;type:varchar(256)"`
}

type Fornecedor struct {
	gorm.Model
	Nome      string `json:"nome" gorm:"not null;type:varchar(50)"`
	Email     string `json:"email" gorm:"not null;uniqueIndex;type:varchar(128)"`
	Telefone  string `json:"telefone" gorm:"type:varchar(11)"`
	TelefoneB string `json:"telefone_b" gorm:"type:varchar(11)"`
	CPF       string `json:"cpf" gorm:"uniqueIndex;type:varchar(14)"`
	CNPJ      string `json:"cnpj" gorm:"uniqueIndex;type:varchar(18)"`
}

type Produto struct {
	gorm.Model
	Id_Fornecedor uint    `json:"id_Fornecedor" gorm:"not null;index;foreignKey:id"`
	Nome          string  `json:"nome" gorm:"not null;type:varchar(50)"`
	Descricao     string  `json:"descricao" gorm:"not null;type:varchar(2048)"`
	Preco         float64 `json:"preco" gorm:"not null"`
}

type Carrinho struct {
	Id         int `json:"id" gorm:"primaryKey"`
	Id_Usuario int `json:"id_usuario" gorm:"not null"`
}

type Items_Carrinho struct {
	Id          int     `json:"id" gorm:"primaryKey"`
	Id_Carrinho int     `json:"id_carrinho" gorm:"not null"`
	ID_Produto  int     `json:"id_produto" gorm:"not null"`
	Quantidade  float32 `json:"quantidade" gorm:"not null"`
}

type Pedido struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Id_Carrinho int    `json:"nome" gorm:"not null"`
	Id_Usuario  int    `json:"id_carrinho" gorm:"not null"`
	Situacao    string `json:"situacao" gorm:"not null"`
}

type Favoritos struct {
	Id_Usuario int `json:"id_usuario" gorm:"primaryKey"`
	Id_Produto int `json:"id_produto" gorm:"primaryKey"`
}
