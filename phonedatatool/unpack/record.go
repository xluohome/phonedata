package unpack

import (
	"bytes"
	"github.com/xluohome/phonedata/phonedatatool"
	"sort"
	"strconv"
	"strings"
)

type RecordItem struct {
	Province phonedatatool.ProvinceName
	City     phonedatatool.CityName
	ZipCode  phonedatatool.ZipCode
	AreaCode phonedatatool.AreaCode
}
type RecordOffset int64

func (ro RecordOffset) String() string {
	return strconv.FormatInt(int64(ro), 10)
}

type RecordID int64

func (rid RecordID) String() string {
	return strconv.FormatInt(int64(rid), 10)
}

type RecordIDList []RecordID

func (l RecordIDList) Len() int {
	return len(l)
}
func (l RecordIDList) Less(i, j int) bool {
	return l[i] < l[j]
}
func (l RecordIDList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type RecordPart struct {
	offset2item map[RecordOffset]*RecordItem
	offset2id   map[RecordOffset]RecordID
	id2offset   map[RecordID]RecordOffset
}

func (rp *RecordPart) Bytes() []byte {
	w := bytes.NewBuffer(nil)
	var idList RecordIDList
	for k, _ := range rp.id2offset {
		idList = append(idList, k)
	}
	sort.Sort(idList)
	for _, id := range idList {
		item := rp.offset2item[rp.id2offset[id]]
		w.WriteString(strings.Join([]string{
			id.String(),
			item.Province.String(),
			item.City.String(),
			item.ZipCode.String(),
			item.AreaCode.String(),
		}, "|"))
		w.WriteByte('\n')
	}
	return w.Bytes()
}

func (rp *RecordPart) Parse(reader *bytes.Reader) error {
	return nil
}
