package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CadastrarUsuario(c *gin.Context) {
	var usuarioTemp struct {
		Nome       string
		Telefone   string
		TelefoneB  string
		Email      string
		CPF        string
		Privilegio string
		Senha      string
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

	if result := inicializadores.BD.Create(&usuario); result.Error != nil {
		c.Status(400)
		return
	}

	senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(usuarioTemp.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.Status(400)
		return
	}

	senha := modelos.Senhas{
		Id_Usuario: int(usuario.ID),
		SenhaA:     string(senhaCriptografada),
	}

	if result := inicializadores.BD.Create(&senha); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"usuario":          usuario,
		"senha_cadastrada": senha.SenhaA,
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

func AtualizarSenhaUsuario(c *gin.Context) {
	// Ler o id do usuário da rota
	id := c.Param("id")

	var senhaTemp struct {
		Senha string
	}

	// Ler a nova senha do corpo da requisição
	c.Bind(&senhaTemp)

	// Criptografar a nova senha
	senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(senhaTemp.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.Status(400)
		return
	}

	// Buscar o usuário no banco de dados pelo id
	var usuario modelos.Usuario
	if result := inicializadores.BD.First(&usuario, id); result.Error != nil {
		c.Status(400)
		return
	}

	// Atualizar a senha do usuário no banco de dados
	var senha modelos.Senhas
	if result := inicializadores.BD.Where("usuario_id = ?", id).First(&senha); result.Error != nil {
		c.Status(400)
		return
	}
	senha.SenhaA = string(senhaCriptografada)
	if result := inicializadores.BD.Save(&senha); result.Error != nil {
		c.Status(400)
		return
	}

	// Responder com a nova senha criptografada
	c.JSON(200, gin.H{
		"Nova_senha": senha.SenhaA,
	})
}
