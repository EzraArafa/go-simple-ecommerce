package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EzraArafa/go-simple-ecommerce/internal/models"
	"github.com/EzraArafa/go-simple-ecommerce/internal/services"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var item models.CartItem

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	err = h.service.AddToCart(item)
	if err != nil {
		if err.Error() == "produk tidak ditemukan" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err.Error() == "jumlah barang minimal 1" || err.Error() == " ID produk tidak valid" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{"messege": "Barang berhasil ditambahkan ke keranjang"}
	json.NewEncoder(w).Encode(response)
}
