package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	s := NewStore()
	err := s.AddProduct("apple", 20000, 10)
	assert.Nil(t, err)
	c, err2 := s.GetProductCount("apple")
	assert.Nil(t, err2)
	assert.Equal(t, 10, c)
	p, err3 := s.GetProductPrice("apple")
	assert.Nil(t, err3)
	assert.Equal(t, 20000.0, p)
	err = s.Order("apple", 8)
	assert.Nil(t, err)
	c, _ = s.GetProductCount("apple")
	assert.Equal(t, 2, c)
}

func TestAddProductWithNegativePrice(t *testing.T) {
	s := NewStore()
	err := s.AddProduct("a", -23, 3)
	assert.Equal(t, err.Error(), "price should be positive")
	err = s.AddProduct("b", 0, 4)
	assert.Equal(t, err.Error(), "price should be positive")
}

func TestAddProductWithNegativeCount(t *testing.T) {
	s := NewStore()
	err := s.AddProduct("a", 23, -8)
	assert.Equal(t, err.Error(), "count should be positive")
	err = s.AddProduct("b", 3, 0)
	assert.Equal(t, err.Error(), "count should be positive")
}

func TestAddProductAlreadyExistProduct(t *testing.T) {
	name := "a"
	s := NewStore()
	err := s.AddProduct(name, 1, 2)
	assert.Nil(t, err)
	err = s.AddProduct(name, 3, 4)
	assert.Equal(t, err.Error(), fmt.Sprintf("%v already exists", name))
}

func TestGetProductCountCheckCorrectStock(t *testing.T) {
	stock := 2
	s := NewStore()
	s.AddProduct("a", 1, stock)
	res, err := s.GetProductCount("a")
	assert.Nil(t, err)
	assert.Equal(t, res, stock)
}

func TestGetProductCountIncorrectName(t *testing.T) {
	s := NewStore()
	s.AddProduct("a", 1, 2)
	res, err := s.GetProductCount("b")
	assert.Equal(t, res, 0)
	assert.Equal(t, err.Error(), "invalid product name")
}
