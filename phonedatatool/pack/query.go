package pack

import (
	"bytes"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool"
	"strconv"
)

type Querier struct {
}

func NewQuerier() *Querier {
	return &Querier{}
}

func (q *Querier) Query(phoneDataBuf []byte, number string) (*phonedatatool.QueryResult, error) {
	if len(number) < 7 {
		return nil, fmt.Errorf("number %v is too short", number)
	}
	var numberPrefix NumberPrefix
	if v, err := strconv.Atoi(number[:7]); err != nil {
		return nil, err
	} else {
		numberPrefix = NumberPrefix(v)
	}

	if result, err := new(Unpacker).unpack(bytes.NewReader(phoneDataBuf)); err != nil {
		return nil, err
	} else {
		var indexItem *IndexItem
		if item, ok := result.indexPart.prefix2item[numberPrefix]; !ok {
			return nil, fmt.Errorf("unknown prefix of number %v", numberPrefix)
		} else {
			indexItem = item
		}
		recordItem := result.recordPart.id2item[result.offset2id[indexItem.recordOffset]]
		return &phonedatatool.QueryResult{
			PhoneNumber:  phonedatatool.PhoneNumber(number),
			AreaCode:     phonedatatool.AreaCode(recordItem.areaCode),
			CardTypeID:   indexItem.cardTypeID,
			CityName:     phonedatatool.CityName(recordItem.city),
			ZipCode:      phonedatatool.ZipCode(recordItem.zipCode),
			ProvinceName: phonedatatool.ProvinceName(recordItem.province),
		}, nil
	}
}
