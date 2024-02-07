package crypt

type EncryptNewProjectInput interface {
	KeyCount() int
	Data() []byte
}

type EncryptInput interface {
	ProjectID() string
	PublicKeys() []string
	Data() []byte
}
