package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func deletarUsuario(c *gin.Context) {
	id := c.Param("id")

	inicializadores.BD.Delete(&modelos.Usuario{}, id)

	c.Status(200)
}
