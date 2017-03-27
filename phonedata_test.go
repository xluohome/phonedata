package phonedata

import (
	"fmt"
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
