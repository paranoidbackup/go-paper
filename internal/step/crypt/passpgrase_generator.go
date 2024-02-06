package crypt

type PassphraseGenerator interface {
	GeneratePassphrase() (string, error)
}

type passphraseGeneratorImpl struct {
}

func newPassphraseGenerator() *passphraseGeneratorImpl {
	return &passphraseGeneratorImpl{}
}

func (g *passphraseGeneratorImpl) GeneratePassphrase() (string, error) {
	return "", nil // TODO
}
