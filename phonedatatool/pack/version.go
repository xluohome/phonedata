package pack

import (
	"bytes"
	"fmt"
)

type VersionPart struct {
	version string
}

// ParsePlainText 从文本文件读取
func (p *VersionPart) ParsePlainText(reader *bytes.Reader) error {
	buf := make([]byte, 4)
	if n, err := reader.Read(buf); err != nil {
		return err
	} else if n != 4 {
		return fmt.Errorf("expect read 4 bytes, but read %v bytes", n)
	}
	p.version = string(buf)
	return nil
}

// Bytes 打包成二进制文件里的样子
func (p *VersionPart) Bytes() []byte {
	w := bytes.NewBuffer(nil)
	w.WriteString(p.version)
	for w.Len() < 4 {
		w.WriteByte(0)
	}
	return w.Bytes()
}

func (p *VersionPart) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 4)
	if _, err := reader.Read(buf); err != nil {
		return err
	}
	p.version = string(buf)
	return nil
}

func (p *VersionPart) BytesPlainText() []byte {
	w := bytes.NewBuffer(nil)
	w.Write([]byte(p.version))
	w.WriteByte('\n')
	return w.Bytes()
}
