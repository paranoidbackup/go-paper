package document

type PageData struct {
	ProjectID               string
	DocID                   string
	DocName                 string
	Passphrases             []string
	PrivateKeysQrCodesPaths [][]string
	DataQrCodesPaths        []string
}
