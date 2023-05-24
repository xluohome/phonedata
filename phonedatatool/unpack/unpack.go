package unpack

import (
	"bytes"
	"fmt"
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
		return nil, nil, nil, fmt.Errorf("failed to read version part: %v", err)
	}
	offsetPart := new(IndexPartOffsetPart)
	if err := offsetPart.Parse(reader); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read index-part-offset part: %v", err)
	}
	recordPart := NewRecordPart()
	if err := recordPart.Parse(reader, offsetPart.IndexPartOffset); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read record part: %v", err)
	}
	indexPart := NewIndexPart()
	if err := indexPart.Parse(reader); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read index part: %v", err)
	}
	if err := indexPart.MatchRecordOffsetToRecordID(recordPart.offset2id); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to match offset to record id: %v", err)
	}

	versionPlainTextBuf = versionPart.Bytes()
	recordPlainTextBuf = recordPart.Bytes()
	indexPlainTextBuf = indexPart.Bytes()
	return versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf, nil
}
