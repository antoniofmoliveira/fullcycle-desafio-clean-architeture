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


foi necessário mover wire_gen.go para /internal/inject/wire_gen.go e mudar o package para inject


alterar /internal/entity
- adicionar listOrders em interface.go
- implementar listOrders em order.go
- adicionar teste em order_test.go

alterar /internal/infra/database
- adicionar consulta ao banco de dados

alterar /internal/usecase
- adicionar list_order.go
- checar necessidade alterar algo em /pkg/events (em princípio nada)

alterar /internal/infra/grpc
- adicionar service e messages to ./protofiles/order.proto
- usar protoc para gerar arquivos em ./pb
- adicionar list_orders.go em ./service
- adicionar service ao grpc server em /cmd/ordersystem/main.go

alterar /internal/infra/graph (seguir https://gqlgen.com/getting-started/)
- executar go run github.com/99designs/gqlgen init
- copiar generated.go gerado em /graph para /internal/infra/graph (dispatch agora precisa de contexto) 
- alterar schema.graphqls para adicionar query
- adicionar resolver a schema.resolvers.go
- adicionar usecase a resolver.go
- adicionar query ao graph server em /cmd/ordersystem/main.go

alterar /internal/infra/web
- criar listorders_handler.go
- adicionar handler ao webserver em /cmd/ordersystem/main.go

alterar /api
- adicionar arquivos .http para criar orders e listar orders
