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

	_, err := Find("10074872323")
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

func TestFindPhone1952947(t *testing.T) {

	//1952947,广西,玉林,移动
	info, err := Find("1952947")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1952947,广西,玉林,中国移动" {
		t.Fatal("验证失败")
	}

}

func TestFindPhone1669981(t *testing.T) {

	//1669981,新疆,乌鲁木齐,联通
	info, err := Find("1669981")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1669981,新疆,乌鲁木齐,中国联通" {
		t.Fatal("验证失败")
	}

}

func TestFindPhone1921306(t *testing.T) {

	//1921306,河北,保定,中国广电
	info, err := Find("1921306")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1921306,河北,保定,中国广电" {
		t.Fatal("验证失败")
	}

}

func TestFindPhone1936137(t *testing.T) {

	//1936137,广东,广州,中国电信
	info, err := Find("1936137")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1936137,广东,广州,中国电信" {
		t.Fatal("验证失败")
	}

}

func TestFindPhone1903845(t *testing.T) {

	//1903845,贵州,遵义,中国电信
	info, err := Find("1903845")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1903845,贵州,遵义,中国电信" {
		t.Fatal("验证失败")
	}

}
