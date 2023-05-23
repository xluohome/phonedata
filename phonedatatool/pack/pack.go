package pack

import "github.com/xluohome/phonedata/phonedatatool"

type Packer struct{}

func NewPacker() phonedatatool.Packer {
	return new(Packer)
}

func (p *Packer) Pack(plainDirectoryPath string, phoneDataFilePath string) error {
	return nil
}
