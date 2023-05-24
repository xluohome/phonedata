package phonedatatool

type Unpacker interface {
	// Unpack 将压缩文件 phoneDataFilePath 解包后保存到明文目录 plainDirectoryPath 里。
	Unpack(phoneDataFilePath string, plainDirectoryPath string) error
}

type Packer interface {
	// Pack 将明文目录 plainDirectoryPath 的数据打包成压缩文件 phoneDataFilePath。
	Pack(plainDirectoryPath string, phoneDataFilePath string) error
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
