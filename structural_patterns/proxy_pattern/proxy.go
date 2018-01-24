package image

type ProxyImage struct {
	RealImage *RealImage
	FileName  string
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{
		FileName: filename,
	}
}

func (pi *ProxyImage) Display() {
	if pi.RealImage == nil {
		pi.RealImage = NewRealImage(pi.FileName)
	}
	pi.RealImage.Display()
}
