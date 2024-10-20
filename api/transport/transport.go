// transport/http.go
package transport

import (
	"context"
	"encoding/json"

	"net/http"
	"source/endpoint"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/api/v1/product/{id}").Handler(kithttp.NewServer(
		endpoints.GetProductEndpoint,
		decodeGetProductRequest,
		encodeResponse,
	))

	// Tüm kullanıcıları getiren yeni route
	r.Methods("GET").Path("/api/v1/product").Handler(kithttp.NewServer(
		endpoints.GetAllProductsEndpoint,
		decodeGetAllProductsRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/api/v1/product").Handler(kithttp.NewServer(
		endpoints.CreateProductEndpoint,
		decodeCreateProductRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/api/v1/product/{id}").Handler(kithttp.NewServer(
		endpoints.UpdateProductEndpoint,
		decodeUpdateProductRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/api/v1/product/{id}").Handler(kithttp.NewServer(
		endpoints.DeleteProductEndpoint,
		decodeDeleteProductRequest,
		encodeResponse,
	))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	return handler
}

func decodeGetProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return endpoint.GetProductRequest{ID: int(id)}, nil
}

func decodeGetAllProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil // Çünkü herhangi bir parametre göndermiyoruz
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Product); err != nil {

		return nil, err
	}

	return req, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateProductRequest
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	req.ID = int(id)
	if err := json.NewDecoder(r.Body).Decode(&req.Product); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return endpoint.DeleteProductRequest{ID: int(id)}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(w).Encode(response)

}
