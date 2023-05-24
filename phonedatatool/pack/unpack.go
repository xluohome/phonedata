package pack

import (
	"bytes"
	"github.com/xluohome/phonedata/phonedatatool"
)

type Unpacker struct {
}

func NewUnpacker() phonedatatool.Unpacker {
	return new(Unpacker)
}

func (u *Unpacker) Unpack(phoneDataBuf []byte) (versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf []byte, err error) {
	reader := bytes.NewReader(phoneDataBuf)
	versionPart := new(VersionPart)
	if err := versionPart.Parse(reader); err != nil {
		return nil, nil, nil, err
	}
	versionPlainTextBuf = versionPart.BytesPlainText()

	var indexPartOffset Offset
	if err := indexPartOffset.Parse(reader); err != nil {
		return nil, nil, nil, err
	}

	recordBuf := make([]byte, indexPartOffset-RecordPartBaseOffset)
	if _, err := reader.Read(recordBuf); err != nil {
		return nil, nil, nil, err
	}

	recordPart := NewRecordPart()
	offset2id := make(map[Offset]RecordID)
	if o2i, err := recordPart.Parse(bytes.NewReader(recordBuf), RecordPartBaseOffset); err != nil {
		return nil, nil, nil, err
	} else {
		offset2id = o2i
	}
	recordPlainTextBuf = recordPart.BytesPlainText()

	indexPart := NewIndexPart()
	if err := indexPart.Parse(reader); err != nil {
		return nil, nil, nil, err
	}
	indexPart.BytesPlainText(offset2id)
	indexPlainTextBuf = indexPart.BytesPlainText(offset2id)

	return versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf, nil
}
