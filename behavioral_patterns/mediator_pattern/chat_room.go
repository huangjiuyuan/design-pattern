package mediator

import (
	"fmt"
)

type ChatRoom struct{}

func (cr *ChatRoom) ShowMessage(user *User, message string) {
	fmt.Printf("User %s said: %s\n", user.GetName(), message)
}

type User struct {
	ChatRoom *ChatRoom
	Name     string
}

func NewUser(cr *ChatRoom, name string) *User {
	return &User{
		ChatRoom: cr,
		Name:     name,
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SendMessage(message string) {
	u.ChatRoom.ShowMessage(u, message)
}
