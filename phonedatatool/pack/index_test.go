package pack

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexPart_ParsePlainText(t *testing.T) {
	indexPart := NewIndexPart()
	reader := bytes.NewReader([]byte("1300000|251|2\n1300001|176|2\n1300002|1|2\n"))
	id2offset := map[RecordID]Offset{
		1:   100,
		176: 200,
		251: 300,
	}
	assert.NoError(t, indexPart.ParsePlainText(reader, id2offset))
	assert.Equal(t, map[NumberPrefix]*IndexItem{
		1300000: {
			recordOffset: 300,
			cardType:     2,
		},
		1300001: {
			recordOffset: 200,
			cardType:     2,
		},
		1300002: {
			recordOffset: 100,
			cardType:     2,
		},
	}, indexPart.prefix2item)
}
