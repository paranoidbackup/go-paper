package document

type PageData struct {
	ProjectID               string
	DocID                   string
	DocName                 string
	DocDate                 string
	Passphrases             []string
	PrivateKeysQrCodesPaths [][]string
	DataQrCodesPaths        []string
}
