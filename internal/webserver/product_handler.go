package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Alefeoliveira/imersao17/goapi/internal/entity"
	"github.com/Alefeoliveira/imersao17/goapi/internal/service"
	"github.com/go-chi/chi"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: &productService}
}

func (wch *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wch.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wch *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wch.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wch *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	products, err := wch.ProductService.GetProductByCategoryID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wch *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wch.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
