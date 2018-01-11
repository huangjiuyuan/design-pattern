package player

import (
	"fmt"
)

type MediaPlayer interface {
	Play(audiotype string, filename string)
}

type MediaAdapter interface {
	MediaPlayer
	AdvancedMediaPlayer
}

type AudioPlayer struct {
	Player AdvancedMediaPlayer
}

func (ap *AudioPlayer) MediaAdapter(audiotype string) {
	if audiotype == "vlc" {
		ap.Player = &VlcPlayer{}
	} else if audiotype == "mp4" {
		ap.Player = &Mp4Player{}
	}
}

func (ap *AudioPlayer) Play(audiotype string, filename string) {
	if audiotype == "mp3" {
		fmt.Printf("Playing mp3 file: %s\n", filename)
	} else if audiotype == "vlc" {
		ap.MediaAdapter(audiotype)
		ap.Player.PlayVlc(filename)
	} else if audiotype == "mp4" {
		ap.MediaAdapter(audiotype)
		ap.Player.PlayMp4(filename)
	} else {
		fmt.Printf("Invalid media: %s\n", filename)
	}
}
