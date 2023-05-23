package unpack

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool"
	"sort"
	"strconv"
	"strings"
)

type PhoneNumberPrefix string

func (pnp PhoneNumberPrefix) String() string {
	return string(pnp)
}

type PhoneNumberPrefixList []PhoneNumberPrefix

func (pl PhoneNumberPrefixList) Len() int {
	return len(pl)
}
func (pl PhoneNumberPrefixList) Less(i, j int) bool {
	return pl[i] < pl[j]
}
func (pl PhoneNumberPrefixList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

type IndexItem struct {
	NumberPrefix PhoneNumberPrefix
	CardTypeID   phonedatatool.CardTypeID
	RecordOffset RecordOffset
	RecordID     RecordID
}

func (ii *IndexItem) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 9)
	if _, err := reader.Read(buf); err != nil {
		return err
	}
	ii.NumberPrefix = PhoneNumberPrefix(strconv.Itoa(int(binary.LittleEndian.Uint32(buf[:4]))))
	ii.RecordOffset = RecordOffset(binary.LittleEndian.Uint32(buf[4:8]))
	ii.CardTypeID = phonedatatool.CardTypeID(buf[8])
	return nil
}

type IndexPart struct {
	prefix2item map[PhoneNumberPrefix]*IndexItem
}

func NewIndexPart() *IndexPart {
	return &IndexPart{
		prefix2item: make(map[PhoneNumberPrefix]*IndexItem),
	}
}

func (ip *IndexPart) Bytes() []byte {
	w := bytes.NewBuffer(nil)
	var prefixList PhoneNumberPrefixList
	for k, _ := range ip.prefix2item {
		prefixList = append(prefixList, k)
	}
	sort.Sort(prefixList)
	for _, prefix := range prefixList {
		item := ip.prefix2item[prefix]
		w.WriteString(strings.Join([]string{
			prefix.String(),
			item.RecordID.String(),
			item.CardTypeID.String(),
		}, "|"))
		w.WriteByte('\n')
	}
	return w.Bytes()
}

func (ip *IndexPart) Parse(reader *bytes.Reader) error {
	for reader.Len() > 0 {
		item := new(IndexItem)
		if err := item.Parse(reader); err != nil {
			return err
		}
		ip.prefix2item[item.NumberPrefix] = item
	}
	return nil
}

func (ip *IndexPart) MatchRecordOffsetToRecordID(offset2id map[RecordOffset]RecordID) error {
	for _, v := range ip.prefix2item {
		if id, ok := offset2id[v.RecordOffset]; ok {
			v.RecordID = id
		} else {
			return fmt.Errorf("failed to find record id for record offset %v", v.RecordOffset)
		}
	}
	return nil
}
