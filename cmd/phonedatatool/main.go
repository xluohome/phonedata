package main

import (
	"flag"
	"fmt"
	"github.com/xluohome/phonedata/phonedatatool/pack"
	"github.com/xluohome/phonedata/phonedatatool/query"
	"github.com/xluohome/phonedata/phonedatatool/util"
	"os"
	"path"
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

const (
	VersionFileName = "version.txt"
	RecordFileName  = "record.txt"
	IndexFileName   = "index.txt"
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
		if err := Unpack(*source, *destination); err != nil {
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
		if err := Pack(*source, *destination); err != nil {
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

func Pack(plainDirectoryPath string, phoneDataFilePath string) error {
	if err := util.AssureFileNotExist(phoneDataFilePath); err != nil {
		return err
	}
	var versionPlainTextBuf []byte
	if buf, err := os.ReadFile(path.Join(plainDirectoryPath, VersionFileName)); err != nil {
		return err
	} else {
		versionPlainTextBuf = buf
	}

	var recordPlainTextBuf []byte
	if buf, err := os.ReadFile(path.Join(plainDirectoryPath, RecordFileName)); err != nil {
		return err
	} else {
		recordPlainTextBuf = buf
	}

	var indexPlainTextBuf []byte
	if buf, err := os.ReadFile(path.Join(plainDirectoryPath, IndexFileName)); err != nil {
		return err
	} else {
		indexPlainTextBuf = buf
	}

	if buf, err := pack.NewPacker().Pack(versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf); err != nil {
		return err
	} else {
		return os.WriteFile(phoneDataFilePath, buf, 0)
	}
}

func Unpack(phoneDataFilePath string, plainDirectoryPath string) error {
	if err := os.MkdirAll(plainDirectoryPath, 0); err != nil {
		return fmt.Errorf("target directory %v not exist and can't be created: %v", plainDirectoryPath, err)
	}

	versionFilePath := path.Join(plainDirectoryPath, VersionFileName)
	recordFilePath := path.Join(plainDirectoryPath, RecordFileName)
	indexFilePath := path.Join(plainDirectoryPath, IndexFileName)

	if err := util.AssureAllFileNotExist(versionFilePath, recordFilePath, indexFilePath); err != nil {
		return err
	}

	var rawBuf []byte
	if b, err := os.ReadFile(phoneDataFilePath); err != nil {
		return err
	} else {
		rawBuf = b
	}

	if versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf, err := pack.NewUnpacker().Unpack(rawBuf); err != nil {
		return err
	} else {
		if err := os.WriteFile(versionFilePath, versionPlainTextBuf, 0); err != nil {
			return err
		}
		if err := os.WriteFile(recordFilePath, recordPlainTextBuf, 0); err != nil {
			return err
		}
		if err := os.WriteFile(indexFilePath, indexPlainTextBuf, 0); err != nil {
			return err
		}
		return nil
	}
}
