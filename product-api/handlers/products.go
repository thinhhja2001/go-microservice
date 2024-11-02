package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"example.com/hello/data"
	protos "github.com/thinhhja2001/currency/protos"
)

// Products is a http.Handler
type Products struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{l, cc}
}

type KeyProduct struct{}

func (p Products) ProductMiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
