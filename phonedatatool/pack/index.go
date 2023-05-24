package pack

import (
	"bytes"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool"
	"github.com/xluohome/phonedata/phonedatatool/util"
	"strconv"
	"strings"
)

type NumberPrefix int32
type NumberPrefixList []NumberPrefix

func (pl NumberPrefixList) Len() int {
	return len(pl)
}
func (pl NumberPrefixList) Less(i, j int) bool {
	return pl[i] < pl[j]
}
func (pl NumberPrefixList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

type IndexItem struct {
	recordOffset Offset
	cardType     phonedatatool.CardTypeID
}

type IndexPart struct {
	prefix2item map[NumberPrefix]*IndexItem
}

func NewIndexPart() *IndexPart {
	return &IndexPart{
		prefix2item: make(map[NumberPrefix]*IndexItem),
	}
}

func (p *IndexPart) ParsePlainText(reader *bytes.Reader, id2offset map[RecordID]Offset) error {
	for reader.Len() > 0 {
		var words []string
		if buf, err := util.ReadUntil(reader, '\n'); err != nil {
			return err
		} else {
			words = strings.Split(string(buf), "|")
		}

		if len(words) != 3 {
			return fmt.Errorf("expect words len is 3, got %v, %v", len(words), words)
		}

		var numberPrefix NumberPrefix
		if v, err := strconv.Atoi(words[0]); err != nil {
			return fmt.Errorf("invalid number prefix %v: %v", words[0], err)
		} else {
			numberPrefix = NumberPrefix(v)
		}

		var recordOffset Offset
		if v, err := strconv.Atoi(words[1]); err != nil {
			return fmt.Errorf("invalid record id %v: %v", words[1], err)
		} else if offset, ok := id2offset[RecordID(v)]; !ok {
			return fmt.Errorf("no offset for record id %v", v)
		} else {
			recordOffset = offset
		}

		var cardTypeID phonedatatool.CardTypeID
		if v, err := strconv.Atoi(words[2]); err != nil {
			return fmt.Errorf("invalid card type id %v: %v", words[2], err)
		} else {
			cardTypeID = phonedatatool.CardTypeID(v)
		}

		p.prefix2item[numberPrefix] = &IndexItem{
			recordOffset: recordOffset,
			cardType:     cardTypeID,
		}
	}
	return nil
}

func (p *IndexPart) Bytes() []byte {
	return nil
}
