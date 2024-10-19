# Descrição do Projeto: API em Golang

Este projeto é uma API RESTful desenvolvida em Go (Golang), projetada para gerenciar usuários, produtos, fornecedores e carrinhos de compras. A API utiliza várias bibliotecas e frameworks populares para facilitar o desenvolvimento e a manutenção.

- ## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **Godotenv:** Para gerenciamento de variáveis de ambiente.
- **Gin**: Framework para construção de APIs web.
- **GORM**: ORM (Object Relational Mapping) para interagir com o banco de dados.
- **JWT-Go:** Para implementação de autenticação baseada em tokens JWT.
- **bcrypt**: Para criptografia de senhas.
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

# Resumo da Estrutura do Projeto

## Estrutura Geral

O projeto é construído em Go utilizando o framework Gin para desenvolvimento de APIs REST. O código é organizado em pacotes, cada um com uma responsabilidade específica.

## Pacotes Principais

### 1. `controladores`

Contém os controladores que gerenciam as rotas e a lógica de negócios. Aqui estão algumas das principais funções:

- **Gerenciamento de Usuários**: Gerencia a criação e atualização de usuários, incluindo validação de senha.
- **Gerenciamento de Fornecedores**: Permite cadastrar, atualizar, buscar e excluir fornecedores.
- **Gerenciamento de Produtos**: Controla as operações de cadastro, atualização, busca e exclusão de produtos.
- **Autenticação de Usuários**: Implementa o login e a geração de tokens JWT para autenticação e criptografia.

### 2. `modelos`

Define as estruturas de dados que representam as entidades do sistema, como:

- **Usuario**: Armazena informações do usuário, incluindo dados de contato e privilégios.
- **Senhas**: Armazena a senha criptografada associada a um usuário.
- **Fornecedor**: Representa um fornecedor com seus dados de contato e identificação.
- **Produto**: Representa um produto com informações relacionadas, como preço e descrição.
- **Carrinho**: Gerencia os itens que um usuário adiciona para compra.
- **Favoritos**: Permite que usuários marquem produtos como favoritos.

### 3. `inicializadores`

Contém funções para configurar e inicializar a aplicação, incluindo:

- **Conexão com o Banco de Dados**: Gerencia a conexão com o banco de dados usando GORM.
- **Carregamento de Variáveis de Ambiente**: Carrega as variáveis de configuração necessárias para a aplicação.

### 4. `migracao.go`

Responsável pela migração do banco de dados. Utiliza o método `AutoMigrate` do GORM para criar e atualizar as tabelas de acordo com os modelos definidos.

## Funcionalidades Principais

- **Cadastro e Autenticação de Usuários**: Permite que os usuários se registrem e façam login utilizando senhas criptografadas.
- **Gerenciamento de Fornecedores e Produtos**: Oferece CRUD (criação, leitura, atualização e exclusão) completo para fornecedores e produtos.
- **Favoritos e Carrinho de Compras**: Permite que usuários adicionem produtos aos favoritos e gerenciem seu carrinho de compras.

## Conclusão

O projeto segue uma arquitetura modular, onde cada pacote tem uma responsabilidade específica. Isso facilita a manutenção e a expansão do sistema, permitindo a adição de novas funcionalidades de forma organizada.
