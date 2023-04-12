package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func CadastrarProduto(c *gin.Context)  {
	var ProdutoTemp struct{
		Id_Fornecedor int
		Nome          string
		Descricao     string
		Preco         float64
	}

	c.Bind(&ProdutoTemp)

	Produto := modelos.Produto{
		Id_Fornecedor: ProdutoTemp.Id_Fornecedor,
		Nome: ProdutoTemp.Nome,
		Descricao: ProdutoTemp.Descricao,
		Preco: ProdutoTemp.Preco,
	}

	if result := inicializadores.BD.Create(&Produto); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"produto": Produto,
	})
}