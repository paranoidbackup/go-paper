package crypt

type EncrypterOutput struct {
	docID      string
	key        []string
	passphrase []string
	data       string
}

func (b *EncrypterOutput) DocID() string {
	return b.docID
}
func (b *EncrypterOutput) Key() []string {
	return b.key
}
func (b *EncrypterOutput) Passphrase() []string {
	return b.passphrase
}
func (b *EncrypterOutput) Data() string {
	return b.data
}
