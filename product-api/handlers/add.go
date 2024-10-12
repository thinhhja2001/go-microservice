package handlers

import (
	"demo/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a product
// responses:
// 200: createProductResponse

// AddProduct returns the products from the data store
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	rw.Header().Set("Content-Type", "application/json")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProducts(&prod)

	p.l.Printf("Prod: %#v", prod)
}
