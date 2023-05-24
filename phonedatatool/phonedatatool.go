package phonedatatool

type Unpacker interface {
	// Unpack 将压缩文件 phoneDataFilePath 解包后保存到明文目录 plainDirectoryPath 里。
	Unpack(phoneDataFilePath string, plainDirectoryPath string) error
}

type Packer interface {
	// Pack 将版本文件、记录文件、索引文件的内容打包成二进制文件。
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
