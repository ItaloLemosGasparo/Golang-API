package controladores

import (
	"errors"
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CadastrarProduto(c *gin.Context) {
	var ProdutoTemp struct {
		Id_Fornecedor int     `json:"id_fornecedor"`
		Nome          string  `json:"nome"`
		Descricao     string  `json:"descricao"`
		Preco         float64 `json:"preco"`
	}

	c.Bind(&ProdutoTemp)

	Produto := modelos.Produto{
		Id_Fornecedor: ProdutoTemp.Id_Fornecedor,
		Nome:          ProdutoTemp.Nome,
		Descricao:     ProdutoTemp.Descricao,
		Preco:         ProdutoTemp.Preco,
	}

	if err := inicializadores.BD.Create(&Produto).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao cadastrar o produto": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Produto cadastrado com sucesso"})
}

func DeletarProduto(c *gin.Context) {
	id := c.Param("id")

	if err := inicializadores.BD.Delete(&modelos.Produto{}, id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao excluir o produto": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Produto excluído com sucesso"})
}

func AtualizarProduto(c *gin.Context) {
	id := c.Param("id")

	var produtoTemp struct {
		Id_Fornecedor int     `json:"id_fornecedor"`
		Nome          string  `json:"nome"`
		Descricao     string  `json:"descricao"`
		Preco         float64 `json:"preco"`
	}

	c.Bind(&produtoTemp)

	var produto modelos.Fornecedor

	if err := inicializadores.BD.First(&produto, id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao buscar o produto ": err.Error()})
		return
	}

	if err := inicializadores.BD.Model(&produto).Updates(modelos.Produto{
		Id_Fornecedor: produtoTemp.Id_Fornecedor,
		Nome:          produtoTemp.Nome,
		Descricao:     produtoTemp.Descricao,
		Preco:         produtoTemp.Preco,
	}).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao atualizar o produto ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"produto": produto})
}

func BuscarProduto(c *gin.Context) {
	id := c.Param("id")

	var produto modelos.Produto

	if err := inicializadores.BD.First(&produto, id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao buscar o produto ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"produto": produto})
}

func BuscarProdutos(c *gin.Context) {
	var produtos []modelos.Produto

	if err := inicializadores.BD.Find(&produtos).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"produtos": produtos})
}

func AdicionarProdutoFavorito(c *gin.Context) {
	var request struct {
		Id_Usuario int `json:"id_usuario"`
		Id_Produto int `json:"id_produto"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var usuario modelos.Usuario
	var favorito modelos.Favoritos
	var produto modelos.Produto

	//Usuario existe?
	if err := inicializadores.BD.First(&usuario, request.Id_Usuario).Error; err != nil {
		c.JSON(400, gin.H{"Usuario não encontrado ": err.Error()})
		return
	}

	//Produto existe?
	if err := inicializadores.BD.First(&produto, request.Id_Produto).Error; err != nil {

		favorito = modelos.Favoritos{
			Id_Usuario: request.Id_Usuario,
			Id_Produto: request.Id_Produto,
		}

		if err := inicializadores.BD.Create(&favorito).Error; err != nil {
			c.JSON(400, gin.H{"Erro ao criar o favorito": err.Error()})
			return
		}

	} else {
		c.JSON(400, gin.H{"Erro ao buscar o produto ": err.Error()})
		return
	}
}

func AdicionarProdutoCarrinho(c *gin.Context) {
	var request struct {
		Id_Usuario int `json:"id_usuario"`
		Id_Produto int `json:"id_produto"`
		Quantidade int `json:"quantidade"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	carrinho := modelos.Carrinho{
		Id_Usuario: request.Id_Usuario,
	}

	if err := inicializadores.BD.First(&carrinho, "id_usuario = ?", request.Id_Usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//Criando carrinho
			if err := inicializadores.BD.Create(&carrinho).Error; err != nil {
				c.JSON(400, gin.H{"erro ao criar o carrinho ": err.Error()})
				return
			}

			item_Carrinho := modelos.Items_Carrinho{
				Id_Carrinho: carrinho.Id,
				ID_Produto:  request.Id_Produto,
				Quantidade:  request.Quantidade,
			}

			//Adicionando o produto ao mesmo
			if err := inicializadores.BD.Create(&item_Carrinho).Error; err != nil {
				c.JSON(400, gin.H{"Erro ao adicionar o produto ao carrinho ": err.Error()})
				return
			}
		} else {
			c.JSON(400, gin.H{"Erro ao buscar o carrinho do usuario ": err.Error()})
			return
		}
	} else {
		item_Carrinho := modelos.Items_Carrinho{
			Id_Carrinho: carrinho.Id,
			ID_Produto:  request.Id_Produto,
			Quantidade:  request.Quantidade,
		}
		//Adicionando o produto ao carrinho
		if err := inicializadores.BD.Create(&item_Carrinho).Error; err != nil {
			c.JSON(400, gin.H{"Erro ao adicionar o produto ao carrinho ": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{"message": "Produto adicionado ao carrinho com sucesso"})
}
