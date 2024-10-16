package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/antoniofmoliveira/cleanarch/configs"
	"github.com/antoniofmoliveira/cleanarch/internal/event/handler"
	"github.com/antoniofmoliveira/cleanarch/internal/infra/web/webserver"
	"github.com/antoniofmoliveira/cleanarch/internal/inject"
	"github.com/antoniofmoliveira/cleanarch/pkg/events"

	"github.com/streadway/amqp"

	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
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

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := inject.NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

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

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

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

	// if _, err := db.Exec(
	// 	"CREATE TABLE IF NOT EXISTS message (value TEXT PRIMARY KEY)"); err != nil {
	// 	return nil, err
	// }

	return db, nil
}
