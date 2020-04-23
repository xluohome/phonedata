手机号码归属地信息库、手机号归属地查询
----------------------------

### 这可能是github上能找到的最新最全的中国境内手机号归属地信息库
基于GO语言实现，使用二分查找法。

 - 归属地信息库文件大小：4,040,893 字节
 - 归属地信息库最后更新：2020年04月
 - 手机号段记录条数：447893

### phone.dat文件格式

        | 4 bytes |                     <- phone.dat 版本号（如：1701即17年1月份）
        ------------
        | 4 bytes |                     <-  第一个索引的偏移
        -----------------------
        |  offset - 8            |      <-  记录区
        -----------------------
        |  index                 |      <-  索引区
        -----------------------

1. 头部为8个字节，版本号为4个字节，第一个索引的偏移为4个字节；
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

### 快速使用

cmd 目录下phonedata是一个命令行查询手机号归属地信息的终端程序。
```

Linux:
#PHONE_DATA_DIR=../ ./phonedata  18957509123

Windows:
>set PHONE_DATA_DIR=../
>phonedata.exe  18957509123
```
stdout:
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

go version go1.13.4 windows/amd64

```
go test -v --bench="."

BenchmarkFindPhone-4     1964847               607 ns/op

```

### 我仅想要phone.dat的csv文本文件?

好。下载地址
https://git.oschina.net/oss/phonedata/attach_files


### 其他语言实现

python: https://github.com/lovedboy/phone

php:  https://github.com/shitoudev/phone-location , https://github.com/iwantofun/php_phone

php ext: https://github.com/jonnywang/phone

java: https://github.com/fengjiajie/phone-number-geo , https://github.com/EeeMt/phone-number-geo

Node: https://github.com/conzi/phone

C++: https://github.com/yanxijian/phonedata

C#: https://github.com/sndnvaps/Phonedata

Rust: https://github.com/vincascm/phonedata

### 安全保证

手机号归属地信息是通过网上公开数据进行收集整理。

对手机号归属地信息数据的绝对正确，我不做任何保证。因此在生产环境使用前请您自行校对测试。


### 客户案例

- [360](https://www.360.cn/)
- [MAGAPP](http://www.magapp.cc/)
- ...

### 感谢
@lovedboy https://github.com/lovedboy

@zhengji  https://github.com/zheng-ji/gophone

### 联系作者

加作者微信

![wx.jpg](https://ucc.alicdn.com/pic/developer-ecology/f41fd688affb41fc8853c4f99abd3d45.jpg)
