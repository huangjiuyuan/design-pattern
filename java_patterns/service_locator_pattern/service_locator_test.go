package locator

import (
	"testing"
)

func TestServiceLocator(t *testing.T) {
	sl := NewServiceLocator()
	sa := sl.GetService("ServiceA")
	sa.Execute()
	sb := sl.GetService("ServiceB")
	sb.Execute()
	sa = sl.GetService("ServiceA")
	sa.Execute()
	sb = sl.GetService("ServiceB")
	sb.Execute()
}
