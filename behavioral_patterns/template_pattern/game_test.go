package game

import (
	"testing"
)

func TestGame(t *testing.T) {
	cricket := NewCricket()
	cricket.Play()

	football := NewFootball()
	football.Play()
}
