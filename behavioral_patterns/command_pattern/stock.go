package stock

import (
	"fmt"
)

type Stock struct {
	Name     string
	Quantity int
}

func (s *Stock) Buy() {
	fmt.Printf("Stock [Name: %s, Quantity: %d] bought\n", s.Name, s.Quantity)
}

func (s *Stock) Sell() {
	fmt.Printf("Stock [Name: %s, Quantity: %d] sold\n", s.Name, s.Quantity)
}

type BuyStock struct {
	Stock
}

func NewBuyStock(name string, quantity int) *BuyStock {
	return &BuyStock{
		Stock{
			Name:     name,
			Quantity: quantity,
		},
	}
}

func (bs *BuyStock) Execute() {
	bs.Buy()
}

type SellStock struct {
	Stock
}

func NewSellStock(name string, quantity int) *SellStock {
	return &SellStock{
		Stock{
			Name:     name,
			Quantity: quantity,
		},
	}
}

func (ss *SellStock) Execute() {
	ss.Sell()
}
