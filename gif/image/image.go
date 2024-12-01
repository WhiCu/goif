package image

import (
	"fmt"
	"goif/gif/gce"
)

type ImageData struct {
	Extensions []*gce.GCE
	Start      byte
	StartX     [2]byte
	StartY     [2]byte
	Width      [2]byte
	Height     [2]byte
	Flags      byte
	LZWMinSize byte
	SizeData   byte
	Data       []byte
	End        byte
}

func New(Extensions []*gce.GCE, startX, startY uint16, width, height uint16, Flags, LZWMinSize, SizeData byte, data []byte) *ImageData {
	return &ImageData{
		Start:      0x2C,
		Extensions: Extensions,
		StartX:     [2]byte{byte(startX), byte(startX >> 8)},
		StartY:     [2]byte{byte(startY), byte(startY >> 8)},
		Width:      [2]byte{byte(width), byte(width >> 8)},
		Height:     [2]byte{byte(width), byte(width >> 8)},
		Flags:      Flags,
		LZWMinSize: LZWMinSize,
		SizeData:   SizeData,
		Data:       data,
		End:        0x00,
	}
}

func (i *ImageData) Bytes() []byte {
	var data []byte

	for _, value := range i.Extensions {
		data = append(data, value.Bytes()...)
	}

	data = append(data, i.Start)
	data = append(data, i.StartX[:]...)
	data = append(data, i.StartY[:]...)
	data = append(data, i.Width[:]...)
	data = append(data, i.Height[:]...)
	data = append(data, i.Flags, i.LZWMinSize, i.SizeData)
	data = append(data, i.Data...)
	data = append(data, i.End)

	return data
}

func (i *ImageData) String() string {
	if i == nil {
		return ""
	}

	var result string

	for _, value := range i.Extensions {
		result += value.String() + " "
	}

	result += fmt.Sprintf("%02X ", i.Start)

	for _, value := range i.StartX {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range i.StartY {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range i.Width {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range i.Height {
		result += fmt.Sprintf("%02X ", value)
	}

	result += fmt.Sprintf("%02X %02X %02X ", i.Flags, i.LZWMinSize, i.SizeData)

	for _, value := range i.Data {
		result += fmt.Sprintf("%02X ", value)
	}

	result += fmt.Sprintf("%02X", i.End)

	return result
}
