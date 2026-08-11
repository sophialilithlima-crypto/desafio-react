# Desafio Técnico — CRUD Full-Stack Nível 3

Aplicação full-stack de cadastro de produtos, categorias e fornecedores, desenvolvida com Go, React e PostgreSQL.

O projeto contempla os requisitos dos níveis anteriores e adiciona os recursos do **Nível 3 — Bônus**: relacionamento N:N entre produtos e fornecedores, autenticação simples, CI, deploy/preview e logs estruturados.

---

## Tecnologias

### Backend

* Go 1.22+
* Gin
* PostgreSQL
* JWT
* `log/slog`
* Repository Pattern
* Arquitetura Controller → Service → Repository

### Frontend

* React 18+
* Vite
* JavaScript
* React Router
* Axios

### Banco de dados

* PostgreSQL 14+

### Infraestrutura

* Docker
* Docker Compose
* GitHub Actions
* Render

---

## Funcionalidades

### Categoria

* Criar categoria
* Listar categorias
* Consultar categoria por ID
* Editar categoria
* Excluir categoria
* Buscar por nome
* Paginação
* Validação
* Nome único
* Bloqueio da exclusão quando existem produtos vinculados

### Fornecedor

* Criar fornecedor
* Listar fornecedores
* Consultar fornecedor por ID
* Editar fornecedor
* Excluir fornecedor
* Buscar por nome
* Paginação
* Validação
* Validação de e-mail

### Produto

* Criar produto
* Listar produtos
* Consultar produto por ID
* Editar produto
* Excluir produto
* Buscar por nome
* Paginação
* SKU único
* Validação de preço
* Validação de estoque
* Relação com categoria
* Relação com múltiplos fornecedores

---

## Relacionamentos

### Categoria → Produto

A relação entre categoria e produto é **1:N**.

```text
Categoria 1 ───── N Produto
```

Um produto pertence a uma categoria e uma categoria pode possuir vários produtos.

A exclusão de uma categoria que possui produtos vinculados é bloqueada.

### Produto ↔ Fornecedor

A relação entre produto e fornecedor é **N:N**.

```text
Produto N ───── N Fornecedor
       \         /
        \       /
     produto_fornecedor
```

A relação é armazenada na tabela:

```text
produto_fornecedor
├── produto_id
└── fornecedor_id
```

A chave primária é composta pelos dois IDs.

---

## Autenticação

O sistema possui autenticação simples utilizando JWT.

### Login

```http
POST /auth/login
Content-Type: application/json
```

Exemplo:

```json
{
  "usuario": "admin",
  "senha": "admin123"
}
```

Após o login, a API retorna um token JWT.

As rotas protegidas utilizam:

```http
Authorization: Bearer SEU_TOKEN
```

As credenciais devem ser configuradas através das variáveis de ambiente.

---

## API

### Autenticação

| Método | Endpoint      |
| ------ | ------------- |
| POST   | `/auth/login` |

### Categorias

| Método | Endpoint              |
| ------ | --------------------- |
| GET    | `/api/categorias`     |
| POST   | `/api/categorias`     |
| GET    | `/api/categorias/:id` |
| PUT    | `/api/categorias/:id` |
| DELETE | `/api/categorias/:id` |

### Fornecedores

| Método | Endpoint                |
| ------ | ----------------------- |
| GET    | `/api/fornecedores`     |
| POST   | `/api/fornecedores`     |
| GET    | `/api/fornecedores/:id` |
| PUT    | `/api/fornecedores/:id` |
| DELETE | `/api/fornecedores/:id` |

### Produtos

| Método | Endpoint            |
| ------ | ------------------- |
| GET    | `/api/produtos`     |
| POST   | `/api/produtos`     |
| GET    | `/api/produtos/:id` |
| PUT    | `/api/produtos/:id` |
| DELETE | `/api/produtos/:id` |

O padrão REST utilizado segue o contrato definido no desafio.

---

## Paginação e busca

As listagens possuem paginação e busca.

Exemplo:

```http
GET /api/produtos?page=1&limit=20&q=notebook
```

A resposta utiliza um envelope consistente:

```json
{
  "data": [],
  "error": null,
  "meta": {
    "total": 0,
    "page": 1,
    "limit": 20
  }
}
```

---

## Estrutura do projeto

```text
desafio-react3/
│
├── .github/
│   └── workflows/
│       └── ci.yml
│
├── backend/
│   ├── cmd/
│   ├── config/
│   ├── controllers/
│   ├── middleware/
│   ├── migrations/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   ├── utils/
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── database/
│   ├── database.sql
│   └── migration_n3.sql
│
├── frontend/
│   ├── public/
│   ├── src/
│   ├── Dockerfile
│   ├── package.json
│   └── vite.config.js
│
├── .env.example
├── docker-compose.yml
├── render.yaml
└── README.md
```

---

## Arquitetura

O backend utiliza três camadas principais:

```text
Controller
    ↓
Service
    ↓
Repository
    ↓
PostgreSQL
```

### Controller

Responsável pelo recebimento das requisições HTTP e retorno das respostas.

### Service

