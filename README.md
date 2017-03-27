手机号码归属地信息库、手机号归属地查询
----------------------------

### 这可能是github上能找到的最新最全的中国境内手机号归属地信息库
基于GO语言实现，使用二分查找法。

 - 归属地信息库文件大小：3,203,029 字节
 - 归属地信息库最后更新：2017年1月
 - 手机号段记录条数：354522

### phone.dat文件格式
 
        | 4 bytes |                     <- phone.dat 版本号（如：1701即17年1月份）
        ------------
        | 4 bytes |                     <-  第一个索引的偏移
        -----------------------
        |  offset - 8            |      <-  记录区
        -----------------------
        |  index                 |      <-  索引区
        -----------------------

1. 头部为8个字节，版本号为4个字节，第一个索引的偏移为4个字节 ；
2. 记录区 中每条记录的格式为"<省份>|<城市>|<邮编>|<长途区号>\0"。 每条记录以'\0'结束；
3. 索引区 中每条记录的格式为"<手机号前七位><记录区的偏移><卡类型>"，每个索引的长度为9个字节； 

### 安装使用

 vi test.go

```
package main

import (
	"fmt"

	"github.com/xluohome/phonedata"
)

func main() {
	pr, err := phonedata.Find("18957509123")
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
}

````
go run test.go

```
PhoneNum: 18957509123
AreaZone: 0575
CardType: 中国电信
City: 绍兴
ZipCode: 312000
Province: 浙江
```

### 性能测试
Thinkpad s3 (Intel(R) Core(TM) i5-4200U CPU @ 1.60GHz)

```
go test -v --bench="."

BenchmarkFindPhone-4      2000000              710 ns/op

```



### 其他语言实现

python: https://github.com/lovedboy/phone
 
php :  https://github.com/shitoudev/phone-location , https://github.com/iwantofun/php_phone

php ext: https://github.com/jonnywang/phone

java: https://github.com/fengjiajie/phone-number-geo

### 感谢
@lovedboy https://github.com/lovedboy

@zhengji  https://github.com/zheng-ji/gophone

