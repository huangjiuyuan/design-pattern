package singleton

import (
	"fmt"
)

type object struct {
	message string
}

var instance = &object{
	message: "Hello World",
}

func GetInstance() *object {
	return instance
}

func (o *object) ShowMessage() {
	fmt.Println(o.message)
}
