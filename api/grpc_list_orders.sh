
grpcurl -plaintext -d '{"id": "1", "price": 100, "tax": 10}' 172.18.0.4:50051 pb.OrderService/CreateOrder

grpcurl -plaintext 172.18.0.4:50051 pb.ListOrderService/ListOrders

