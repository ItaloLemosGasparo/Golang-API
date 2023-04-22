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

	if err := inicializadores.BD.Delete(&modelos.Usuario{}, id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuario excluído com sucesso"})
}

func BuscarUsuarios(c *gin.Context) {
	var usuarios []modelos.Usuario

	if err := inicializadores.BD.Find(&usuarios).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"usuarios": usuarios})
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuario modelos.Usuario

	if err := inicializadores.BD.First(&usuario, id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"usuario": usuario})
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

	if err := inicializadores.BD.First(&usuario, id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := inicializadores.BD.Model(&usuario).Updates(modelos.Usuario{
		Nome:       usuarioTemp.Nome,
		Telefone:   usuarioTemp.Telefone,
		TelefoneB:  usuarioTemp.TelefoneB,
		Email:      usuarioTemp.Email,
		CPF:        usuarioTemp.CPF,
		Privilegio: usuarioTemp.Privilegio,
	}).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"usuario": usuario})
}

func AtualizarSenhaUsuario(c *gin.Context) {
	id := c.Param("id")

	//Verifica se o usuario existe
	if err := inicializadores.BD.First("usuarios", id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Verifica se o usuario tem uma senha cadastrada
	var senha modelos.Senhas

	if err := inicializadores.BD.Where("Id_Usuario = ?", id).First(&senha).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var senhaTemp struct {
		Senha string
	}

	c.Bind(&senhaTemp)

	//Verifica se a nova senha é igual a atual
	if bcrypt.CompareHashAndPassword([]byte(senha.SenhaA), []byte(senhaTemp.Senha)) != nil {
		senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(senhaTemp.Senha), bcrypt.DefaultCost)
		if err != nil {
			c.Status(400)
			return
		}

		senha.SenhaA = string(senhaCriptografada)
		if result := inicializadores.BD.Save(&senha); result.Error != nil {
			c.Status(400)
			return
		}

		c.JSON(200, gin.H{
			"senhaA": senha.SenhaA,
		})
	} else {
		c.Status(400)
		return
	}

}

func AtualizarEndereco(c *gin.Context) {
	var enderecoTemp struct {
		Id_usuario int
		Logradouro string
		Numero     int
		Bairro     string
		Cidade     string
		Uf         string
		Cep        string
	}

	c.Bind(&enderecoTemp)

	endereco := modelos.Endereco{
		Id_Usuario: enderecoTemp.Id_usuario,
		Logradouro: enderecoTemp.Logradouro,
		Numero:     enderecoTemp.Numero,
		Bairro:     enderecoTemp.Bairro,
		Cidade:     enderecoTemp.Cidade,
		Uf:         enderecoTemp.Uf,
		Cep:        enderecoTemp.Cep,
	}

	result := inicializadores.BD.Create(&endereco)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"endereco": endereco,
	})
}

func AtualizarTelefone(c *gin.Context) {
	id := c.Param("id")

	var telefoneTemp struct {
		Telefone  string
		TelefoneB string
	}

	c.Bind(&telefoneTemp)

	if telefoneTemp.Telefone != telefoneTemp.TelefoneB {
		var usuario modelos.Usuario
		inicializadores.BD.First(&usuario, id)

		inicializadores.BD.Model(&usuario).Updates(modelos.Usuario{
			Telefone:  telefoneTemp.Telefone,
			TelefoneB: telefoneTemp.TelefoneB,
		})

		c.JSON(200, gin.H{
			"usuario": usuario,
		})
	} else {
		c.Status(400)
		return
	}
}

func BuscarCarrinho(c *gin.Context) {
	id := c.Param("id")

	var Carrinho modelos.Carrinho

	if err := inicializadores.BD.First(&Carrinho, id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"carrinho": Carrinho})
}

func BuscarFavorito(c *gin.Context) {
	idU := c.Param("id")
	idP := c.Param("id")

	var Favorito modelos.Favoritos

	if err := inicializadores.BD.First(&Favorito.Id_Usuario, idU, &Favorito.Id_Produto, idP).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"favorito": Favorito})
}

func BuscarFavoritos(c *gin.Context) {
	var Favoritos []modelos.Favoritos

	if err := inicializadores.BD.Find(&Favoritos).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"favoritos": Favoritos})
}
