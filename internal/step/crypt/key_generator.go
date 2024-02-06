package crypt

import "github.com/ProtonMail/gopenpgp/v2/crypto"

type KeyGenerator interface {
	GenerateKey(docId string, passphrase string) (*crypto.Key, error)
}

type keyGeneratorImpl struct{}

func newKeyGenerator() *keyGeneratorImpl {
	return &keyGeneratorImpl{}
}

func (g *keyGeneratorImpl) GenerateKey(docId string, passphrase string) (*crypto.Key, error) {
	key, err := crypto.GenerateKey(docId, "", "rsa", 4096)
	if err != nil {
		return nil, err
	}
	defer key.ClearPrivateParams()

	return key.Lock([]byte(passphrase))
}
