//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/antoniofmoliveira/cleanarch/internal/entity"
	"github.com/antoniofmoliveira/cleanarch/internal/event"
	"github.com/antoniofmoliveira/cleanarch/internal/infra/database"
	"github.com/antoniofmoliveira/cleanarch/internal/infra/web"
	"github.com/antoniofmoliveira/cleanarch/internal/usecase"
	"github.com/antoniofmoliveira/cleanarch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),

	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderListedEvent = wire.NewSet(
	event.NewOrderListed,
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewWebOrderListHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebListOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		web.NewWebListOrderHandler,
	)
	return &web.WebListOrderHandler{}
}
