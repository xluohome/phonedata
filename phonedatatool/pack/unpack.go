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

type unpackResult struct {
	versionPart *VersionPart
	recordPart  *RecordPart
	offset2id   map[Offset]RecordID
	indexPart   *IndexPart
}

func (u *Unpacker) Unpack(phoneDataBuf []byte) (versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf []byte, err error) {
	if result, err := u.unpack(bytes.NewReader(phoneDataBuf)); err != nil {
		return nil, nil, nil, err
	} else {
		return result.versionPart.BytesPlainText(), result.recordPart.BytesPlainText(), result.indexPart.BytesPlainText(result.offset2id), nil
	}
}

func (u *Unpacker) unpack(reader *bytes.Reader) (*unpackResult, error) {
	versionPart := new(VersionPart)
	if err := versionPart.Parse(reader); err != nil {
		return nil, err
	}

	var indexPartOffset Offset
	if err := indexPartOffset.Parse(reader); err != nil {
		return nil, err
	}

	recordBuf := make([]byte, indexPartOffset-RecordPartBaseOffset)
	if _, err := reader.Read(recordBuf); err != nil {
		return nil, err
	}

	recordPart := NewRecordPart()
	offset2id := make(map[Offset]RecordID)
	if o2i, err := recordPart.Parse(bytes.NewReader(recordBuf), RecordPartBaseOffset); err != nil {
		return nil, err
	} else {
		offset2id = o2i
	}

	indexPart := NewIndexPart()
	if err := indexPart.Parse(reader); err != nil {
		return nil, err
	}
	indexPart.BytesPlainText(offset2id)

	return &unpackResult{
		versionPart: versionPart,
		recordPart:  recordPart,
		offset2id:   offset2id,
		indexPart:   indexPart,
	}, nil
}
