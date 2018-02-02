package view

import (
	"testing"
)

func TestFrontController(t *testing.T) {
	fc := NewFrontController()
	fc.DispatchRequest("student")
	fc.DispatchRequest("home")
	fc.DispatchRequest("school")
}
