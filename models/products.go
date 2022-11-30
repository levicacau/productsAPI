package models

import (
	"errors"
	"fmt"
	"net/http"
)

type Product struct {
	Codigobarras  string  `json:"barcode"`
	Precogondola1 float64 `json:"price"`
	Descricao     string  `json:"description"`
	Unidade       string  `json:"unit"`
}
type ProductList struct {
	Products []Product `json:"products"`
}

func (i *Product) Bind(r *http.Request) error {
	if i.Descricao == "" {
		return fmt.Errorf("a descrição é obrigatória")
	}
	return nil
}
func (*ProductList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Product) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewProduct(barcode string, price float64, description string, unit string) (*Product, error) {
	product := &Product{
		Codigobarras:  barcode,
		Precogondola1: price,
		Descricao:     description,
		Unidade:       unit,
	}
	err := product.IsValid()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (o *Product) IsValid() error {
	if o.Codigobarras == "" {
		return errors.New("invalid barcode")
	}
	if o.Precogondola1 <= 0 {
		return errors.New("invalid price")
	}
	if o.Descricao == "" {
		return errors.New("invalid description")
	}
	if o.Unidade == "" {
		return errors.New("invalid unit")
	}
	return nil
}
