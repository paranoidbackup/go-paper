package crypt

type EncrypterInput interface {
	DocID() string
	KeyCount() int
	Data() []byte
}
