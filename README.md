# Descrição do Projeto: API em Golang

Este projeto é uma API RESTful desenvolvida em Go (Golang), projetada para gerenciar usuários, produtos, fornecedores e carrinhos de compras. A API utiliza várias bibliotecas e frameworks populares para facilitar o desenvolvimento e a manutenção.

## Tecnologias Utilizadas:

- **Godotenv:** Para gerenciamento de variáveis de ambiente.
- **Gin-Gonic:** Framework web para construção de APIs de alta performance.
- **GORM:** ORM (Object-Relational Mapping) para facilitar a interação com o banco de dados.
- **JWT-Go:** Para implementação de autenticação baseada em tokens JWT.
- **PostgreSQL:** Banco de dados utilizado para armazenamento de dados.

## Principais Funcionalidades:

- **Gerenciamento de Usuários:** 
  - Cadastro, atualização e exclusão de usuários.
  - Atualização de senhas e informações adicionais (como telefone e endereço).
  - Autenticação de usuários via login.

- **Gerenciamento de Produtos:**
  - Cadastro, atualização e consulta de produtos.
  
- **Gerenciamento de Fornecedores:**
  - Cadastro, atualização e consulta de fornecedores.

- **Gerenciamento de Carrinho e Favoritos:**
  - Adição de produtos aos favoritos e ao carrinho de compras.
  - Consulta de produtos no carrinho e favoritos.

## Estrutura do Código:
