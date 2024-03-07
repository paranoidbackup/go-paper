package document

type ExporterInput interface {
	ProjectID() string
	DocID() string
	PublicKeys() []string
	Passphrases() []string
	PrivateKeysSplit() [][]string
	DataSplit() []string
	PrivateKeysQrCodesPNG() [][][]byte
	DataQrCodesPNG() [][]byte
}
