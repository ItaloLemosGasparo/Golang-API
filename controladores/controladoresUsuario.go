package controladores

import (
	"projeto_def/modelos"

	"github.com/gin-gonic/gin"
)

func CriarCadastroUsuario(c *gin.Context) {
	//Receber dados

	var usuarioTemp struct {
		Nome       string
		Telefone   string
		TelefoneB  string
		Email      string
		CPF        string
		Privilegio string
	}

	c.Bind(&usuarioTemp)

	//Criar o post

	usuario := modelos.Usuario{Nome: usuarioTemp.Nome,
		Telefone:   usuarioTemp.Telefone,
		TelefoneB:  usuarioTemp.TelefoneB,
		Email:      usuarioTemp.Email,
		CPF:        usuarioTemp.CPF,
		Privilegio: usuarioTemp.Privilegio,
	}

	result := inicializadores.BD.Create(&usuario)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//retornar func

	c.JSON(200, gin.H{
		"post": usuario,
	})

}
