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

func DeletarProduto(c *gin.Context) {
	id := c.Param("id")

	if err := inicializadores.BD.Delete(&modelos.Produto{}, id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Produto excluído com sucesso"})
}

func AtualizarProduto(c *gin.Context) {
	id := c.Param("id")

	var produtoTemp struct {
		Id_Fornecedor int
		Nome          string
		Descricao     string
		Preco         float64
	}

	c.Bind(&produtoTemp)

	var produto modelos.Fornecedor

	if err := inicializadores.BD.First(&produto, id).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := inicializadores.BD.Model(&produto).Updates(modelos.Produto{
		Id_Fornecedor: produtoTemp.Id_Fornecedor,
		Nome:          produtoTemp.Nome,
		Descricao:     produtoTemp.Descricao,
		Preco:         produtoTemp.Preco,
	}).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"produto": produto})
}

func BuscarProduto(c *gin.Context) {
	id := c.Param("id")

	var produto modelos.Produto
	inicializadores.BD.First(&produto, id)

	c.JSON(200, gin.H{
		"produto": produto,
	})
}

func BuscarProdutos(c *gin.Context) {
	var produtos []modelos.Produto
	inicializadores.BD.Find(&produtos)

	c.JSON(200, gin.H{
		"produtos": produtos,
	})
}

func AdicionarProdutoFavorito(c *gin.Context) {
	idU := c.Param("idU")
	idP := c.Param("idP")

	var usuario modelos.Usuario
	inicializadores.BD.First(&usuario, idU)

	var favorito modelos.Favoritos

	var produto modelos.Produto

	if inicializadores.BD.First(&produto, idP) != nil {
		idUs, err := strconv.Atoi(idU)
		idPr, err2 := strconv.Atoi(idP)

		if err == nil && err2 == nil {
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

func AdicionarProdutoCarrinho(c *gin.Context) {
	idU := c.Param("idU")
	idP := c.Param("idP")
	qtd := c.Param("qtd")

	idUs, err := strconv.Atoi(idU)
	idPr, err2 := strconv.Atoi(idP)
	qtdP, err3 := strconv.ParseFloat(qtd, 64)

	if err == nil && err2 == nil && err3 == nil {
		carrinho := modelos.Carrinho{
			Id_Usuario: idUs,
		}

		if inicializadores.BD.First(&carrinho.Id_Usuario, idP) != nil {
			if result := inicializadores.BD.Create(&carrinho); result.Error != nil {
				c.Status(400)
				return
			}
			item_Carrinho := modelos.Items_Carrinho{
				Id_Carrinho: carrinho.Id,
				ID_Produto:  idPr,
				Quantidade:  qtdP,
			}
			if result := inicializadores.BD.Create(&item_Carrinho); result.Error != nil {
				c.Status(400)
				return
			}
		} else {
			item_Carrinho := modelos.Items_Carrinho{
				Id_Carrinho: carrinho.Id,
				ID_Produto:  idPr,
				Quantidade:  qtdP,
			}
			if result := inicializadores.BD.Create(&item_Carrinho); result.Error != nil {
				c.Status(400)
				return
			}
		}

	} else {
		c.Status(400)
		return
	}
}
