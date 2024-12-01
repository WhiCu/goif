package gif

import (
	"bytes"
	"fmt"
	"goif/gif/gce"
	"goif/gif/header"
	"goif/gif/image"
	"goif/gif/lsd"
	"goif/gif/palette"
	"goif/gif/trailer"
	"io"
	"os"
)

type GIF struct {
	Header  *header.Header
	LSD     *lsd.LSD
	Palette *palette.Palette
	//Extension []*gce.GCE
	ImageData []*image.ImageData
	Trailer   *trailer.Trailer
}

func New() *GIF {
	return &GIF{}
}
func (g *GIF) String() string {
	var result string

	result += fmt.Sprintf("%s %s %s ", g.Header, g.LSD, g.Palette)

	for _, value := range g.ImageData {
		result += value.String() + " "
	}

	result += g.Trailer.String()

	return result
}

func (g *GIF) Bytes() []byte {
	var data []byte

	data = append(data, g.Header.Bytes()...)
	data = append(data, g.LSD.Bytes()...)
	data = append(data, g.Palette.Bytes()...)

	for _, value := range g.ImageData {
		data = append(data, value.Bytes()...)
	}

	data = append(data, g.Trailer.Bytes()...)

	return data
}

func U16toBytes(u uint16) []byte {
	var b [2]byte

	b[0] = byte(u)
	b[1] = byte(u >> 8)
	return b[:]
}

func BytesToU16(b []byte) uint16 {
	return uint16(b[1])<<8 | uint16(b[0])
	//return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func GIFFromFile(file *os.File) *GIF {
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var gif GIF

	if (bytes.Equal(data[:6], header.StandardNineBytes()) || bytes.Equal(data[:6], header.StandardSevenBytes())) && data[len(data)-1] == trailer.StandardByte() {
		gif = GIF{
			Header:  header.New(),
			LSD:     lsd.New(BytesToU16(data[6:8]), BytesToU16(data[8:10]), data[10], data[11], data[12]),
			Trailer: trailer.New(),
		}
		pallette := &palette.Palette{}
		position := 13
		if data[10]&0x80 == 0x80 {
			position = 13 + (2<<(data[10]&0x07))*3
			colorsBytes := data[13:position]

			for i := 0; i < 2<<(data[10]&0x07); i++ {
				color := palette.NewColor(colorsBytes[3*i], colorsBytes[3*i+1], colorsBytes[3*i+2])
				pallette.Colors = append(pallette.Colors, color)
			}
		}
		gif.Palette = pallette

		images := []*image.ImageData{}

		// position = 45
		// size := data[position+2]
		// fmt.Printf("%02X %d - ", data[position+1], size)
		// for _, value := range data[position+3 : position+3+int(size)+1+int(data[position+3+int(size)])] {
		// 	fmt.Printf("%02X ", value)
		// }

		counter := 0
		for data[position] != 0x3B && counter < 10 {

			var Extensions []*gce.GCE

			for data[position] == 0x21 {
				var g *gce.GCE

				switch data[position+1] {
				case 0xFF:
					size := data[position+2]
					g = gce.New(data[position+1], size, data[position+3:position+3+int(size)+1+int(data[position+3+int(size)])])
					position += 3 + int(size) + 1 + int(data[position+3+int(size)]) + 1
				case 0xF9:
					size := data[position+2]
					g = gce.New(data[position+1], size, data[position+3:position+3+int(size)])
					position += int(size) + 3 + 1
				default:
					panic(fmt.Sprintf("Invalid GCE %02X: NOT FOUND GCE", data[position+1]))
				}

				Extensions = append(Extensions, g)

			}

			if data[position] != 0x2C {
				panic("Invalid GIF: NOT FOUND 0x2C")
			}
			img := image.New(
				Extensions,
				BytesToU16(data[position+1:position+3]),
				BytesToU16(data[position+3:position+5]),
				BytesToU16(data[position+5:position+7]),
				BytesToU16(data[position+7:position+9]),
				data[position+9],
				data[position+10],
				data[position+11],
				data[position+12:position+12+int(data[position+11])],
			)

			position += 12 + int(data[position+11]) + 1

			images = append(images, img)
			counter++
		}

		gif.ImageData = images
		//gif.Extension = gce.GCEsFromBytes(data[16+BytesToU16(data[14:16])*3:])

	}
	return &gif
}

// data, err := io.ReadAll(file)
// if err != nil {
// 	panic(err)
// }

// var png PNG

// if bytes.Equal(data[:8], header.New().StandardBytes()) && bytes.Equal(data[len(data)-12:], chunk.StandardIEND()) {
// 	png = PNG{
// 		Header: header.New(),
// 		IHDR:   chunk.IHDRFromBytes(data[8:33]),
// 		IDAT:   chunk.ChunksFromBytes(data[33 : len(data)-12]),
// 		IEND:   chunk.NewIEND(),
// 	}
// } else {
// 	panic(fmt.Sprintf("File %s is not a PNG file", file.Name()))
// }

// return &png
