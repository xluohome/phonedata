package unpack

import (
	"bytes"
	"github.com/xluohome/phonedata/phonedatatool"
	"sort"
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
	CardTypeID   phonedatatool.CardTypeID
	RecordOffset RecordOffset
	RecordID     RecordID
}

type IndexPart struct {
	prefix2item map[PhoneNumberPrefix]*IndexItem
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
			item.RecordOffset.String(),
			item.CardTypeID.String(),
		}, "|"))
		w.WriteByte('\n')
	}
	return w.Bytes()
}

func (ip *IndexPart) Parse(reader *bytes.Reader) error {
	return nil
}
