package order

import (
	"context"
	ecommercev1 "github/closidx/golang-grpc-ecommerce/protos-contract/gen/go/ecommerce"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ecommercev1.UnimplementedOrderServiceServer
}

func Register(gRPC *grpc.Server) {
	ecommercev1.RegisterOrderServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) CreateOrder(
	ctx context.Context,
	req *ecommercev1.CreateOrderRequest,
) (*ecommercev1.CreateOrderResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetOrder(
	ctx context.Context,
	req *ecommercev1.GetOrderRequest,
) (*ecommercev1.GetOrderResponse, error) {
	panic("implement me")
}

func (s *serverAPI) UpdateOrder(
	ctx context.Context,
	req *ecommercev1.UpdateOrderRequest,
) (*ecommercev1.UpdateOrderResponse, error) {
	panic("implement me")
}

func (s *serverAPI) CancelOrder(
	ctx context.Context,
	req *ecommercev1.CancelOrderRequest,
) (*ecommercev1.CancelOrderResponse, error) {
	panic("implement me")
}