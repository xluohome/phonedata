package main

import (
	"flag"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool/pack"
	"github.com/xluohome/phonedata/phonedatatool/query"
	"github.com/xluohome/phonedata/phonedatatool/unpack"
)

// 这里编译出来的可执行程序具备打包、查询、解包三个功能。
// ./phonedatatool -unpack -i phone.dat -o tmp
// ./phonedatatool -pack -i tmp -o phone.dat
// ./phonedatatool -query -i phone.dat -number 13000001234

const (
	Name     = "phonedatatool"
	Version  = "0.1.0"
	Author   = "seedjyh@gmail.com"
	FullName = Name + "_" + Version + "(" + Author + ")"
)

func main() {
	showVersionFlag := flag.Bool("v", false, "Just show version string")
	showHelpFlag := flag.Bool("help", false, "Just print help.")
	unpackFlag := flag.Bool("unpack", false, "Unpack phone data to plain text")
	packFlag := flag.Bool("pack", false, "Pack plain text to phone data")
	queryFlag := flag.Bool("query", false, "Query number from phone data")
	source := flag.String("i", "", "Source of operation")
	destination := flag.String("o", "", "Destination of operation")
	number := flag.String("number", "", "Number to query")
	flag.Parse()
	if *showVersionFlag {
		fmt.Println("Version:", FullName)
		return
	}
	if *showHelpFlag {
		showHelp()
		return
	}
	if *unpackFlag {
		if source == nil {
			fmt.Println("ERROR! No source")
			return
		}
		if destination == nil {
			fmt.Println("ERROR! No destination")
			return
		}
		if err := unpack.NewUnpacker().Unpack(*source, *destination); err != nil {
			fmt.Println("ERROR! Unpack failed.", err)
			return
		} else {
			fmt.Println("Unpack completed.")
			return
		}
	}
	if *packFlag {
		if source == nil {
			fmt.Println("ERROR! No source")
			return
		}
		if destination == nil {
			fmt.Println("ERROR! No destination")
			return
		}
		if err := pack.NewPacker().Pack(*source, *destination); err != nil {
			fmt.Println("ERROR! Pack failed.", err)
			return
		} else {
			fmt.Println("Pack completed.")
			return
		}
	}
	if *queryFlag {
		if source == nil {
			fmt.Println("ERROR! No source")
			return
		}
		if number == nil {
			fmt.Println("ERROR! No number to query")
			return
		}
		if info, err := query.NewQuerier(*source).QueryNumber(*number); err != nil {
			fmt.Println("ERROR! Query failed.", err)
			return
		} else {
			fmt.Println("PhoneNum: ", info.PhoneNumber)
			fmt.Println("AreaZone: ", info.AreaCode)
			fmt.Println("CardType: ", info.CardTypeID.ToName().String())
			fmt.Println("City: ", info.CityName)
			fmt.Println("ZipCode: ", info.ZipCode)
			fmt.Println("Province: ", info.ProvinceName)
			fmt.Println("Query completed.")
			return
		}
	}
	fmt.Println("Did nothing.")
	showHelp()
	return
}

func showHelp() {
	fmt.Println("<< HELP >>")
	fmt.Println("./phonedatatool -unpack -i phone.dat -o tmp")
	fmt.Println("./phonedatatool -pack -i tmp -o phone.dat")
	fmt.Println("./phonedatatool -query -i phone.dat -n 13000001234")
}
