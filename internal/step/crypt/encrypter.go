package crypt

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

type Encrypter interface {
	Encrypt(input EncrypterInput) (*EncrypterOutput, error)
}

type EncrypterImpl struct {
	keyGenerator        KeyGenerator
	keySplitter         KeySplitter
	passphraseGenerator PassphraseGenerator
}

func NewEncrypter() *EncrypterImpl {
	return &EncrypterImpl{
		keyGenerator:        newKeyGenerator(),
		keySplitter:         newKeySplitter(),
		passphraseGenerator: newPassphraseGenerator(),
	}
}

func (e *EncrypterImpl) Encrypt(input EncrypterInput) (*EncrypterOutput, error) {
	passphrase, err := e.passphraseGenerator.GeneratePassphrase()
	if err != nil {
		return nil, err
	}
	keyObject, err := e.keyGenerator.GenerateKey(input.DocID(), passphrase)
	if err != nil {
		return nil, err
	}
	data, err := e.doEncrypt(keyObject, input.Data())
	if err != nil {
		return nil, err
	}
	keyArmored, err := keyObject.Armor()
	if err != nil {
		return nil, err
	}
	keySplit, passphraseSplit, err := e.keySplitter.Split(keyArmored, passphrase, input.KeyShards())
	if err != nil {
		return nil, err
	}

	return &EncrypterOutput{
		docID:      input.DocID(),
		key:        keySplit,
		passphrase: passphraseSplit,
		data:       data,
	}, nil
}

func (e *EncrypterImpl) doEncrypt(key *crypto.Key, inputData []byte) (string, error) {
	message := crypto.NewPlainMessage(inputData)
	publicKey, err := key.ToPublic()
	if err != nil {
		return "", err
	}
	publicKeyRing, err := crypto.NewKeyRing(publicKey)
	if err != nil {
		return "", err
	}
	result, err := publicKeyRing.Encrypt(message, nil)
	if err != nil {
		return "", err
	}

	return result.GetArmored()
}
