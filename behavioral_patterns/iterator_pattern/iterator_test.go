package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	nr := &NameRepository{
		Names: []string{"Robert", "John", "Julie", "Lora"},
	}

	iter := nr.GetIterator()
	for iter.HasNext() {
		name := iter.Next().(string)
		fmt.Println(name)
	}
}
