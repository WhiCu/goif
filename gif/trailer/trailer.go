package trailer

import "fmt"

type Trailer byte

func New() *Trailer {
	t := Trailer(0x3B)
	return &t
}

func StandardByte() byte {
	return 0x3B
}

func (t *Trailer) Bytes() []byte {
	return []byte{byte(*t)}
}

func (t *Trailer) String() string {
	return fmt.Sprintf("%02X", *t)
}
