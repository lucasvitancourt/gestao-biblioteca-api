
# Gerenciamento de Biblioteca

Este é um projeto de API RESTful para um sistema de gerenciamento de biblioteca. A API permite gerenciar livros, autores e empréstimos de livros.

## Estrutura do Projeto

```
gerenciamento-biblioteca/
├── main.go
├── modelos/
│   ├── livro.go
│   ├── autor.go
│   └── emprestimo.go
├── controladores/
│   ├── livro_controlador.go
│   ├── autor_controlador.go
│   └── emprestimo_controlador.go
├── roteadores/
│   └── roteador.go
├── banco/
│   └── banco.go
├── middleware/
│   └── autenticacao.go
├── go.mod
└── go.sum
```

## Pré-requisitos

- [Go](https://golang.org/doc/install) 1.16+
- [Gorilla Mux](https://github.com/gorilla/mux)
- [GORM](https://gorm.io/index.html)
- [SQLite](https://www.sqlite.org/index.html)

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/lucasvitancourt/gerenciamento-biblioteca.git
    cd gerenciamento-biblioteca
    ```

2. Inicialize o módulo Go:
    ```sh
    go mod init gerenciamento-biblioteca
    go mod tidy
    ```

## Configuração do Banco de Dados

O projeto usa SQLite como banco de dados. O arquivo do banco de dados será criado automaticamente na primeira execução.

## Inicialização do Servidor

Para iniciar o servidor, execute o comando:
```sh
go run main.go
```

O servidor será iniciado na porta 8080.

## Endpoints da API

### Livros

- **Obter todos os livros**
  ```
  GET /livros
  ```
  **Resposta de exemplo:**
  ```json
  [
    {
      "id": 1,
      "titulo": "Aprender Go",
      "autor_id": 1,
      "disponivel": true
    }
  ]
  ```

- **Obter um livro por ID**
  ```
  GET /livros/{id}
  ```
  **Resposta de exemplo:**
  ```json
  {
    "id": 1,
    "titulo": "Aprender Go",
    "autor_id": 1,
    "disponivel": true
  }
  ```

- **Criar um novo livro**
  ```
  POST /livros
  ```
  **Corpo da requisição:**
  ```json
  {
    "titulo": "Novo Livro",
    "autor_id": 1,
    "disponivel": true
  }
  ```

- **Atualizar um livro**
  ```
  PUT /livros/{id}
  ```
  **Corpo da requisição:**
  ```json
  {
    "titulo": "Livro Atualizado",
    "autor_id": 1,
    "disponivel": false
  }
  ```

- **Deletar um livro**
  ```
  DELETE /livros/{id}
  ```

### Autores

- **Obter todos os autores**
  ```
  GET /autores
  ```
  **Resposta de exemplo:**
  ```json
  [
    {
      "id": 1,
      "nome": "Autor Exemplo",
      "livros": [
        {
          "id": 1,
          "titulo": "Aprender Go",
          "autor_id": 1,
          "disponivel": true
        }
      ]
    }
  ]
  ```

- **Obter um autor por ID**
  ```
  GET /autores/{id}
  ```
  **Resposta de exemplo:**
  ```json
  {
    "id": 1,
    "nome": "Autor Exemplo",
    "livros": [
      {
        "id": 1,
        "titulo": "Aprender Go",
        "autor_id": 1,
        "disponivel": true
      }
    ]
  }
  ```

- **Criar um novo autor**
  ```
  POST /autores
  ```
  **Corpo da requisição:**
  ```json
  {
    "nome": "Novo Autor"
  }
  ```

- **Atualizar um autor**
  ```
  PUT /autores/{id}
  ```
  **Corpo da requisição:**
  ```json
  {
    "nome": "Autor Atualizado"
  }
  ```

- **Deletar um autor**
  ```
  DELETE /autores/{id}
  ```

### Empréstimos

- **Obter todos os empréstimos**
  ```
  GET /emprestimos
  ```
  **Resposta de exemplo:**
  ```json
  [
    {
      "id": 1,
      "livro_id": 1,
      "usuario": "Usuário Exemplo",
      "data_emprestimo": "2022-01-01T00:00:00Z",
      "data_devolucao": "2022-01-15T00:00:00Z"
    }
  ]
  ```

- **Obter um empréstimo por ID**
  ```
  GET /emprestimos/{id}
  ```
  **Resposta de exemplo:**
  ```json
  {
    "id": 1,
    "livro_id": 1,
    "usuario": "Usuário Exemplo",
    "data_emprestimo": "2022-01-01T00:00:00Z",
    "data_devolucao": "2022-01-15T00:00:00Z"
  }
  ```

- **Criar um novo empréstimo**
  ```
  POST /emprestimos
  ```
  **Corpo da requisição:**
  ```json
  {
    "livro_id": 1,
    "usuario": "Novo Usuário",
    "data_emprestimo": "2022-02-01T00:00:00Z",
    "data_devolucao": "2022-02-15T00:00:00Z"
  }
  ```

- **Atualizar um empréstimo**
  ```
  PUT /emprestimos/{id}
  ```
  **Corpo da requisição:**
  ```json
  {
    "livro_id": 1,
    "usuario": "Usuário Atualizado",
    "data_emprestimo": "2022-03-01T00:00:00Z",
    "data_devolucao": "2022-03-15T00:00:00Z"
  }
  ```

- **Deletar um empréstimo**
  ```
  DELETE /emprestimos/{id}
  ```
