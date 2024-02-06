package crypt

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/armor"
	"github.com/ProtonMail/gopenpgp/v2/constants"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

type Encrypter interface {
	Encrypt(input EncrypterInput) (*EncrypterOutput, error)
}

type EncrypterImpl struct {
	keyGen  KeyGenerator
	passGen PassphraseGenerator
}

type encryptRoundResult struct {
	key        string
	passphrase string
	data       []byte
}

func NewEncrypter() *EncrypterImpl {
	return &EncrypterImpl{
		keyGen:  newKeyGenerator(),
		passGen: newPassphraseGenerator(),
	}
}

func (e *EncrypterImpl) Encrypt(input EncrypterInput) (*EncrypterOutput, error) {
	if input.KeyCount() <= 0 {
		return nil, fmt.Errorf("key count should be >=1, got %v", input.KeyCount())
	}

	keys := make([]string, 0)
	passphrases := make([]string, 0)
	data := input.Data()

	for round := 0; round < input.KeyCount(); round++ {
		roundResult, err := e.encryptRound(round, input.DocID(), data)
		if err != nil {
			return nil, err
		}
		keys = append(keys, roundResult.key)
		passphrases = append(passphrases, roundResult.passphrase)
		data = roundResult.data
	}

	dataArmored, err := armor.ArmorWithType(data, constants.PGPMessageHeader)
	if err != nil {
		return nil, err
	}

	return &EncrypterOutput{
		docID:      input.DocID(),
		key:        keys,
		passphrase: passphrases,
		data:       dataArmored,
	}, nil
}

func (e *EncrypterImpl) encryptRound(roundId int, docId string, inputData []byte) (*encryptRoundResult, error) {
	passphrase, err := e.passGen.GeneratePassphrase()
	if err != nil {
		return nil, err
	}
	keyObject, err := e.keyGen.GenerateKey(fmt.Sprintf("%v_%v", docId, roundId), passphrase)
	if err != nil {
		return nil, err
	}
	data, err := e.doEncrypt(keyObject, inputData)
	if err != nil {
		return nil, err
	}
	key, err := keyObject.Armor()
	if err != nil {
		return nil, err
	}
	return &encryptRoundResult{
		key:        key,
		passphrase: passphrase,
		data:       data,
	}, nil
}

func (e *EncrypterImpl) doEncrypt(key *crypto.Key, inputData []byte) ([]byte, error) {
	message := crypto.NewPlainMessage(inputData)
	publicKey, err := key.ToPublic()
	if err != nil {
		return nil, err
	}
	publicKeyRing, err := crypto.NewKeyRing(publicKey)
	if err != nil {
		return nil, err
	}
	result, err := publicKeyRing.Encrypt(message, nil)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
