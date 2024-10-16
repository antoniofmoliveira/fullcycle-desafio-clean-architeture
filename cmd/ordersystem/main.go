package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/antoniofmoliveira/cleanarch/configs"
	"github.com/antoniofmoliveira/cleanarch/internal/graph"
	"github.com/antoniofmoliveira/cleanarch/internal/event/handler"
	"github.com/antoniofmoliveira/cleanarch/internal/infra/web/webserver"
	"github.com/antoniofmoliveira/cleanarch/internal/inject"
	"github.com/antoniofmoliveira/cleanarch/pkg/events"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := initStorage(*configs)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// RabbitMQ init

	rabbitMQChannel := getRabbitMQChannel(*configs)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	eventDispatcher.Register("OrderListed", &handler.OrderListedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// RabbitMQ end

	// webserver start

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := inject.NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)

	webOrderListHandler := inject.NewWebOrderListHandler(db, eventDispatcher)
	webserver.AddHandler("/orders", webOrderListHandler.List)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// webserver end

	//graphql server start

	createOrderUseCase := inject.NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := inject.NewListOrderUseCase(db, eventDispatcher)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

	// graphql server end

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-termChan
	log.Println("server: shutting down")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Could not shutdown the server: %v\n", err)
	// }
	fmt.Println("Server stopped")
	os.Exit(0)

}

func getRabbitMQChannel(configs configs.Config) *amqp.Channel {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/", configs.AmqpUser, configs.AmqpPassword, configs.AmqpHost, configs.AmqpPort)

	conn, err := amqp.Dial(addr)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func initStorage(configs configs.Config) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configs.DBUser,
		configs.DBPassword,
		configs.DBHost,
		configs.DBPort,
		configs.DBName,
	)

	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open(configs.DBDriver, connString)

	if err != nil {
		return nil, err
	}
	return db, nil
}
