package handlers

import (
	"demo/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//	201: noContent

// DeleteProduct deletes the product from a database

func (p Products) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("")

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(writer, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, "Product Not Found", http.StatusInternalServerError)
		return
	}
}