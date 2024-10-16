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


utilizando docker
- mariadb@latest
- rabbitmq 4-management-alpine

alterar /internal/entity
- adicionar listOrders em interface.go ✔
- implementar listOrders em order.go ✔
- adicionar teste em order_test.go ✔

alterar /internal/usecase
- adicionar list_order.go ✔
- checar necessidade alterar algo em /pkg/events (em princípio nada) ✔

alterar /internal/infra/database
- adicionar consulta ao banco de dados ✔
- criar migrations com sql para criação da tabela ✔

alterar /internal/event
- adicionar evento e handler para list order ✔

alterar /internal/infra/web
- criar listorders_handler.go ✔
- altera wire.go para incluir list orders ✔
- executa wire ✔
- migrar wire_gen.go para outro package para evitar problemas de namespace ✔
- adicionar handler ao webserver em /cmd/ordersystem/main.go ✔
- foi necessário mover wire_gen.go para /internal/inject/wire_gen.go e mudar o package para inject wm virtude de conflito no namespace main

alterar /api
- adicionar arquivos .http para criar orders e listar orders ✔

alterar /internal/infra/graph (seguir https://gqlgen.com/getting-started/)
- executar go run github.com/99designs/gqlgen init  ✔
- alterar schema.graphqls para adicionar query ✔
- executar go run github.com/99designs/gqlgen generate ✔
- adicionar resolver a schema.resolvers.go  ✔
- implementar Orders em resolver.go ✔
- adicionar query ao graph server em /cmd/ordersystem/main.go ✔
- mover graph para /internal/ ✔

alterar /internal/infra/grpc
- adicionar service e messages to ./protofiles/order.proto
- usar protoc para gerar arquivos em ./pb
- adicionar list_orders.go em ./service
- adicionar service ao grpc server em /cmd/ordersystem/main.go



