package composite

import (
	"testing"
)

func TestCompositeEntity(t *testing.T) {
	c := NewClient()
	c.SetData("First", "Hello, it's me")
	c.PrintData()
	c.SetData("Second", "Hello from the other side")
	c.PrintData()
}
