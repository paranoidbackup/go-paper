package crypt

type EncryptNewProjectInput interface {
	DocName() string
	Data() []byte
}

type EncryptInput interface {
	DocName() string
	ProjectID() string
	PublicKeys() []string
	Data() []byte
}
