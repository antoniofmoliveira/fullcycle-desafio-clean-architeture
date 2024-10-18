createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:grilo007@tcp(localhost:3306)/ordersystem" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:grilo007@tcp(localhost:3306)/ordersystem" -verbose down

cleandirtyflag:
	migrate -path=sql/migrations -database "mysql://root:grilo007@tcp(localhost:3306)/ordersystem" force 1

.PHONY: migrate migratedown cleandirtyflag
