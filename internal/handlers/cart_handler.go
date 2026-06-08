package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (h *CartHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetCartItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(items)
}

func (h *CartHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteCartItem(id)
	if err != nil {
		if err.Error() == "item keranjang tidak ditemukan" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "Item berhasil dihapus dari keranjang"}
	json.NewEncoder(w).Encode(response)
}

func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	err := h.service.Checkout()
	if err != nil {
		if err.Error() == "keranjang kosong" || err.Error() == "stok produk tidak mencukupi untuk diproses" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content- Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "Checkout berhasil! Stok telah dikurangi dan keranjang dikosongkan."}
	json.NewEncoder(w).Encode(response)
}
