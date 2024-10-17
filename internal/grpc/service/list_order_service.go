package service

import (
	"context"

	"github.com/antoniofmoliveira/cleanarch/internal/grpc/pb"
	"github.com/antoniofmoliveira/cleanarch/internal/usecase"
)

type ListOrderService struct {
	pb.UnimplementedListOrderServiceServer
	ListOrderUseCase usecase.ListOrderUseCase
}

func NewListOrderService(listOrderUseCase usecase.ListOrderUseCase) *ListOrderService {
	return &ListOrderService{
		ListOrderUseCase: listOrderUseCase,
	}
}

func (s *ListOrderService) ListOrders(ctx context.Context, in *pb.Empty) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	listOrdersResponse := make([]*pb.OrderResponse, len(output))
	for i, order := range output {
		listOrdersResponse[i] = &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
	}
	return &pb.ListOrdersResponse{
		Orders: listOrdersResponse,
	}, nil
}
