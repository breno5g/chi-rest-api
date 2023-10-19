package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("test", 10.0)
	assert.Nil(t, err)

	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	_, err := NewProduct("", 10.0)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("test", 0)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	_, err := NewProduct("test", -10.0)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("test", 10.0)
	assert.Nil(t, err)

	assert.Nil(t, product.Validate())
}
