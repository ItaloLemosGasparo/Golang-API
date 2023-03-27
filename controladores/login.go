package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Função para criar um token JWT
func criarToken(usuario modelos.Usuario) (string, error) {
	// Definir o tempo de expiração do token
	expiraEm := time.Now().Add(time.Hour * 24).Unix()

	// Definir o payload do token
	claims := jwt.MapClaims{
		"id":    usuario.ID,
		"email": usuario.Email,
		"exp":   expiraEm,
	}

	// Gerar o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func Login(c *gin.Context) {
	var usuarioTemp struct {
		Email string
		Senha string
	}

	c.Bind(&usuarioTemp)

	//Buscando o usuario
	var usuario modelos.Usuario
	inicializadores.BD.Where("email = ?", usuarioTemp.Email).First(&usuario)

	//Buscando a senha do usuario
	var senha modelos.Senhas
	inicializadores.BD.Where("id_Usuario = ?", usuario.ID).First(&senha)

	//Pegar o id do usuario e dar um find tanto em usuario quanto em senha vaseado no id -----------------------------------------
	err := bcrypt.CompareHashAndPassword([]byte(senha.SenhaA), []byte(usuarioTemp.Senha))

	if err != nil {
		c.Status(401)
		return
	}

	// Criar token JWT
	token, err := criarToken(usuario)
	if err != nil {
		c.Status(500)
		return
	}

	// Responder com o token JWT
	c.JSON(200, gin.H{
		"token": token,
	})
}
