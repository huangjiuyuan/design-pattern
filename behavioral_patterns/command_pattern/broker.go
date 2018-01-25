package stock

import (
	"sync"
)

type Broker struct {
	sync.Mutex
	OrderList []Order
}

func (b *Broker) TakeOrder(order Order) {
	b.Lock()
	defer b.Unlock()
	b.OrderList = append(b.OrderList, order)
}

func (b *Broker) PlacePrders() {
	b.Lock()
	b.Unlock()
	for _, o := range b.OrderList {
		o.Execute()
	}
}
