package qrcode

type Output struct {
	projectId             string
	docId                 string
	docName               string
	publicKeys            []string
	passphrases           []string
	privateKeysSplit      [][]string
	dataSplit             []string
	privateKeysQrCodesPng [][][]byte
	dataQrCodesPng        [][]byte
}

func (o *Output) ProjectID() string {
	return o.projectId
}
func (o *Output) DocID() string {
	return o.docId
}
func (o *Output) DocName() string {
	return o.docName
}
func (o *Output) PublicKeys() []string {
	return o.publicKeys
}
func (o *Output) Passphrases() []string {
	return o.passphrases
}
func (o *Output) PrivateKeysSplit() [][]string {
	return o.privateKeysSplit
}
func (o *Output) DataSplit() []string {
	return o.dataSplit
}
func (o *Output) PrivateKeysQrCodesPNG() [][][]byte {
	return o.privateKeysQrCodesPng
}
func (o *Output) DataQrCodesPNG() [][]byte {
	return o.dataQrCodesPng
}
