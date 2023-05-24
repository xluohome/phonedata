package pack

import "encoding/binary"

type Offset int64

func (o Offset) Bytes() []byte {
	return binary.LittleEndian.AppendUint32(nil, uint32(o))
}
