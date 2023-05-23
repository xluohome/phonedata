package unpack

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type IndexPartOffsetPart struct {
	IndexPartOffset int64
}

func (op *IndexPartOffsetPart) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 4)
	if _, err := reader.Read(buf); err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	op.IndexPartOffset = int64(binary.LittleEndian.Uint64(buf))
	return nil
}
