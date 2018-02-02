package view

import (
	"fmt"
)

type FrontController struct {
	dispatcher *Dispatcher
}

func NewFrontController() *FrontController {
	return &FrontController{
		dispatcher: new(Dispatcher),
	}
}

func (fc *FrontController) IsAuthenticUser() bool {
	fmt.Println("User is authenticated successfully")
	return true
}

func (fc *FrontController) TrackRequest(req string) {
	fmt.Printf("Page requested: %s\n", req)
}

func (fc *FrontController) DispatchRequest(req string) {
	fc.TrackRequest(req)
	if fc.IsAuthenticUser() {
		fc.dispatcher.Dispatch(req)
	}
}
