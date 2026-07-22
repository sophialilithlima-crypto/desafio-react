# Desafio React - Sistema de Gerenciamento de Produtos

## 📌 Sobre o projeto

Este projeto foi desenvolvido como parte de um desafio técnico utilizando **React**, **Go** e **PostgreSQL**.

O sistema permite o gerenciamento de:

- Categorias
- Fornecedores
- Produtos

Cada produto possui uma categoria e um fornecedor associados por meio de chaves estrangeiras no banco de dados.

---

## 🚀 Tecnologias Utilizadas

### Frontend
- React
- Vite
- JavaScript

### Backend
- Go
- Gin Gonic
- PostgreSQL Driver (lib/pq)

### Banco de Dados
- PostgreSQL
- pgAdmin 4

---

## 📁 Estrutura do Projeto

```text
desafio-react/
│
├── backend/
│   ├── cmd/
│   ├── config/
│   ├── controllers/
│   ├── models/
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── vite.config.js
│
└── README.md
```

---

## 🗄️ Banco de Dados

O projeto utiliza PostgreSQL com três tabelas:

- categoria
- fornecedor
- produto

Relacionamentos:

- Uma categoria pode possuir vários produtos.
- Um fornecedor pode fornecer vários produtos.
- Cada produto pertence a uma categoria e a um fornecedor.

---

## ⚙️ Como executar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/SEU-USUARIO/desafio-react.git
```

Entre na pasta do projeto:

```bash
cd desafio-react
```

---

## Backend

Entre na pasta:

```bash
cd backend
```

Instale as dependências:

```bash
go mod tidy
```

Execute o servidor:

```bash
go run ./cmd
```

O backend será iniciado em:

```
http://localhost:8080
```

---

## Frontend

Abra outro terminal.

Entre na pasta:

```bash
cd frontend
```

Instale as dependências:

```bash
npm install
```

Execute o projeto:

```bash
npm run dev
```

O frontend será iniciado em:

```
http://localhost:5173
```

---

## 📡 Endpoints da API

### Categoria

| Método | Endpoint |
|---------|----------|
| GET | `/categoria` |
| POST | `/categoria` |
| PUT | `/categoria/:id` |
| DELETE | `/categoria/:id` |

### Fornecedor

| Método | Endpoint |
|---------|----------|
| GET | `/fornecedor` |
| POST | `/fornecedor` |
| PUT | `/fornecedor/:id` |
| DELETE | `/fornecedor/:id` |

### Produto

| Método | Endpoint |
|---------|----------|
| GET | `/produto` |
| POST | `/produto` |
| PUT | `/produto/:id` |
| DELETE | `/produto/:id` |

---

## 📄 Exemplos de Requisições

### Categoria

```json
{
  "nome": "Eletrônicos"
}
```

### Fornecedor

```json
{
  "nome": "Fornecedor A",
  "email": "fornecedor@email.com"
}
```

### Produto

```json
{
  "nome": "Notebook",
  "preco": 3500.00,
  "categoria_id": 1,
  "fornecedor_id": 1
}
```

---

## 🛠️ Funcionalidades

- Cadastro de categorias
- Listagem de categorias
- Atualização de categorias
- Exclusão de categorias

- Cadastro de fornecedores
- Listagem de fornecedores
- Atualização de fornecedores
- Exclusão de fornecedores

- Cadastro de produtos
- Listagem de produtos
- Atualização de produtos
- Exclusão de produtos

---

## 👩‍💻 Desenvolvedora

**Sophia Lilith**

Projeto desenvolvido como parte de um desafio de desenvolvimento Full Stack utilizando React, Go e PostgreSQL.
