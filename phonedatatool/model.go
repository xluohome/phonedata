package phonedatatool

import (
	"github.com/xluohome/phonedata"
	"strconv"
)

type PhoneNumber string // 手机号码
type AreaCode string    // 区号
func (ac AreaCode) String() string {
	return string(ac)
}

type CardTypeID byte // 卡类型 ID，单字节

func (ctid CardTypeID) Bytes() []byte {
	return []byte{byte(ctid)}
}

func (ctid CardTypeID) String() string {
	return strconv.Itoa(int(ctid))
}
func (ctid CardTypeID) ToName() CardTypeName {
	if v, ok := phonedata.CardTypemap[byte(ctid)]; ok {
		return CardTypeName(v)
	} else {
		return "---"
	}
}

type CardTypeName string // 卡类型中文名
func (n CardTypeName) String() string {
	return string(n)
}

type CityName string // 城市名
func (cn CityName) String() string {
	return string(cn)
}

type ProvinceName string // 省名
func (pn ProvinceName) String() string {
	return string(pn)
}

type ZipCode string // 邮政编码
func (zc ZipCode) String() string {
	return string(zc)
}
