package customer

import (
	"testing"
)

func TestCustomer(t *testing.T) {
	cf := &CustomerFactory{
		Names: []string{"Rob", "Joe", "Julie"},
	}

	rob := cf.GetCustomer("Rob")
	bob := cf.GetCustomer("Bob")
	julie := cf.GetCustomer("Julie")
	laura := cf.GetCustomer("Laura")

	if rob.GetName() != "Rob" {
		t.Errorf("expected Rob, got %s\n", rob.GetName())
	}
	if bob.GetName() != "Unavailable in customer database" {
		t.Errorf("expected Bob, got %s\n", bob.GetName())
	}
	if julie.GetName() != "Julie" {
		t.Errorf("expected Julie, got %s\n", julie.GetName())
	}
	if laura.GetName() != "Unavailable in customer database" {
		t.Errorf("expected Laura, got %s\n", laura.GetName())
	}
}
