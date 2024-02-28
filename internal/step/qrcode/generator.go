package qrcode

import qrcode "github.com/skip2/go-qrcode"

type Generator interface {
	Generate(input Input) (*Output, error)
}

type GeneratorImpl struct {
}

func NewGenerator() (*GeneratorImpl, error) {
	return &GeneratorImpl{}, nil
}

func (g *GeneratorImpl) Generate(input Input) (*Output, error) {

	// example
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Highest, 256)

	return &Output{
		projectId:   input.ProjectID(),
		docId:       input.DocID(),
		publicKeys:  input.PublicKeys(),
		passphrases: input.Passphrases(),
	}, nil // TODO
}
