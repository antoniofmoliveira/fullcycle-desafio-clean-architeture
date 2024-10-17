
grpcurl -plaintext -d '{"id": "1", "price": 100, "tax": 10}' localhost:50051 pb.OrderService/CreateOrder

grpcurl -plaintext localhost:50051 pb.ListOrderService/ListOrders

