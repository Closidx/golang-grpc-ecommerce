package product

import (
	"context"
	ecommercev1 "github/closidx/golang-grpc-ecommerce/protos-contract/gen/go/ecommerce"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ecommercev1.UnimplementedProductServiceServer
}

func Register(gRPC *grpc.Server) {
	ecommercev1.RegisterProductServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) CreateProduct(
	ctx context.Context,
	req *ecommercev1.CreateProductRequest,
) (*ecommercev1.CreateProductResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetProducts(
	ctx context.Context,
	req *ecommercev1.GetProductsRequest,
) (*ecommercev1.GetProductsResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetProduct(
	ctx context.Context,
	req *ecommercev1.GetProductRequest,
) (*ecommercev1.GetProductResponse, error) {
	panic("implement me")
}

func (s *serverAPI) UpdateProduct(
	ctx context.Context,
	req *ecommercev1.UpdateProductRequest,
) (*ecommercev1.UpdateProductResponse, error) {
	panic("implement me")
}

func (s *serverAPI) DeleteProduct(
	ctx context.Context,
	req *ecommercev1.DeleteProductRequest,
) (*ecommercev1.DeleteProductResponse, error) {
	panic("implement me")
}