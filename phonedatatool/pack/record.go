package pack

import (
	"bytes"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool/util"
	"strconv"
	"strings"
)

type RecordID int64

type RecordItem struct {
	province string
	city     string
	zipCode  string
	areaCode string
}

type RecordPart struct {
	id2item map[RecordID]*RecordItem
}

func NewRecordPart() *RecordPart {
	return &RecordPart{id2item: make(map[RecordID]*RecordItem)}
}

func (p *RecordPart) ParsePlainText(reader *bytes.Reader) error {
	for reader.Len() > 0 {

		var words []string
		if b, err := util.ReadUntil(reader, '\n'); err != nil {
			return err
		} else {
			words = strings.Split(string(b), "|")
		}
		if len(words) != 5 {
			return fmt.Errorf("invalid record line. expect 5 words (id, province, city, zipCode, areaCode), got %v words, %v", len(words), words)
		}

		var recordID RecordID
		if id, err := strconv.Atoi(words[0]); err != nil {
			return fmt.Errorf("invalid id format, raw=%v, err=%v", words[0], err)
		} else {
			recordID = RecordID(id)
		}

		if _, ok := p.id2item[recordID]; ok {
			return fmt.Errorf("duplicate recordID %v", recordID)
		}

		p.id2item[recordID] = &RecordItem{
			province: words[1],
			city:     words[2],
			zipCode:  words[3],
			areaCode: words[4],
		}
	}
	return nil
}
