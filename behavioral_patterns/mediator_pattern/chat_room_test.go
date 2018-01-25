package mediator

import (
	"testing"
)

func TestChatRoom(t *testing.T) {
	cr := &ChatRoom{}
	robert := NewUser(cr, "Robert")
	john := NewUser(cr, "John")

	robert.SendMessage("I am Robert.")
	john.SendMessage("I am John.")
}
