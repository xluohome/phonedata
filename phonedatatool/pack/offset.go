package pack

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Offset int64

func (o Offset) Bytes() []byte {
	return binary.LittleEndian.AppendUint32(nil, uint32(o))
}
func (o *Offset) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 4)
	if _, err := reader.Read(buf); err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	*o = Offset(binary.LittleEndian.Uint32(buf))
	return nil
}
