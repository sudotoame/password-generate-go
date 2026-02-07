package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (cloud *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (cloud *CloudDb) Write(content []byte) {

}
