package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CadastrarUsuario(c *gin.Context) {
	var usuarioTemp struct {
		Nome       string `json:"nome"`
		Telefone   string `json:"telefone"`
		TelefoneB  string `json:"telefoneb"`
		Email      string `json:"email"`
		CPF        string `json:"cpf"`
		Privilegio string `json:"privilegio"`
		Senha      string `json:"senha"`
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

	if err := inicializadores.BD.Create(&usuario).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao cadastrar o usuario ": err.Error()})
		return
	}

	senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(usuarioTemp.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error ": err.Error()})
		return
	}

	senha := modelos.Senhas{
		Id_Usuario: int(usuario.ID),
		SenhaA:     string(senhaCriptografada),
	}

	if err := inicializadores.BD.Create(&senha).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao cadastrar a senha ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuario cadastrado com sucesso"})
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
		c.JSON(400, gin.H{"Erro ao buscar os usuarios ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"usuarios": usuarios})
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuario modelos.Usuario

	if err := inicializadores.BD.First(&usuario, id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao buscar o usuario ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"usuario": usuario})
}

func AtualizarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuarioTemp struct {
		Nome       string `json:"nome"`
		Telefone   string `json:"telefone"`
		TelefoneB  string `json:"telefoneb"`
		Email      string `json:"email"`
		CPF        string `json:"cpf"`
		Privilegio string `json:"privilegio"`
	}

	c.Bind(&usuarioTemp)

	var usuario modelos.Usuario

	if err := inicializadores.BD.First(&usuario, id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao buscar o usuario ": err.Error()})
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
		c.JSON(400, gin.H{"Erro ao atualizar o cadastro do usuario ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuario atualizado com sucesso"})
}

func AtualizarSenhaUsuario(c *gin.Context) {
	id := c.Param("id")

	//Verifica se o usuario existe
	if err := inicializadores.BD.First("usuarios", id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao buscar o usuario ": err.Error()})
		return
	}

	var senha modelos.Senhas

	//Verifica se o usuario tem uma senha cadastrada
	if err := inicializadores.BD.First(&senha, "id_usuario = ?", id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao buscar a senha do usuario ": err.Error()})
		return
	}

	var senhaTemp string
	c.Bind(&senhaTemp)

	//Verifica se a nova senha é igual a atual
	if err := bcrypt.CompareHashAndPassword([]byte(senha.SenhaA), []byte(senhaTemp)); err != nil {
		senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(senhaTemp), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{"Error ": err.Error()})
			return
		}

		senha.SenhaA = string(senhaCriptografada)
		if result := inicializadores.BD.Save(&senha); result.Error != nil {
			c.JSON(400, gin.H{"Error ": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "senha atualizada com sucesso"})
	} else {
		c.JSON(400, gin.H{"Error ": err.Error()})
		return
	}
}

func AtualizarEndereco(c *gin.Context) {
	var enderecoTemp struct {
		Id_usuario int    `json:"id_usuario"`
		Logradouro string `json:"logradouro"`
		Numero     int    `json:"numero"`
		Bairro     string `json:"bairro"`
		Cidade     string `json:"cidade"`
		Uf         string `json:"uf"`
		Cep        string `json:"cep"`
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

	if err := inicializadores.BD.Create(&endereco).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao atualizar o endereço do usuario ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "endereço atualizada com sucesso"})
}

func AtualizarTelefone(c *gin.Context) {
	id := c.Param("id")

	var telefoneTemp struct {
		Telefone  string `json:"telefone"`
		TelefoneB string `json:"telefoneb"`
	}

	c.Bind(&telefoneTemp)

	if telefoneTemp.Telefone != telefoneTemp.TelefoneB {
		var usuario modelos.Usuario
		inicializadores.BD.First(&usuario, id)

		if err := inicializadores.BD.Model(&usuario).Updates(modelos.Usuario{
			Telefone:  telefoneTemp.Telefone,
			TelefoneB: telefoneTemp.TelefoneB,
		}).Error; err != nil {
			c.JSON(400, gin.H{"Erro ao atualizar o telefone do usuario ": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "telefone atualizada com sucesso"})
	} else {
		c.JSON(400, gin.H{"message": "telefoneA e TelefoneB não podem ser iguais"})
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
	var request struct {
		Id_Usuario int `json:"id_usuario"`
		Id_Produto int `json:"id_produto"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var Favorito modelos.Favoritos

	//Buscando o produto x nos favoritos
	if err := inicializadores.BD.First(&Favorito, "id_usuario = ?", request.Id_Usuario, &Favorito, "id_produto = ?", request.Id_Produto).Error; err != nil {
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
