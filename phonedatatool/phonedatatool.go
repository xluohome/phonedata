package phonedatatool

type Unpacker interface {
	// Unpack 将二进制文件的内容解包成版本文件、记录文件、索引文件的内容。
	Unpack(phoneDataBuf []byte) (versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf []byte, err error)
}

type Packer interface {
	// Pack 将版本文件、记录文件、索引文件的内容打包成二进制文件的内容。
	Pack(versionPlainTextBuf, recordPlainTextBuf, indexPlainTextBuf []byte) ([]byte, error)
}

type QueryResult struct {
	PhoneNumber  PhoneNumber
	AreaCode     AreaCode
	CardTypeID   CardTypeID
	CityName     CityName
	ZipCode      ZipCode
	ProvinceName ProvinceName
}

type Querier interface {
	QueryNumber(number string) (*QueryResult, error)
}

const (
	VersionFileName = "version.txt"
	RecordFileName  = "record.txt"
	IndexFileName   = "index.txt"
)
