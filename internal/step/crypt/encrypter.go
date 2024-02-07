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

func NewEncrypter() (*EncrypterImpl, error) {
	passGen, err := newPassphraseGenerator()
	if err != nil {
		return nil, err
	}
	return &EncrypterImpl{
		keyGen:  newKeyGenerator(),
		idGen:   newIDGenerator(),
		passGen: passGen,
	}, nil
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
		result, err := e.roundWithKeyGen(round, projectId, data)
		if err != nil {
			return nil, err
		}
		publicKeys = append(publicKeys, result.publicKey)
		privateKeys = append(privateKeys, result.privateKey)
		passphrases = append(passphrases, result.passphrase)
		data = result.data
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

func (e *EncrypterImpl) roundWithKeyGen(roundId int, projectId string, inputData []byte) (*encryptRoundWithKeyGenResult, error) {
	passphrase, err := e.passGen.GeneratePassphrase(12)
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
	publicKey, err := keyObject.GetArmoredPublicKey()
	if err != nil {
		return nil, err
	}
	return &encryptRoundWithKeyGenResult{
		publicKey:  publicKey,
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
