package crypt

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/armor"
	"github.com/ProtonMail/gopenpgp/v2/constants"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

type Encrypter interface {
	EncryptNewProject(input EncryptNewProjectInput) (*EncryptNewProjectOutput, error)
}

type EncrypterImpl struct {
	keyGen  KeyGenerator
	idGen   IDGenerator
	passGen PassphraseGenerator
}

type encryptRoundWithKeyGenResult struct {
	publicKey  string
	privateKey string
	passphrase string
	data       []byte
}

func NewEncrypter() *EncrypterImpl {
	return &EncrypterImpl{
		keyGen:  newKeyGenerator(),
		idGen:   newIDGenerator(),
		passGen: newPassphraseGenerator(),
	}
}

func (e *EncrypterImpl) EncryptNewProject(input EncryptNewProjectInput) (*EncryptNewProjectOutput, error) {
	if input.KeyCount() <= 0 {
		return nil, fmt.Errorf("privateKey count should be >=1, got %v", input.KeyCount())
	}

	projectId, err := e.idGen.GenerateID()
	if err != nil {
		return nil, err
	}
	docId, err := e.idGen.GenerateID()
	if err != nil {
		return nil, err
	}

	publicKeys := make([]string, 0)
	privateKeys := make([]string, 0)
	passphrases := make([]string, 0)
	data := input.Data()

	for round := 0; round < input.KeyCount(); round++ {
		roundResult, err := e.encryptRoundWithKeyGen(round, projectId, data)
		if err != nil {
			return nil, err
		}
		publicKeys = append(publicKeys, roundResult.publicKey)
		privateKeys = append(privateKeys, roundResult.privateKey)
		passphrases = append(passphrases, roundResult.passphrase)
		data = roundResult.data
	}

	dataArmored, err := armor.ArmorWithType(data, constants.PGPMessageHeader)
	if err != nil {
		return nil, err
	}

	return &EncryptNewProjectOutput{
		projectID:   projectId,
		docID:       docId,
		publicKeys:  publicKeys,
		privateKeys: privateKeys,
		passphrases: passphrases,
		data:        dataArmored,
	}, nil
}

func (e *EncrypterImpl) encryptRoundWithKeyGen(roundId int, projectId string, inputData []byte) (*encryptRoundWithKeyGenResult, error) {
	passphrase, err := e.passGen.GeneratePassphrase()
	if err != nil {
		return nil, err
	}
	keyObject, err := e.keyGen.GenerateKey(fmt.Sprintf("%v_%v", projectId, roundId), passphrase)
	if err != nil {
		return nil, err
	}
	data, err := e.doEncrypt(keyObject, inputData)
	if err != nil {
		return nil, err
	}
	privateKey, err := keyObject.Armor()
	if err != nil {
		return nil, err
	}
	return &encryptRoundWithKeyGenResult{
		privateKey: privateKey,
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
