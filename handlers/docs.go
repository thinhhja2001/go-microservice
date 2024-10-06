package handlers

import "demo/data"

// A list of products returns in the response
// swagger:response productsResponse
type productResponseWrapper struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// swagger:parameters deleteProduct updateProduct
type productIDParameterWrapper struct {
	// ID of the product
	// in:path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productNoContent struct {
}

// swagger:response createProductResponse
type createProductResponseWrapper struct {
	// Return newly created product
	// in:body
	Body data.Product
}
