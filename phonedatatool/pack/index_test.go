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
			numberPrefix: 1300000,
			recordOffset: 300,
			cardTypeID:   2,
		},
		1300001: {
			numberPrefix: 1300001,
			recordOffset: 200,
			cardTypeID:   2,
		},
		1300002: {
			numberPrefix: 1300002,
			recordOffset: 100,
			cardTypeID:   2,
		},
	}, indexPart.prefix2item)
}

func TestIndexPart_Bytes(t *testing.T) {
	var indexPart = &IndexPart{prefix2item: map[NumberPrefix]*IndexItem{
		1300000: {
			numberPrefix: 1300000,
			recordOffset: 0x1A4E,
			cardTypeID:   2,
		},
		1300001: {
			numberPrefix: 1300001,
			recordOffset: 0x122C,
			cardTypeID:   2,
		},
		1300002: {
			numberPrefix: 1300002,
			recordOffset: 0x0008,
			cardTypeID:   2,
		},
	}}
	assert.Equal(t, []byte("\x20\xD6\x13\x00\x4E\x1A\x00\x00\x02\x21\xD6\x13\x00\x2C\x12\x00\x00\x02\x22\xD6\x13\x00\x08\x00\x00\x00\x02"), indexPart.Bytes())

}
