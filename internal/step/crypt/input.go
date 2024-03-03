package crypt

type EncryptNewProjectInput interface {
	Data() []byte
}

type EncryptInput interface {
	ProjectID() string
	PublicKeys() []string
	Data() []byte
}
