#!/bin/bash

# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

migrate -source file://. -database mysql://ordersystem:ordersystem@mariadbca:3306/ordersystem up 4