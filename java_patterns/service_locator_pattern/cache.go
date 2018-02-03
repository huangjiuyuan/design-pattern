package locator

import (
	"fmt"
)

type Cache struct {
	Services []Service
}

func (c *Cache) GetServices(name string) Service {
	for _, s := range c.Services {
		if s.GetName() == name {
			fmt.Printf("Returning cached %s object", name)
			return s
		}
	}
	return nil
}

func (c *Cache) AddService(service Service) {
	var existed bool
	for _, s := range c.Services {
		if s.GetName() == service.GetName() {
			existed = true
		}
	}
	if !existed {
		c.Services = append(c.Services, service)
	}
}
