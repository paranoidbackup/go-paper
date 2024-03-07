package crypt

type EncryptOutput struct {
	projectID   string
	docID       string
	publicKeys  []string
	privateKeys []string
	passphrases []string
	data        string
}

func (b *EncryptOutput) ProjectID() string {
	return b.projectID
}
func (b *EncryptOutput) DocID() string {
	return b.docID
}
func (b *EncryptOutput) PublicKeys() []string {
	return b.publicKeys
}
func (b *EncryptOutput) PrivateKeys() []string {
	return b.privateKeys
}
func (b *EncryptOutput) Passphrases() []string {
	return b.passphrases
}
func (b *EncryptOutput) Data() string {
	return b.data
}
