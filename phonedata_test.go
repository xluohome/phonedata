package phonedata

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func since(t time.Time) int64 {
	return time.Since(t).Nanoseconds()
}

func init() {

	Debug()
}

func BenchmarkFindPhone(b *testing.B) {

	b.RunParallel(func(p *testing.PB) {

		var i = 0
		for p.Next() {
			i++
			_, err := Find(fmt.Sprintf("%s%d%s", "1897", i&10000, "45"))
			if err != nil {
				b.Fatal(err)
			}
		}

	})

}

func TestFindPhone1(t *testing.T) {

	_, err := Find("13580198235123123213213")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}
func TestFindPhone2(t *testing.T) {

	_, err := Find("1300")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}
func TestFindPhone3(t *testing.T) {

	pr, err := Find("1703576")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pr)
}

func TestFindPhone4(t *testing.T) {

	_, err := Find("19174872323")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestFindPhone5(t *testing.T) {

	_, err := Find("afsd32323")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestDump(t *testing.T) {
	records, err := Dump()
	if err != nil {
		t.Fatal("错误的结果：", err)
	}

	if len(records) != 387695 {
		t.Fatal("错误的数量：" + strconv.Itoa(len(records)))
	}

	if records[0].String() != `PhoneNum: 1300000xxxx
AreaZone: 0531
CardType: 中国联通
City: 济南
ZipCode: 250000
Province: 山东
` {
		t.Fatal("错误的结果：", records[0].String())
	}
}
