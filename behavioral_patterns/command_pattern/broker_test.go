package stock

import (
	"testing"
)

func TestBroker(t *testing.T) {
	alibaba := NewBuyStock("Alibaba", 100)
	nvidia := NewBuyStock("Nvidia", 50)
	apple := NewSellStock("Apple", 80)

	broker := &Broker{}
	broker.TakeOrder(alibaba)
	broker.TakeOrder(nvidia)
	broker.TakeOrder(apple)
	broker.PlacePrders()
}
