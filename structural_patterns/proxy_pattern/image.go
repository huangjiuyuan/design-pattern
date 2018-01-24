package image

import (
	"fmt"
)

type Image interface {
	Display()
}

type RealImage struct {
	FileName string
}

func NewRealImage(filename string) *RealImage {
	ri := &RealImage{
		FileName: filename,
	}
	ri.LoadFromDisk(filename)
	return ri
}

func (ri *RealImage) Display() {
	fmt.Printf("Displaying %s\n", ri.FileName)
}

func (ri *RealImage) LoadFromDisk(filename string) {
	fmt.Printf("Loading %s\n", filename)
}
