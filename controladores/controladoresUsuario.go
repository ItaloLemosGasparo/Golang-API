package controladores

import (
	"projeto_def/inicializadores"
	"projeto_def/modelos"

	"github.com/gin-gonic/gin"
)

func deletarUsuario(c *gin.Context) {
	id := c.Param("id")

	inicializadores.BD.Delete(&modelos.Usuario{}, id)

	c.Status(200)
}
