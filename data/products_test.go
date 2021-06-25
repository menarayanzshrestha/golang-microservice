package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "RAm",
		Price: 3,
		SKU:   "abc-dd-df",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
