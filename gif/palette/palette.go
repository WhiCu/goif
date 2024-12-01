package palette

import "fmt"

type Palette struct {
	Colors []Color
}

func New(colors []Color) *Palette {
	return &Palette{
		Colors: colors,
	}
}

func (p *Palette) Bytes() []byte {
	var data []byte

	for _, value := range p.Colors {
		data = append(data, value.R, value.G, value.B)
	}

	return data
}

func (p *Palette) String() string {
	if p == nil {
		return ""
	}

	var result string

	for _, value := range p.Colors {
		result += value.String() + " "
	}

	return result[:len(result)-1]
}

type Color struct {
	R, G, B byte
}

func NewColor(r, g, b byte) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func (c *Color) String() string {
	return fmt.Sprintf("%02X %02X %02X", c.R, c.G, c.B)
}
