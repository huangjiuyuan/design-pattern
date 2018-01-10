package builder

import (
	"fmt"
)

type Meal struct {
	Items []Item
}

func (m *Meal) AddItem(i Item) {
	m.Items = append(m.Items, i)
}

func (m *Meal) GetCost() float64 {
	var cost float64
	for _, item := range m.Items {
		cost += item.Price()
	}
	return cost
}

func (m *Meal) ShowItems() {
	for _, item := range m.Items {
		fmt.Printf("Item: %s, Packing: %s, Price: %f\n", item.Name(), item.Packing().Pack(), item.Price())
	}
}

type MealBuilder struct{}

func (mb *MealBuilder) PrepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(VegBurger))
	meal.AddItem(new(Coke))
	return meal
}

func (mb *MealBuilder) PrepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(ChickenBurger))
	meal.AddItem(new(Pepsi))
	return meal
}
