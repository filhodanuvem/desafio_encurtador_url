# desafio_encurtador_url
Desafio encurtador de url proposto pela plataforma Devgym

## Tecnologias utilizadas:
```
uuid v1.3.0
testify v1.8.1
pq v1.10.7
godotenv v1.4.0
```

## Como rodar o projeto?
Anteriomente, devemos criar um arquivo 2 arquivos, .env e .db.env com as seguintes configurações:

.env deve conter o `BASE_URL` para a conexão com o banco de dados
```
BASE_URL="postgres://username:password@host:port/dbname?sslmode=disable"
```

.db.env
```
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_DB=dbname
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

## Comandos utilizados:
```
go test ./...
go run ./...
```

## API
Há 2 rotas disponíveis:

POST /cut passando um json: { "url": "myurl" } retorna { "short": "mycode" }

GET /uncut?code=mycode retorna { "url": "myurl" }
