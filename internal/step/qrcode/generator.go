package qrcode

type Generator interface {
	Generate(input Input) (*Output, error)
}

type GeneratorImpl struct {
}

func NewGenerator() (*GeneratorImpl, error) {
	return &GeneratorImpl{}, nil
}

func (g *GeneratorImpl) Generate(input Input) (*Output, error) {
	return nil, nil // TODO
}
