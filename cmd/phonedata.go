package main

import (
	"fmt"
	"os"

	"github.com/xluohome/phonedata"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Print("请输入手机号")
		return
	}
	pr, err := phonedata.Find(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Print(pr)
}
