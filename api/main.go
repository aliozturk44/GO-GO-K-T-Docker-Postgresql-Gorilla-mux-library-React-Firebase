package main

import (
	"net/http"
	"source/config"
	"source/endpoint"
	"source/service"
	"source/transport"

	"log"
)

func main() {
	// config.InitialMigration()
	// r := mux.NewRouter()
	// r.HandleFunc("/api/v1/product", controllers.GetProducts).Methods("GET")
	// r.HandleFunc("/api/v1/product/{id}", controllers.GetProductsById).Methods("GET")
	// r.HandleFunc("/api/v1/product/{id}", controllers.UpdateProduct).Methods("PUT")
	// r.HandleFunc("/api/v1/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	// r.HandleFunc("/api/v1/product", controllers.AddProduct).Methods("POST")

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// })

	// handler := c.Handler(r)
	// http.ListenAndServe(":8080", handler)
	config.InitialMigration()
	productService := service.NewProductService(config.DB)
	endpoints := endpoint.MakeEndpoints(productService)
	handler := transport.NewHTTPHandler(endpoints)

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
