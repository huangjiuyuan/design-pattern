package composite

import (
	"fmt"
)

type Client struct {
	ce *CompositeEntity
}

func NewClient() *Client {
	return &Client{
		ce: NewCompositeEntity(),
	}
}

func (c *Client) PrintData() {
	for _, data := range c.ce.GetData() {
		fmt.Printf("Data: %s\n", data)
	}
}

func (c *Client) SetData(a, b string) {
	c.ce.SetData(a, b)
}
