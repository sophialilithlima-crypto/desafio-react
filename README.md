# Desafio React 3 — CRUD Full Stack

Aplicação web full stack desenvolvida como parte do Desafio React 3, utilizando **React**, **Go**, **Gin**, **PostgreSQL** e **Docker**.

O projeto permite realizar o gerenciamento de **categorias, fornecedores e produtos**, com operações de criação, consulta, atualização e exclusão.

---

## Tecnologias utilizadas

### Frontend

* React
* Vite
* JavaScript
* HTML5
* CSS3
* Node.js
* npm

### Backend

* Go
* Gin
* PostgreSQL
* API REST

### Infraestrutura

* Docker
* Docker Compose

---

## Funcionalidades

### Categorias

* Listar categorias
* Buscar categoria por ID
* Cadastrar categoria
* Atualizar categoria
* Excluir categoria
* Validação dos dados

### Fornecedores

* Listar fornecedores
* Buscar fornecedor por ID
* Cadastrar fornecedor
* Atualizar fornecedor
* Excluir fornecedor
* Validação dos dados

### Produtos

* Listar produtos
* Buscar produto por ID
* Cadastrar produto
* Atualizar produto
* Excluir produto
* Validação dos dados
* Associação com categoria
* Controle de preço e estoque

### Outros recursos

* Paginação
* Busca
* Validações no backend
* Respostas padronizadas da API
* Persistência dos dados em PostgreSQL
* Testes automatizados no backend
* Configuração para execução com Docker Compose

---

## Estrutura do projeto

```text
desafio-react3/
│
├── backend/
│   ├── cmd/
│   ├── config/
│   ├── controllers/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   ├── utils/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── ...
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── ...
│   ├── public/
│   ├── Dockerfile
│   ├── package.json
│   └── ...
│
├── database/
│   └── ...
│
├── docker-compose.yml
└── README.md
```

---

## Requisitos

Para executar o projeto utilizando Docker, é necessário ter instalado:

* Docker
* Docker Compose

O Docker Desktop já inclui o Docker Compose nas versões atuais.

---

## Como executar o projeto

### 1. Clone o repositório

```bash
git clone <URL_DO_REPOSITORIO>
```

Entre na pasta do projeto:

```bash
cd desafio-react3
```

---

### 2. Inicie os containers

Execute:

```bash
docker compose up --build
```

O Docker irá:

1. Construir a imagem do frontend.
2. Construir a imagem do backend.
3. Criar o container do PostgreSQL.
4. Criar a rede da aplicação.
5. Criar o volume de persistência do banco.
6. Iniciar o backend.
7. Iniciar o frontend.

Na primeira execução, o PostgreSQL pode levar alguns segundos para inicializar. O backend possui reinício automático caso tente se conectar antes de o banco estar pronto.

---

## Acessando a aplicação

Depois que os containers estiverem em execução, acesse:

**Frontend**

```text
http://localhost:5173
```

**Backend**

```text
http://localhost:8080
```

A aplicação frontend utiliza a API disponibilizada pelo backend para realizar as operações no banco de dados.

---

## API

O backend disponibiliza endpoints para as principais entidades da aplicação.

### Categorias

```text
GET    /categorias
POST   /categorias
PUT    /categorias/:id
DELETE /categorias/:id
GET    /categoria/:id
```

### Fornecedores

```text
GET    /fornecedores
POST   /fornecedores
PUT    /fornecedores/:id
DELETE /fornecedores/:id
GET    /fornecedor/:id
```

### Produtos

```text
GET    /produtos
POST   /produtos
PUT    /produtos/:id
DELETE /produtos/:id
GET    /produto/:id
```

---

## Paginação e busca

Os endpoints de listagem possuem suporte a paginação e busca.

Exemplo:

```text
GET /produtos?page=1&limit=10&q=notebook
```

Onde:

* `page` representa a página atual.
* `limit` representa a quantidade de registros por página.
* `q` representa o termo de busca.

---

## Banco de dados

O projeto utiliza **PostgreSQL**.

O banco é executado através de um container Docker e seus dados são armazenados em um volume:

```text
postgres_data
```

Isso permite que os dados continuem disponíveis mesmo após a parada dos containers.

As configurações de conexão são definidas através de variáveis de ambiente utilizadas pelo backend.

---

## Testes

Os testes do backend podem ser executados dentro da pasta `backend` com:

```bash
go test ./...
```

O comando executa os testes disponíveis nos diferentes pacotes do backend.

---

## Parando a aplicação

Para parar os containers:

```bash
docker compose down
```

Para parar os containers e também remover os volumes:

```bash
docker compose down -v
```

> **Atenção:** o segundo comando remove o volume do PostgreSQL e, consequentemente, os dados armazenados nele.

---

## Verificando os containers

Para verificar os containers em execução:

```bash
docker compose ps
```

Para visualizar os logs:

```bash
docker compose logs
```

Para visualizar os logs de um serviço específico:

```bash
docker compose logs backend
```

ou:

```bash
docker compose logs frontend
```

ou:

```bash
docker compose logs postgres
```

---

## Execução validada

A aplicação foi executada e testada utilizando:

```bash
docker compose up --build
```

Durante a inicialização, foram criados e executados:

* Container do PostgreSQL
* Container do backend
* Container do frontend
* Rede Docker da aplicação
* Volume para persistência do PostgreSQL

Após a inicialização do PostgreSQL, o backend realizou a conexão com o banco com sucesso e iniciou o servidor na porta `8080`.

O frontend foi iniciado pelo Vite na porta `5173`.

As funcionalidades da aplicação foram testadas após a inicialização dos containers.

---

## Status do projeto

**Projeto concluído.**

* [x] Frontend React
* [x] Backend Go/Gin
* [x] PostgreSQL
* [x] CRUD de categorias
* [x] CRUD de fornecedores
* [x] CRUD de produtos
* [x] Validações
* [x] Paginação
* [x] Busca
* [x] Testes
* [x] Docker
* [x] Docker Compose
* [x] Integração frontend/backend
* [x] Persistência do banco de dados
* [x] Testes da aplicação em ambiente Docker

