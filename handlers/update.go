package handlers

import (
	"demo/data"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route PUT /products/{id} products updateProduct
// Returns a list of products
// responses:
//	200: noContent

// UpdateProduct update the product from a database
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Update Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, prod)
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

}
