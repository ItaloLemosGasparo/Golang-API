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
