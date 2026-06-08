package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EzraArafa/go-simple-ecommerce/internal/config"
	"github.com/EzraArafa/go-simple-ecommerce/internal/handlers"
	"github.com/EzraArafa/go-simple-ecommerce/internal/repository"
	"github.com/EzraArafa/go-simple-ecommerce/internal/services"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	cartRepo := repository.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]string{"status": "OK", "message": "E-Commerce API is  running smoothly"}
		json.NewEncoder(w).Encode(response)
	})

	router.HandleFunc("POST /products", productHandler.CreateProduct)
	router.HandleFunc("GET /products", productHandler.GetAllProducts)
	router.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	router.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	router.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	router.HandleFunc("POST /cart", cartHandler.AddToCart)
	router.HandleFunc("GET /cart", cartHandler.GetCartItems)
	router.HandleFunc("DELETE /cart/{id}", cartHandler.DeleteCartItem)

	port := ":8080"
	log.Printf("Server berhasil berjalan di port %s\n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
