FROM golang:1.23-alpine 

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.* ./

RUN go mod download

COPY . .
COPY .env_container .env

RUN go build -o main cmd/ordersystem/main.go

EXPOSE 8080
EXPOSE 8081
EXPOSE 50051

CMD [ "air", "-c", ".air.toml" ]

