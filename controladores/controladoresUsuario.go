package controladores

import (
	"projeto/inicializadores"
	"projeto_def/modelos"

	"github.com/gin-gonic/gin"
)

func CadastrarUsuario(c *gin.Context) {
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

func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")

	inicializadores.BD.Delete(&modelos.Usuario{}, id)

	c.Status(200)
}
