package player

import (
	"fmt"
)

type AdvancedMediaPlayer interface {
	PlayVlc(filename string)
	PlayMp4(filename string)
}

type VlcPlayer struct{}

func (vlc *VlcPlayer) PlayVlc(filename string) {
	fmt.Printf("Playing vlc file: %s\n", filename)
}

func (vlc *VlcPlayer) PlayMp4(filename string) {
	fmt.Printf("Not working\n")
}

type Mp4Player struct{}

func (mp4 *Mp4Player) PlayVlc(filename string) {
	fmt.Printf("Not working\n")
}

func (mp4 *Mp4Player) PlayMp4(filename string) {
	fmt.Printf("Playing mp4 file: %s\n", filename)
}
