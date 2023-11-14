package pack

import (
	"bytes"
	"github.com/xluohome/phonedata/phonedatatool"
)

type Packer struct{}

const RecordPartBaseOffset = Offset(8) // record part 首字节偏移量

func NewPacker() phonedatatool.Packer {
	return new(Packer)
}

func (p *Packer) Pack(versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf []byte) ([]byte, error) {
	versionPart := new(VersionPart)
	if err := versionPart.ParsePlainText(bytes.NewReader(versionPlainTextBuf)); err != nil {
		return nil, err
	}
	versionPartBuf := versionPart.Bytes()

	recordPart := NewRecordPart()
	if err := recordPart.ParsePlainText(bytes.NewReader(recordPlainTextBuf)); err != nil {
		return nil, err
	}
	recordPartBuf, recordID2Offset := recordPart.Bytes(RecordPartBaseOffset)

	indexPartOffsetPart := RecordPartBaseOffset + Offset(len(recordPartBuf))

	indexPart := NewIndexPart()
	if err := indexPart.ParsePlainText(bytes.NewReader(indexPlainTextBuf), recordID2Offset); err != nil {
		return nil, err
	}
	indexPartBuf := indexPart.Bytes()

	w := bytes.NewBuffer(nil)
	if _, err := w.Write(versionPartBuf); err != nil {
		return nil, err
	}
	if _, err := w.Write(indexPartOffsetPart.Bytes()); err != nil {
		return nil, err
	}
	if _, err := w.Write(recordPartBuf); err != nil {
		return nil, err
	}
	if _, err := w.Write(indexPartBuf); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
