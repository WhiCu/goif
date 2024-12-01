package gce

import "fmt"

type GCE struct {
	Start          byte
	ExtensionCodes byte
	ByteCount      byte
	Data           []byte
	End            byte
}

func New(extensionCodes, byteCount byte, bytes []byte) *GCE {
	switch extensionCodes {
	case 0xFF:
		if int(byteCount)+int(bytes[byteCount]) != (len(bytes) - 1) {
			panic(fmt.Sprintf("Invalid GCE %02X: int(byteCount){%d}+int(bytes[byteCount]){%d}=%d, len(bytes)-1=%d", extensionCodes, int(byteCount), int(bytes[byteCount]), int(byteCount)+int(bytes[byteCount]), len(bytes)-1))
		}
		return &GCE{
			Start:          0x21,
			ExtensionCodes: extensionCodes,
			ByteCount:      byteCount,
			Data:           bytes,
			End:            0x00,
		}
	case 0xF9:
		if int(byteCount) != len(bytes) {
			panic(fmt.Sprintf("Invalid GCE %02X: int(byteCount)=%d, len(bytes)=%d", extensionCodes, int(byteCount), len(bytes)))
		}
		return &GCE{
			Start:          0x21,
			ExtensionCodes: extensionCodes,
			ByteCount:      byteCount,
			Data:           bytes,
			End:            0x00,
		}
	default:
		panic(fmt.Sprintf("Invalid GCE %02X: NOT FOUND GCE", extensionCodes))

	}
}

func (g *GCE) Bytes() []byte {
	var data []byte

	data = append(data, g.Start, g.ExtensionCodes, g.ByteCount)
	data = append(data, g.Data...)
	data = append(data, g.End)

	return data
}

func (g *GCE) String() string {
	var result string

	result += fmt.Sprintf("%02X %02X %02X ", g.Start, g.ExtensionCodes, g.ByteCount)

	for _, value := range g.Data {
		result += fmt.Sprintf("%02X ", value)
	}

	result += fmt.Sprintf("%02X", g.End)

	return result
}
