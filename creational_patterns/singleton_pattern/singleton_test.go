package singleton

import (
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {
	defaultObj := &object{
		message: "Hello World",
	}
	obj := GetInstance()
	obj.ShowMessage()
	if !reflect.DeepEqual(defaultObj, obj) {
		t.Fatalf("expected obj: %#v, got a different: %#v", defaultObj, obj)
	}
}