Responsável pelas regras de negócio e validações.

### Repository

Responsável pelo acesso ao banco de dados.

As queries utilizam parâmetros, evitando concatenação direta de valores no SQL. Essa separação segue os requisitos técnicos do desafio.

---

## Logs estruturados

O backend utiliza `log/slog` para gerar logs estruturados em JSON.

Exemplo:

```json
{
  "level": "INFO",
  "msg": "http_request",
  "method": "GET",
  "path": "/api/produtos",
  "status": 200,
  "duration_ms": 12
}
```

O objetivo é facilitar a leitura e o processamento dos logs em ambientes de execução.

---

## Variáveis de ambiente

Utilize o arquivo `.env.example` como referência.

Exemplo:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=desafio-react2

PORT=8080

JWT_SECRET=change-this-secret

AUTH_USER=admin
AUTH_PASSWORD=admin123

FRONTEND_URL=http://localhost:5173
```

**Não coloque senhas ou secrets reais no GitHub.**

---

## Executando com Docker

Com o Docker Desktop instalado e aberto, entre na pasta do projeto:

```powershell
cd C:\Users\Sophia\Documents\GitHub\desafio-react3
```

Execute:

```powershell
docker compose up --build
```

O Docker Compose sobe os serviços necessários:

```text
PostgreSQL
    ↓
Backend Go
    ↓
Frontend React
```

O desafio define como requisito que o `docker-compose.yml` seja capaz de subir PostgreSQL, API e Front com um comando.

### Frontend

```text
http://localhost:5173
```

### Backend

```text
http://localhost:8080
```

Para parar os containers:

```powershell
docker compose down
```

Para remover também os volumes:

```powershell
docker compose down -v
```

---

## Executando sem Docker

### Backend

```powershell
cd backend
go mod tidy
go run ./cmd
```

### Frontend

Em outro terminal:

```powershell
cd frontend
npm install
npm run dev
```

---

## Banco de dados

O schema principal está localizado em:

```text
database/database.sql
```

A alteração referente ao relacionamento N:N está em:

```text
database/migration_n3.sql
```

A relação entre produto e fornecedor utiliza a tabela:

```text
produto_fornecedor
```

---

## Testes

Para executar os testes do backend:

```powershell
cd backend
go test ./...
```

Os testes devem cobrir principalmente regras de negócio dos services e endpoints relevantes.

O desafio prioriza testes relevantes em vez de exigir cobertura de 100%.

---

## CI

O projeto possui GitHub Actions configurado em:

```text
.github/workflows/ci.yml
```

O pipeline executa verificações do backend e frontend.

Backend:

```text
go mod download
go test ./...
```

Frontend:

```text
npm ci
npm run build
```

O CI permite verificar automaticamente o projeto a cada alteração enviada ao GitHub.

---

## Deploy / Preview

O projeto possui configuração de deploy em:

```text
render.yaml
```

A configuração foi preparada para permitir o deploy da aplicação utilizando:

* Backend Go
* Frontend React
* PostgreSQL

As variáveis sensíveis devem ser configuradas diretamente no ambiente de deploy.

---

## Decisões de arquitetura

### Go + Gin

Gin foi utilizado como framework HTTP por ser leve e adequado para a construção da API REST.

### Repository Pattern

O Repository Pattern separa o acesso ao banco das regras de negócio e facilita a realização de testes.

### PostgreSQL

PostgreSQL foi escolhido para garantir persistência real e suportar as relações entre as entidades.

### JWT

JWT foi utilizado para implementar uma autenticação simples, conforme o escopo do Nível 3.

### React + Vite

React e Vite foram utilizados para a construção da SPA.

### Docker Compose

Docker Compose permite executar os serviços da aplicação de maneira padronizada.

---

## Trade-offs

O projeto prioriza simplicidade e atendimento ao escopo do desafio.

A autenticação implementada é simples e não possui um sistema completo de usuários, permissões ou refresh tokens.

O sistema também não implementa uma plataforma completa de observabilidade, utilizando logs estruturados como mecanismo de observabilidade previsto no Nível 3.

---

## O que ficou de fora

Não foram adicionadas funcionalidades que não fazem parte do escopo definido para o desafio.

Possíveis evoluções futuras seriam:

* gerenciamento completo de usuários;
* diferentes níveis de permissão;
* refresh tokens;
* documentação Swagger/OpenAPI;
* maior cobertura de testes;
* monitoramento externo;
* pipeline de deploy mais avançado.

---

## Entregáveis

O projeto contém:

* Repositório Git;
* CRUD das três entidades;
* Relação 1:N Categoria → Produto;
* Relação N:N Produto ↔ Fornecedor;
* Validações;
* Paginação;
* Busca;
* Autenticação;
* Logs estruturados;
* Testes;
* Docker Compose;
* GitHub Actions;
* Configuração de deploy;
* README.

Esses elementos correspondem aos requisitos definidos para os níveis do desafio.

---

## Desenvolvedora

**Sophia Lilith**

Projeto desenvolvido como desafio técnico Full-Stack utilizando Go, React e PostgreSQL.
