package pack

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordPart_ParsePlainText(t *testing.T) {
	plainText := []byte("\x31\x7C\xE5\xAE\x89\xE5\xBE\xBD\x7C\xE5\xB7\xA2\xE6\xB9\x96\x7C\x32\x33\x38\x30\x30\x30\x7C\x30\x35\x35\x31\x0A\x32\x7C\xE5\xAE\x89\xE5\xBE\xBD\x7C\xE5\x90\x88\xE8\x82\xA5\x7C\x32\x33\x30\x30\x30\x30\x7C\x30\x35\x35\x31\x0A")
	recordPart := NewRecordPart()
	assert.NoError(t, recordPart.ParsePlainText(bytes.NewReader(plainText)))
	expectedMap := map[RecordID]*RecordItem{
		1: {
			province: "\xE5\xAE\x89\xE5\xBE\xBD",
			city:     "\xE5\xB7\xA2\xE6\xB9\x96",
			zipCode:  "238000",
			areaCode: "0551",
		},
		2: {
			province: "\xE5\xAE\x89\xE5\xBE\xBD",
			city:     "\xE5\x90\x88\xE8\x82\xA5",
			zipCode:  "230000",
			areaCode: "0551",
		},
	}
	assert.Equal(t, expectedMap, recordPart.id2item)
}

func TestRecordPart_Bytes(t *testing.T) {
	recordPart := &RecordPart{id2item: map[RecordID]*RecordItem{
		1: {
			province: "\xE5\xAE\x89\xE5\xBE\xBD",
			city:     "\xE5\xB7\xA2\xE6\xB9\x96",
			zipCode:  "238000",
			areaCode: "0551",
		},
		2: {
			province: "\xE5\xAE\x89\xE5\xBE\xBD",
			city:     "\xE5\x90\x88\xE8\x82\xA5",
			zipCode:  "230000",
			areaCode: "0551",
		},
	}}
	buf, id2offset := recordPart.Bytes(8)
	assert.Equal(t, []byte("\xE5\xAE\x89\xE5\xBE\xBD\x7C\xE5\xB7\xA2\xE6\xB9\x96\x7C\x32\x33\x38\x30\x30\x30\x7C\x30\x35\x35\x31\x00\xE5\xAE\x89\xE5\xBE\xBD\x7C\xE5\x90\x88\xE8\x82\xA5\x7C\x32\x33\x30\x30\x30\x30\x7C\x30\x35\x35\x31\x00"), buf)
	assert.Equal(t, map[RecordID]Offset{
		1: 8,
		2: 34,
	}, id2offset)
}

func TestRecordPart_Parse(t *testing.T) {
	buf := []byte("\xE5\xAE\x89\xE5\xBE\xBD\x7C\xE5\xB7\xA2\xE6\xB9\x96\x7C\x32\x33\x38\x30\x30\x30\x7C\x30\x35\x35\x31\x00\xE5\xAE\x89\xE5\xBE\xBD\x7C\xE5\x90\x88\xE8\x82\xA5\x7C\x32\x33\x30\x30\x30\x30\x7C\x30\x35\x35\x31\x00")
	recordPart := NewRecordPart()
	offset2id, err := recordPart.Parse(bytes.NewReader(buf), 8)
	assert.NoError(t, err)
	assert.Equal(t, map[Offset]RecordID{
		8:  1,
		34: 2,
	}, offset2id)
	assert.Equal(t, map[RecordID]*RecordItem{
		1: {
			province: "\xE5\xAE\x89\xE5\xBE\xBD",
			city:     "\xE5\xB7\xA2\xE6\xB9\x96",
			zipCode:  "238000",
			areaCode: "0551",
		},
		2: {
			province: "\xE5\xAE\x89\xE5\xBE\xBD",
			city:     "\xE5\x90\x88\xE8\x82\xA5",
			zipCode:  "230000",
			areaCode: "0551",
		},
	}, recordPart.id2item)
}
