package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "Rex",
		Price: 20,
		SKU: "abcaskfn-aen-alskn",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal()
	}
}
