package phonedata

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

const (
	CMCC               byte = iota + 0x01 //中国移动
	CUCC                                  //中国联通
	CTCC                                  //中国电信
	CTCC_v                                //电信虚拟运营商
	CUCC_v                                //联通虚拟运营商
	CMCC_v                                //移动虚拟运营商
	INT_LEN            = 4
	CHAR_LEN           = 1
	HEAD_LENGTH        = 8
	PHONE_INDEX_LENGTH = 9
	PHONE_DAT          = "phone.dat"
)

type PhoneRecord struct {
	PhoneNum string
	Province string
	City     string
	ZipCode  string
	AreaZone string
	CardType string
}

type Phonedata struct {
	Province []byte
	City     []byte
	ZipCode  []byte
	AreaZone []byte
	CardType byte
}

var (
	CardTypemap = map[byte]string{
		CMCC:   "中国移动",
		CUCC:   "中国联通",
		CTCC:   "中国电信",
		CTCC_v: "中国电信虚拟运营商",
		CUCC_v: "中国联通虚拟运营商",
		CMCC_v: "中国移动虚拟运营商",
	}
	total_len, firstoffset int32
	version                string
	phonemap               map[int32][][]byte
)

func init() {
	dir := os.Getenv("PHONE_DATA_DIR")
	if dir == "" {
		_, fulleFilename, _, _ := runtime.Caller(0)
		dir = path.Dir(fulleFilename)
	}
	content, err := ioutil.ReadFile(path.Join(dir, PHONE_DAT))
	if err != nil {
		panic(err)
	}
	version = string(content[0:INT_LEN])
	total_len = int32(len(content))
	firstoffset = get4(content[INT_LEN : INT_LEN*2])

	phonemap = make(map[int32][][]byte, total_len)
	var i int32
	for i = 0; i < total_len; i++ {
		offset := firstoffset + i*PHONE_INDEX_LENGTH
		if offset >= total_len {
			break
		}
		cur_phone := get4(content[offset : offset+INT_LEN])
		record_offset := get4(content[offset+INT_LEN : offset+INT_LEN*2])
		card_type := content[offset+INT_LEN*2 : offset+INT_LEN*2+CHAR_LEN]

		cbyte := content[record_offset:]
		end_offset := int32(bytes.Index(cbyte, []byte("\000")))
		data := bytes.Split(cbyte[:end_offset], []byte("|"))
		phonemap[cur_phone] = [][]byte{
			data[0],
			data[1],
			data[2],
			data[3],
			card_type,
		}
	}
}

func Debug() {
	fmt.Println(version)
	fmt.Println((total_len - firstoffset) / PHONE_INDEX_LENGTH)
	fmt.Println(firstoffset)
}

func (pr PhoneRecord) String() string {
	return fmt.Sprintf("PhoneNum: %s\nAreaZone: %s\nCardType: %s\nCity: %s\nZipCode: %s\nProvince: %s\n", pr.PhoneNum, pr.AreaZone, pr.CardType, pr.City, pr.ZipCode, pr.Province)
}

func get4(b []byte) int32 {
	if len(b) < 4 {
		return 0
	}
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

func getN(s string) (uint32, error) {
	var n, cutoff, maxVal uint32
	i := 0
	base := 10
	cutoff = (1<<32-1)/10 + 1
	maxVal = 1<<uint(32) - 1
	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			return 0, errors.New("invalid syntax")
		}
		if v >= byte(base) {
			return 0, errors.New("invalid syntax")
		}

		if n >= cutoff {
			// n*base overflows
			n = (1<<32 - 1)
			return n, errors.New("value out of range")
		}
		n *= uint32(base)

		n1 := n + uint32(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = (1<<32 - 1)
			return n, errors.New("value out of range")
		}
		n = n1
	}
	return n, nil
}

//map查询
func Find(phone_num string) (pr *PhoneRecord, err error) {
	if len(phone_num) < 7 || len(phone_num) > 11 {
		return nil, errors.New("illegal phone length")
	}

	phone_seven_int, err := getN(phone_num[0:7])
	if err != nil {
		return nil, errors.New("illegal phone number")
	}
	phone_seven_int32 := int32(phone_seven_int)

	if data, ok := phonemap[phone_seven_int32]; ok {

		card_str, ok1 := CardTypemap[data[4][0]]
		if !ok1 {
			card_str = "未知电信运营商"
		}
		pr = &PhoneRecord{
			PhoneNum: phone_num,
			Province: string(data[0]),
			City:     string(data[1]),
			ZipCode:  string(data[2]),
			AreaZone: string(data[3]),
			CardType: card_str,
		}
		return
	}
	return nil, errors.New("phone's data not found")
}
