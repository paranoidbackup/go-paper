package qrcode

type Input interface {
	ProjectID() string
	DocID() string
	PublicKeys() []string
	PrivateKeys() []string
	Passphrases() []string
	Data() string
}
