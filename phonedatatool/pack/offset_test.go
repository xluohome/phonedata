package pack

import (
	"bytes"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestOffset_Bytes(t *testing.T) {
	assert.Equal(t, []byte{0xd2, 0x04, 0x00, 0x00}, Offset(1234).Bytes())
}

func TestOffset_Parse(t *testing.T) {
	offset := new(Offset)
	assert.NoError(t, offset.Parse(bytes.NewReader([]byte{0xd2, 0x04, 0x00, 0x00})))
	assert.Equal(t, Offset(1234), *offset)
}
