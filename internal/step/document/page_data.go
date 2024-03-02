package document

type PageData struct {
	ProjectID               string
	DocID                   string
	Passphrases             []string
	PrivateKeysQrCodesPaths [][]string
	DataQrCodesPaths        []string
}
