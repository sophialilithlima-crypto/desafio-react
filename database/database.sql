DROP TABLE IF EXISTS produto;
DROP TABLE IF EXISTS fornecedor;
DROP TABLE IF EXISTS categoria;

CREATE TABLE categoria (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE fornecedor (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    telefone VARCHAR(30)
);

CREATE TABLE produto (
    id SERIAL PRIMARY KEY,

    nome VARCHAR(150) NOT NULL,

    sku VARCHAR(100) NOT NULL UNIQUE,

    preco NUMERIC(10,2) NOT NULL
        CHECK (preco > 0),

    estoque INTEGER NOT NULL
        CHECK (estoque >= 0),

    categoria_id INTEGER NOT NULL,

    criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_categoria
        FOREIGN KEY (categoria_id)
        REFERENCES categoria(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);