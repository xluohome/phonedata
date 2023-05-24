package pack

import "testing"
import "github.com/stretchr/testify/assert"

func TestOffset_Bytes(t *testing.T) {
	assert.Equal(t, []byte{0xd2, 0x04, 0x00, 0x00}, Offset(1234).Bytes())
}
