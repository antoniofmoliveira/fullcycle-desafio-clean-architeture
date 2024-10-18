# Clean Architecture

Olá devs!

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)

- Service ListOrders com GRPC

- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Solução

### utilizando docker

- mariadb:latest

- rabbitmq:4-management-alpine

### alterar /internal/entity

- adicionar listOrders em interface.go ✔

- implementar listOrders em order.go ✔

- adicionar teste em order_test.go ✔

### alterar /internal/usecase

- adicionar list_order.go ✔

- checar necessidade alterar algo em /pkg/events (em princípio nada) ✔

### alterar /internal/database

- adicionar consulta ao banco de dados ✔

- criar migrations com sql para criação da tabela ✔

### alterar /internal/event

- adicionar evento e handler para list order ✔

### alterar /internal/web

- criar listorders_handler.go ✔

- altera wire.go para incluir list orders ✔

- executa wire ✔

- migrar wire_gen.go para outro package para evitar problemas de namespace ✔

- adicionar handler ao webserver em /cmd/ordersystem/main.go ✔

- foi necessário mover wire_gen.go para /internal/inject/wire_gen.go e mudar o package para inject wm virtude de conflito no namespace main

### alterar /api

- adicionar arquivos .http para criar orders e listar orders ✔

### alterar /internal/graph (seguir https://gqlgen.com/getting-started/)

- executar go run github.com/99designs/gqlgen init  ✔

- alterar schema.graphqls para adicionar query ✔

- executar go run github.com/99designs/gqlgen generate ✔

- adicionar resolver a schema.resolvers.go  ✔

- implementar Orders em resolver.go ✔

- adicionar query ao graph server em /cmd/ordersystem/main.go ✔

- mover graph para /internal/ ✔

### alterar /internal/grpc

- adicionar service e messages to ./protofiles/order.proto ✔

- go get google.golang.org/grpc ✔

- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ✔

- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest ✔

- usar protoc para gerar arquivos em ./pb  ✔

- protoc --go-grpc_out --go_out=pb --go_opt=paths=source_relative protofiles/order.proto  ✔

- criar ListOrdersService em /service ✔

- adicionar service ao grpc server em /cmd/ordersystem/main.go ✔

- testar: ✔

    - grpcurl -plaintext -d '{"id": "1", "price": 100, "tax": 10}' localhost:50051 pb.OrderService/CreateOrder

    - grpcurl -plaintext localhost:50051 pb.ListOrderService/ListOrders


### implementar docker composer

- docker build --pull --rm -f "Dockerfile" -t cleanarch:latest "." ✔

- criar a network que vai connectar o docker compose ✔

- docker compose -f "docker-compose.yaml" up -d --build  ✔

- atentar para as variáveis de ambiente para rodar local ✔

  - DB_HOST=172.18.0.3 - tem que ser o ip do container

  - AMQP_HOST=localhost

- atentar para as variaveis de ambiente para rodar no docker compose ✔

  - DB_HOST=mariadbca - tem que ser o hostname colocado no service do docker compose

  - AMQP_HOST=rabbitmq - tem que ser o hostname colocado no service do docker compose

- assim `.env_container` contém as variáveis de ambiente para rodar em container e `.env` contém as necessárias para executar localmente. a linha `COPY .env_container .env` no `dockerfile` copia o `.env_container` para `.env`, automatizando o processoa ao construir o container.

- **IMPORTANTE**: após a carga do container o rabbitmq demora cerca de 3 segundos para iniciar em minha máquina, impedindo o aplicativo de se conectar ao rabbimq causando panic.
isso se traduz na necessidade de adicionar no `docker-compose.yml` um healthcheck no service do rabbitmq e uma condition no depends_on do app.
um adicional healthcheck foi adicionado ao service do mariadb só por garantia.

### executando servidor e log de mensagens

```bash
antonio@DG15:~/DEV/go2/clean-arch$ air

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.61.1, built with Go go1.23.2

watching .
watching api
watching cmd
watching cmd/ordersystem
watching configs
watching internal
watching internal/entity
watching internal/event
watching internal/event/handler
watching internal/graph
watching internal/graph/model
watching internal/grpc
watching internal/grpc/pb
watching internal/grpc/protofiles
watching internal/grpc/service
watching internal/database
watching internal/web
watching internal/web/webserver
watching internal/inject
watching internal/usecase
watching migrations
watching pkg
watching pkg/amqpclientgo
watching pkg/events
!exclude tmp
building...
running...
Starting web server on port :8080
Starting gRPC server on port 50051
Starting GraphQL server on port 8081
Order listed: <nil>
Order listed: [{a 100.5 0.5 101} {b 100.5 0.5 101} {c 100.5 0.5 101}]
Order created: {1 100 10 110}/Order listed: [{a 100.5 0.5 101} {b 100.5 0.5 101} {c 100.5 0.5 101}]
Order listed: [{1 100 10 110} {a 100.5 0.5 101} {b 100.5 0.5 101} {c 100.5 0.5 101}]
```

### interação grpc

```bash
antonio@DG15:~/DEV/go2/clean-arch$ grpcurl -plaintext -d '{"id": "1", "price": 100, "tax": 10}' localhost:50051 pb.OrderService/CreateOrder
{
  "id": "1",
  "price": 100,
  "tax": 10,
  "finalPrice": 110
}
antonio@DG15:~/DEV/go2/clean-arch$ grpcurl -plaintext localhost:50051 pb.ListOrderService/ListOrders
{
  "orders": [
    {
      "id": "1",
      "price": 100,
      "tax": 10,
      "finalPrice": 110
    },
    {
      "id": "a",
      "price": 100.5,
      "tax": 0.5,
      "finalPrice": 101
    },
    {
      "id": "b",
      "price": 100.5,
      "tax": 0.5,
      "finalPrice": 101
    },
    {
      "id": "c",
      "price": 100.5,
      "tax": 0.5,
      "finalPrice": 101
    }
  ]
}
```

### interação graphql

```graphql
{orders {
  id,Price, Tax, FinalPrice
}}
```
### saída graphql
```json
{
  "data": {
    "orders": [
      {
        "id": "1",
        "Price": 100,
        "Tax": 10,
        "FinalPrice": 110
      },
      {
        "id": "a",
        "Price": 100.5,
        "Tax": 0.5,
        "FinalPrice": 101
      },
      {
        "id": "b",
        "Price": 100.5,
        "Tax": 0.5,
        "FinalPrice": 101
      },
      {
        "id": "c",
        "Price": 100.5,
        "Tax": 0.5,
        "FinalPrice": 101
      }
    ]
  }
}
```

### interação api

```http
GET http://localhost:8080/orders HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

### saída api

```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 17 Oct 2024 13:16:26 GMT
Content-Length: 211
Connection: close

[
  {
    "id": "1",
    "price": 100,
    "tax": 10,
    "final_price": 110
  },
  {
    "id": "a",
    "price": 100.5,
    "tax": 0.5,
    "final_price": 101
  },
  {
    "id": "b",
    "price": 100.5,
    "tax": 0.5,
    "final_price": 101
  },
  {
    "id": "c",
    "price": 100.5,
    "tax": 0.5,
    "final_price": 101
  }
]
```

