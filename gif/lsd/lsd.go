package lsd

import "fmt"

type LSD struct {
	Width           [2]byte
	Height          [2]byte
	Flags           byte
	BackgroundIndex byte
	AspectRatio     byte
}

func New(width, height uint16, flags, backgroundIndex, aspectRatio byte) *LSD {
	return &LSD{
		Width:           [2]byte{byte(width), byte(width >> 8)},
		Height:          [2]byte{byte(width), byte(width >> 8)},
		Flags:           flags,
		BackgroundIndex: backgroundIndex,
		AspectRatio:     aspectRatio,
	}
}

func (l *LSD) Bytes() []byte {
	var data []byte

	data = append(data, l.Width[:]...)
	data = append(data, l.Height[:]...)
	data = append(data, l.Flags, l.BackgroundIndex, l.AspectRatio)

	return data
}

func (l *LSD) String() string {
	if l == nil {
		return ""
	}

	var result string

	for _, value := range l.Width {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range l.Height {
		result += fmt.Sprintf("%02X ", value)
	}

	result += fmt.Sprintf("%02X %02X %02X", l.Flags, l.BackgroundIndex, l.AspectRatio)

	return result
}
