package handlers

import (
	"context"
	"errors"
	protos "github.com/thinhhja2001/currency/protos"
	"net/http"
	"strconv"

	"example.com/hello/data"
	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	rw.Header().Set("Content-Type", "application/json")

	// fetch the products from the datastore
	lp := data.GetProducts()
	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products getSingleProduct
// Returns a single product
// responses:
//	200: createProductResponse

// GetSingleProduct returns the products from the data store
func (p Products) GetSingleProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Single Product")
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	// fetch the products from the database
	product, err := data.GetSingleProduct(id)

	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[ERROR] error getting new Rate")
		http.Error(rw, "Unable to get Rate Request", http.StatusInternalServerError)
		return
	}
	product.Price = product.Price * resp.Rate
	err = product.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}
