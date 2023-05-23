package unpack

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type IndexPartOffsetPart struct {
	IndexPartOffset RecordOffset
}

func (op *IndexPartOffsetPart) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 4)
	if _, err := reader.Read(buf); err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	op.IndexPartOffset = RecordOffset(binary.LittleEndian.Uint32(buf))
	return nil
}
