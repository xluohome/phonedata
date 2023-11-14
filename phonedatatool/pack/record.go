package pack

import (
	"bytes"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool/util"
	"sort"
	"strconv"
	"strings"
)

type RecordID int64

func (rid RecordID) String() string {
	return strconv.FormatInt(int64(rid), 10)
}

type RecordIDList []RecordID

func (idl RecordIDList) Len() int {
	return len(idl)
}
func (idl RecordIDList) Less(i, j int) bool {
	return idl[i] < idl[j]
}
func (idl RecordIDList) Swap(i, j int) {
	idl[i], idl[j] = idl[j], idl[i]
}

type RecordItem struct {
	province string
	city     string
	zipCode  string
	areaCode string
}

func (ri *RecordItem) Bytes() []byte {
	w := bytes.NewBuffer(nil)
	w.WriteString(strings.Join([]string{ri.province, ri.city, ri.zipCode, ri.areaCode}, "|"))
	w.WriteByte(0)
	return w.Bytes()
}

// Parse 从压缩文件读取一条 RecordItem。注意 reader 必须以 '\0' 结尾。
func (ri *RecordItem) Parse(reader *bytes.Reader) error {
	if buf, err := util.ReadUntil(reader, 0); err != nil {
		return fmt.Errorf("no term char for record item: %v", err)
	} else {
		words := strings.Split(string(buf), "|")
		if len(words) != 4 {
			return fmt.Errorf("invalid item bytes, %v", string(buf))
		}
		ri.province = words[0]
		ri.city = words[1]
		ri.zipCode = words[2]
		ri.areaCode = words[3]
		return nil
	}
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

func (p *RecordPart) Bytes(baseOffset Offset) ([]byte, map[RecordID]Offset) {
	w := bytes.NewBuffer(nil)
	id2offset := make(map[RecordID]Offset)

	var idList RecordIDList
	for k := range p.id2item {
		idList = append(idList, k)
	}
	sort.Sort(idList)

	for _, id := range idList {
		id2offset[id] = baseOffset + Offset(w.Len())
		w.Write(p.id2item[id].Bytes())
	}
	return w.Bytes(), id2offset
}

func (p *RecordPart) Parse(reader *bytes.Reader, baseOffset Offset) (map[Offset]RecordID, error) {
	offset2id := make(map[Offset]RecordID)
	offset := baseOffset
	for id := RecordID(1); reader.Len() > 0; id++ {
		var itemBuf []byte
		if buf, err := util.ReadUntil(reader, 0); err != nil {
			return nil, err
		} else {
			itemBuf = buf
			itemBuf = append(itemBuf, 0)
		}
		item := new(RecordItem)
		if err := item.Parse(bytes.NewReader(itemBuf)); err != nil {
			return nil, err
		}
		offset2id[offset] = id
		p.id2item[id] = item
		offset += Offset(len(itemBuf))
	}
	return offset2id, nil
}

func (p *RecordPart) BytesPlainText() []byte {
	var idList RecordIDList
	for k := range p.id2item {
		idList = append(idList, k)
	}
	sort.Sort(idList)

	w := bytes.NewBuffer(nil)
	for _, id := range idList {
		item := p.id2item[id]
		w.WriteString(strings.Join([]string{
			id.String(),
			item.province,
			item.city,
			item.zipCode,
			item.areaCode,
		}, "|"))
		w.WriteByte('\n')
	}
	return w.Bytes()
}
