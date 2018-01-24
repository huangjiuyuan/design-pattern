package image

import (
	"testing"
)

func TestProxyImage(t *testing.T) {
	image := NewProxyImage("sexy.jpg")
	// Loading image from disk.
	image.Display()
	// Loading from proxy.
	image.Display()
}
