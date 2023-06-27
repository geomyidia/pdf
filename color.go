package fpdf

var (
	Black = RGB{0, 0, 0}
)

type RGB struct {
	R int
	G int
	B int
}

func (rgb *RGB) Red() int {
	return rgb.R
}

func (rgb *RGB) Green() int {
	return rgb.G
}

func (rgb *RGB) Blue() int {
	return rgb.B
}

type RGBA struct {
	RGB
	A float64
}

func (rgba *RGBA) Alpha() float64 {
	return rgba.A
}

func (rgba *RGBA) Color() RGB {
	return rgba.RGB
}
