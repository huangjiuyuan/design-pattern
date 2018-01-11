package player

import (
	"testing"
)

func TestMediaAdapter(t *testing.T) {
	ap := new(AudioPlayer)
	ap.Play("mp3", "You Belong with Me.mp3")
	ap.Play("mp4", "I Knew You Were Trouble.mp4")
	ap.Play("vlc", "Welcome to New York.vlc")
	ap.Play("avi", "Look What You Made Me Do.avi")
}
