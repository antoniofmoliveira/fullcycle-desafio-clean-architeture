package usecase

import (
	"github.com/antoniofmoliveira/cleanarch/internal/entity"
	"github.com/antoniofmoliveira/cleanarch/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}
	l.EventDispatcher.Dispatch(l.OrderListed)

	dto := make([]OrderOutputDTO, len(orders))
	for i, order := range orders {
		dto[i] = OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}
	return dto, nil
}
