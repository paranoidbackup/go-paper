package crypt

type EncryptNewProjectOutput struct {
	projectID   string
	docID       string
	publicKeys  []string
	privateKeys []string
	passphrases []string
	data        string
}

func (b *EncryptNewProjectOutput) ProjectID() string {
	return b.projectID
}
func (b *EncryptNewProjectOutput) DocID() string {
	return b.docID
}
func (b *EncryptNewProjectOutput) PublicKeys() []string {
	return b.publicKeys
}
func (b *EncryptNewProjectOutput) PrivateKeys() []string {
	return b.privateKeys
}
func (b *EncryptNewProjectOutput) Passphrases() []string {
	return b.passphrases
}
func (b *EncryptNewProjectOutput) Data() string {
	return b.data
}
