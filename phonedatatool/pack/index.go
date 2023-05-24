package pack

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool"
	"github.com/xluohome/phonedata/phonedatatool/util"
	"sort"
	"strconv"
	"strings"
)

type NumberPrefix int32

func (p NumberPrefix) Bytes() []byte {
	return binary.LittleEndian.AppendUint32(nil, uint32(p))
}

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
	numberPrefix NumberPrefix
	recordOffset Offset
	cardTypeID   phonedatatool.CardTypeID
}

func (ii IndexItem) Bytes() []byte {
	w := bytes.NewBuffer(nil)
	w.Write(ii.numberPrefix.Bytes())
	w.Write(ii.recordOffset.Bytes())
	w.Write(ii.cardTypeID.Bytes())
	return w.Bytes()
}
func (ii *IndexItem) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 9)
	if _, err := reader.Read(buf); err != nil {
		return err
	}
	ii.numberPrefix = NumberPrefix(binary.LittleEndian.Uint32(buf[:4]))
	ii.recordOffset = Offset(binary.LittleEndian.Uint32(buf[4:8]))
	ii.cardTypeID = phonedatatool.CardTypeID(buf[8])
	return nil
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
			numberPrefix: numberPrefix,
			recordOffset: recordOffset,
			cardTypeID:   cardTypeID,
		}
	}
	return nil
}

func (p *IndexPart) Bytes() []byte {
	var prefixList NumberPrefixList
	for k := range p.prefix2item {
		prefixList = append(prefixList, k)
	}
	sort.Sort(prefixList)

	w := bytes.NewBuffer(nil)
	for _, prefix := range prefixList {
		w.Write(p.prefix2item[prefix].Bytes())
	}
	return w.Bytes()
}

func (p *IndexPart) Parse(reader *bytes.Reader) error {
	for reader.Len() > 0 {
		item := new(IndexItem)
		if err := item.Parse(reader); err != nil {
			return err
		}
		p.prefix2item[item.numberPrefix] = item
	}
	return nil
}
