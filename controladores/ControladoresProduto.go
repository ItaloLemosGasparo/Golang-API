package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CadastrarProduto(c *gin.Context) {
	var ProdutoTemp struct {
		Id_Fornecedor int
		Nome          string
		Descricao     string
		Preco         float64
	}

	c.Bind(&ProdutoTemp)

	Produto := modelos.Produto{
		Id_Fornecedor: ProdutoTemp.Id_Fornecedor,
		Nome:          ProdutoTemp.Nome,
		Descricao:     ProdutoTemp.Descricao,
		Preco:         ProdutoTemp.Preco,
	}

	if result := inicializadores.BD.Create(&Produto); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"produto": Produto,
	})
}

func CadastrarFavorito(c *gin.Context) {
	idU := c.Param("idU")
	idP := c.Param("idP")

	var usuario modelos.Usuario
	inicializadores.BD.First(&usuario, idU)

	var favorito modelos.Favoritos

	if inicializadores.BD.First(&favorito.Id_Produto, idP) == nil {
		idUs, err := strconv.Atoi(idU)
		idPr, err2 := strconv.Atoi(idP)

		if err != nil && err2 != nil {
			favorito = modelos.Favoritos{
				Id_Usuario: idUs,
				Id_Produto: idPr,
			}
			if result := inicializadores.BD.Create(&favorito); result.Error != nil {
				c.Status(400)
				return
			}
		} else {
			c.Status(400)
			return
		}
	} else {
		c.Status(400)
		return
	}
}

func AddCarrinho(c *gin.Context) {
	//map para armazenar os itens do carrinho.
	var carrinho = make(map[int]int)
	var item Item
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Bad request"})
		return
	}

	// Verifica se o produto já está add no BD
	var produto modelos.Produto
	if err := inicializadores.BD.First(&produto, item.Id_Produto).Error; err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Produto não encontrado"})
		return
	}

	// Adiciona o item ao carrinho
	carrinho[item.IdProduto] += item.Quantidade

	// Retorna a resposta com status de sucesso
	c.JSON(200, gin.H{
		"message":  "Item adicionado ao carrinho",
		"carrinho": carrinho,
	})
}
