package crypt

type EncrypterInput interface {
	DocID() string
	Data() []byte
}
