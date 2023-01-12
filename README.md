# desafio_encurtador_url
Desafio encurtador de url proposto pela plataforma Devgym

## Tecnologias utilizadas:
```
uuid v1.3.0
testify v1.8.1
pq v1.10.7
```

## Como rodar o projeto?
Anteriomente, devemos criar um arquivo .db.env com as seguintes configurações:
```
POSTGRES_USER=my_username
POSTGRES_PASSWORD=my_password
POSTGRES_DB=my_dbname
```
Logo após, devemos rodar o comando `docker compose up -d` para subir o banco de dados

Devemos fazer a conexão e criar a tabela abaixo:
```
CREATE TABLE shortener(
  id VARCHAR(36) PRIMARY KEY,
  longurl TEXT NOT NULL,
  shorturl CHAR(6) NOT NULL UNIQUE,
  expiresin TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
```

Nas strings de conexão, devemos altera-lo de acordo com o usuário criado!

## Comandos utilizados:
```
go test ./...
go run ./...
```

## API
Há 2 rotas disponíveis:

POST /cut passando um json: { "url": "myurl" } retorna { "short": "mycode" }

GET /uncut?code=mycode retorna { "url": "myurl" }
