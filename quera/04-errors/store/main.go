package main

import (
	"fmt"
	"sort"
)

type Store struct {
	stocks map[string]int
	prices map[string]float64
}

func NewStore() *Store {
	return &Store{
		stocks: make(map[string]int),
		prices: make(map[string]float64),
	}
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	if _, found := s.stocks[name]; found {
		return fmt.Errorf("%v already exists", name)
	}

	if price <= 0 {
		return fmt.Errorf("price should be positive")
	}

	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	s.stocks[name] = count
	s.prices[name] = price
	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	if count, found := s.stocks[name]; found {
		return count, nil
	} else {
		return 0, fmt.Errorf("invalid product name")
	}
}

func (s *Store) GetProductPrice(name string) (float64, error) {
	if price, found := s.prices[name]; found {
		return price, nil
	} else {
		return 0, fmt.Errorf("invalid product name")
	}
}

func (s *Store) Order(name string, count int) error {
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	stock, found := s.stocks[name]

	if !found {
		return fmt.Errorf("invalid product name")
	}

	if stock == 0 {
		return fmt.Errorf("there is no %v in the store", name)
	}

	if stock < count {
		return fmt.Errorf("not enough %v in the store. there are %v left", name, stock)
	}

	s.stocks[name] -= count
	return nil
}

func (s *Store) ProductsList() ([]string, error) {
	if len(s.stocks) == 0 {
		return nil, fmt.Errorf("store is empty")
	}

	result := make([]string, 0)
	for product, stock := range s.stocks {
		if stock == 0 {
			continue
		}
		result = append(result, product)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("store is empty")
	}

	sort.Strings(result)

	return result, nil
}
