package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyBarcode_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{}
	assert.Error(t, product.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{Codigobarras: "123"}
	assert.Error(t, product.IsValid(), "invalid price")
}

func TestGivenAnEmptyDescription_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{Codigobarras: "123", Precogondola1: 10}
	assert.Error(t, product.IsValid(), "invalid description")
}

func TestGivenAnEmptyUnit_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{Codigobarras: "123", Precogondola1: 10, Descricao: "Rice"}
	assert.Error(t, product.IsValid(), "invalid unit")
}

func TestGivenAValidParams_WhenICallNewProduct_ThenIShouldReceiveCreateProductWithAllParams(t *testing.T) {
	product := Product{
		Codigobarras:  "123",
		Precogondola1: 10,
		Descricao:     "Rice",
		Unidade:       "UN",
	}
	assert.Equal(t, "123", product.Codigobarras)
	assert.Equal(t, 10.0, product.Precogondola1)
	assert.Equal(t, "Rice", product.Descricao)
	assert.Equal(t, "UN", product.Unidade)
	assert.Nil(t, product.IsValid())
}

func TestGivenAValidParams_WhenICallNewProductFunc_ThenIShouldReceiveCreateProductWithAllParams(t *testing.T) {
	product, err := NewProduct("123", 10.0, "Rice", "UN")
	assert.Nil(t, err)
	assert.Equal(t, "123", product.Codigobarras)
	assert.Equal(t, 10.0, product.Precogondola1)
	assert.Equal(t, "Rice", product.Descricao)
	assert.Equal(t, "UN", product.Unidade)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	product, err := NewProduct("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, product.CalculateFinalPrice())
	assert.Equal(t, 12.0, product.FinalPrice)
}
