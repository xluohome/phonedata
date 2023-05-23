package unpack

import (
	"bytes"
	"fmt"
)

type Version string

func (v Version) String() string {
	return string(v)
}

type VersionPart struct {
	Version Version
}

func (vp *VersionPart) Bytes() []byte {
	w := bytes.NewBuffer(nil)
	w.Write([]byte(vp.Version))
	w.WriteByte('\n')
	return w.Bytes()
}

func (vp *VersionPart) Parse(reader *bytes.Reader) error {
	buf := make([]byte, 4)
	if _, err := reader.Read(buf); err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	vp.Version = Version(buf)
	return nil
}
