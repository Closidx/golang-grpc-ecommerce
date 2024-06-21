package cart

import (
	"context"
	ecommercev1 "github/closidx/golang-grpc-ecommerce/protos-contract/gen/go/ecommerce"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ecommercev1.UnimplementedCartServiceServer
}

func Register(gRPC *grpc.Server) {
	ecommercev1.RegisterCartServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) AddToCart(
	ctx context.Context,
	req *ecommercev1.AddToCartRequest,
) (*ecommercev1.AddToCartResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetCart(
	ctx context.Context,
	req *ecommercev1.GetCartRequest,
) (*ecommercev1.GetCartResponse, error) {
	panic("implement me")
}

func (s *serverAPI) RemoveFromCart(
	ctx context.Context,
	req *ecommercev1.RemoveFromCartRequest,
) (*ecommercev1.RemoveFromCartResponse, error) {
	panic("implement me")
}
