package endpoint

import (
	"context"

	"source/model"
	"source/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetProductEndpoint     endpoint.Endpoint
	GetAllProductsEndpoint endpoint.Endpoint
	CreateProductEndpoint  endpoint.Endpoint
	UpdateProductEndpoint  endpoint.Endpoint
	DeleteProductEndpoint  endpoint.Endpoint
}

func MakeEndpoints(s service.ProductService) Endpoints {
	return Endpoints{
		GetProductEndpoint:     makeGetProductEndpoint(s),
		GetAllProductsEndpoint: makeGetAllProductsEndpoint(s),
		CreateProductEndpoint:  makeCreateProductEndpoint(s),
		UpdateProductEndpoint:  makeUpdateProductEndpoint(s),
		DeleteProductEndpoint:  makeDeleteProductEndpoint(s),
	}
}

func makeGetProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		return s.GetProduct(req.ID)
	}
}

func makeGetAllProductsEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetAllProducts()
		if err != nil {
			return nil, err
		}
		return products, nil
	}
}

func makeCreateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)

		return s.CreateProduct(req.Product)
	}
}

func makeUpdateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProductRequest)
		return s.UpdateProduct(req.ID, req.Product)
	}
}

func makeDeleteProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		return nil, s.DeleteProduct(req.ID)
	}
}

// Request and Response structs
type GetProductRequest struct {
	ID int `json:"id"`
}

type CreateProductRequest struct {
	Product model.Tb_casestudy `json:"product"`
}

type UpdateProductRequest struct {
	ID      int                `json:"id"`
	Product model.Tb_casestudy `json:"product"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}
