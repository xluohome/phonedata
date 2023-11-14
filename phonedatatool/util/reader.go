package util

import "bytes"

// ReadUntil 读取，直到遇到特定字符
func ReadUntil(reader *bytes.Reader, term byte) ([]byte, error) {
	var buf []byte
	for {
		if b, err := reader.ReadByte(); err != nil {
			return nil, err
		} else if b == term {
			return buf, nil
		} else {
			buf = append(buf, b)
		}
	}
}
