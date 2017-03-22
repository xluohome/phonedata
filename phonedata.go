package phonedata

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
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
	PhoneNum []byte
	Province []byte
	City     []byte
	ZipCode  []byte
	AreaZone []byte
	CardType string
}

var (
	content     []byte
	CardTypemap = map[byte]string{
		CMCC:   "中国移动",
		CUCC:   "中国联通",
		CTCC:   "中国电信",
		CTCC_v: "中国电信虚拟运营商",
		CUCC_v: "中国联通虚拟运营商",
		CMCC_v: "中国移动虚拟运营商",
	}
)

func init() {
	_, fulleFilename, _, _ := runtime.Caller(0)
	var err error
	content, err = ioutil.ReadFile(path.Join(path.Dir(fulleFilename), PHONE_DAT))
	if err != nil {
		panic(err)
	}
}

func Debug() {
	fmt.Println(version())
	fmt.Println(totalRecord())
	fmt.Println(firstRecordOffset())
}

func (pr PhoneRecord) String() string {
	return fmt.Sprintf("PhoneNum: %s\nAreaZone: %s\nCardType: %s\nCity: %s\nZipCode: %s\nProvince: %s\n", pr.PhoneNum, pr.AreaZone, pr.CardType, pr.City, pr.ZipCode, pr.Province)
}

func version() string {
	return string(content[0:INT_LEN])
}

func totalRecord() int32 {
	return (int32(len(content)) - firstRecordOffset()) / PHONE_INDEX_LENGTH
}

func firstRecordOffset() int32 {
	var offset int32
	buffer := bytes.NewBuffer(content[INT_LEN : INT_LEN*2])
	binary.Read(buffer, binary.LittleEndian, &offset)
	return offset
}

func indexRecord(offset int32) (phone_prefix int32, record_offset int32, card_type byte) {
	buffer := bytes.NewBuffer(content[offset : offset+INT_LEN])
	binary.Read(buffer, binary.LittleEndian, &phone_prefix)
	buffer = bytes.NewBuffer(content[offset+INT_LEN : offset+INT_LEN*2])
	binary.Read(buffer, binary.LittleEndian, &record_offset)
	buffer = bytes.NewBuffer(content[offset+INT_LEN*2 : offset+INT_LEN*2+CHAR_LEN])
	binary.Read(buffer, binary.LittleEndian, &card_type)
	return
}

// 二分法查询phone数据
func Find(phone_num string) (pr *PhoneRecord, err error) {
	if len(phone_num) < 7 || len(phone_num) > 11 {
		return nil, errors.New("illegal phone length")
	}

	var left int32
	phone_seven_int, err := strconv.ParseInt(phone_num[0:7], 10, 32)
	if err != nil {
		return nil, errors.New("illegal phone number")
	}
	phone_seven_int32 := int32(phone_seven_int)
	total_len := int32(len(content))
	right := totalRecord()
	firstoffset := firstRecordOffset()
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		current_offset := firstoffset + mid*PHONE_INDEX_LENGTH
		if current_offset >= total_len {
			break
		}
		cur_phone, record_offset, card_type := indexRecord(current_offset)
		switch {
		case cur_phone > phone_seven_int32:
			right = mid - 1
		case cur_phone < phone_seven_int32:
			left = mid + 1
		default:
			cbyte := content[record_offset:]
			end_offset := int32(bytes.Index(cbyte, []byte("\000")))
			data := bytes.Split(cbyte[:end_offset], []byte("|"))
			card_str, ok := CardTypemap[card_type]
			if !ok {
				card_str = "未知电信运营商"
			}
			pr = &PhoneRecord{
				PhoneNum: []byte(phone_num),
				Province: data[0],
				City:     data[1],
				ZipCode:  data[2],
				AreaZone: data[3],
				CardType: card_str,
			}
			return
		}
	}
	return nil, errors.New("phone's data not found")
}
