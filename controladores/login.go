package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var login struct {
		Email string
		Senha string
	}

	c.Bind(&login)

	var usuario modelos.Usuario
	if result := inicializadores.BD.Where("email = ?", login.Email).First(&usuario); result.Error != nil {
		c.Status(400)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(login.Senha)); err != nil {
		c.Status(400)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  usuario.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("sua_chave_secreta_aqui"))
	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
