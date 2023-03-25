package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func CadastrarUsuario(c *gin.Context) {
	var usuarioTemp struct {
		Nome       string
		Telefone   string
		TelefoneB  string
		Email      string
		CPF        string
		Privilegio string
	}

	c.Bind(&usuarioTemp)

	usuario := modelos.Usuario{
		Nome:       usuarioTemp.Nome,
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

	c.JSON(200, gin.H{
		"usuarioz": usuario,
	})
}

func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")

	inicializadores.BD.Delete(&modelos.Usuario{}, id)

	c.Status(200)
}

func BuscarUsuarios(c *gin.Context) {
	var usuarios []modelos.Usuario
	inicializadores.BD.Find(&usuarios)

	c.JSON(200, gin.H{
		"usuarios": usuarios,
	})
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuario modelos.Usuario
	inicializadores.BD.First(&usuario, id)

	c.JSON(200, gin.H{
		"usuario": usuario,
	})
}

func AtualizarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuarioTemp struct {
		Nome       string
		Telefone   string
		TelefoneB  string
		Email      string
		CPF        string
		Privilegio string
	}

	c.Bind(&usuarioTemp)

	var usuario modelos.Usuario
	inicializadores.BD.First(&usuario, id)

	inicializadores.BD.Model(&usuario).Updates(modelos.Usuario{
		Nome:       usuarioTemp.Nome,
		Telefone:   usuarioTemp.Telefone,
		TelefoneB:  usuarioTemp.TelefoneB,
		Email:      usuarioTemp.Email,
		CPF:        usuarioTemp.CPF,
		Privilegio: usuarioTemp.Privilegio,
	})

	c.JSON(200, gin.H{
		"usuario": usuario,
	})
}

func CadastrarSenhaUsuario(c *gin.Context) {
	id := c.Param("id")

	var senhaTemp struct {
		Id_Usuario int
		SenhaA     string
		SenhaB     string
	}

	c.Bind(&senhaTemp)

	senha := modelos.Senhas{
		Id_Usuario: senhaTemp.Id_Usuario,
		SenhaA:     senhaTemp.SenhaA,
		SenhaB:     senhaTemp.SenhaB,
	}

	result := inicializadores.BD.Create(&senha)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Senha": senha,
	})
}
