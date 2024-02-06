package crypt

type EncrypterInput interface {
	DocID() string
	KeyShards() int
	Data() []byte
}
