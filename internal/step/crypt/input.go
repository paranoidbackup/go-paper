package crypt

type EncryptNewProjectInput interface {
	KeyCount() int
	Data() []byte
}
