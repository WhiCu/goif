package header

import "fmt"

type Header struct {
	Signature [3]byte // "GIF"
	Version   [3]byte // "87a" or "89a"
}

func New() *Header {
	return &Header{
		Signature: [3]byte{'G', 'I', 'F'},
		Version:   [3]byte{'8', '9', 'a'},
	}
}

func StandardNineBytes() []byte {
	return []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}
}

func StandardSevenBytes() []byte {
	return []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61}
}

func (h *Header) Bytes() []byte {
	var data []byte

	data = append(data, h.Signature[:]...)
	data = append(data, h.Version[:]...)

	return data
}

func (h *Header) String() string {
	if h == nil {
		return ""
	}

	var result string

	for _, value := range h.Signature {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range h.Version {
		result += fmt.Sprintf("%02X ", value)
	}

	return result[:len(result)-1]
}
