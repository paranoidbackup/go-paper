package crypt

import (
	"errors"
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/armor"
	"github.com/ProtonMail/gopenpgp/v2/constants"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

type Encrypter interface {
	Encrypt(input EncryptInput) (*EncryptOutput, error)
	EncryptNewProject(input EncryptNewProjectInput) (*EncryptOutput, error)
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

func (e *EncrypterImpl) Encrypt(input EncryptInput) (*EncryptOutput, error) {
	if len(input.PublicKeys()) == 0 {
		return nil, errors.New("no public keys provided")
	}
	docId, err := e.idGen.GenerateID()
	if err != nil {
		return nil, err
	}

	data := input.Data()
	for _, publicKey := range input.PublicKeys() {
		data, err = e.encryptRound(publicKey, data)
		if err != nil {
			return nil, err
		}
	}

	dataArmored, err := armor.ArmorWithType(data, constants.PGPMessageHeader)
	if err != nil {
		return nil, err
	}

	return &EncryptOutput{
		projectID:  input.ProjectID(),
		docID:      docId,
		publicKeys: input.PublicKeys(),
		data:       dataArmored,
	}, nil
}

func (e *EncrypterImpl) EncryptNewProject(input EncryptNewProjectInput) (*EncryptOutput, error) {
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
		result, err := e.encryptRoundWithKeyGen(round, projectId, data)
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

	return &EncryptOutput{
		projectID:   projectId,
		docID:       docId,
		publicKeys:  publicKeys,
		privateKeys: privateKeys,
		passphrases: passphrases,
		data:        dataArmored,
	}, nil
}

func (e *EncrypterImpl) encryptRound(publicKeyArmored string, inputData []byte) ([]byte, error) {
	publicKey, err := crypto.NewKeyFromArmored(publicKeyArmored)
	if err != nil {
		return nil, err
	}
	return e.doEncrypt(publicKey, inputData)
}

func (e *EncrypterImpl) encryptRoundWithKeyGen(roundId int, projectId string, inputData []byte) (*encryptRoundWithKeyGenResult, error) {
	passphrase, err := e.passGen.GeneratePassphrase(12)
	if err != nil {
		return nil, err
	}
	key, err := e.keyGen.GenerateKey(fmt.Sprintf("%v_%v", projectId, roundId), passphrase)
	if err != nil {
		return nil, err
	}
	publicKey, err := key.ToPublic()
	if err != nil {
		return nil, err
	}
	data, err := e.doEncrypt(publicKey, inputData)
	if err != nil {
		return nil, err
	}
	privateKeyArmored, err := key.Armor()
	if err != nil {
		return nil, err
	}
	publicKeyArmored, err := publicKey.Armor()
	if err != nil {
		return nil, err
	}
	return &encryptRoundWithKeyGenResult{
		publicKey:  publicKeyArmored,
		privateKey: privateKeyArmored,
		passphrase: passphrase,
		data:       data,
	}, nil
}

func (e *EncrypterImpl) doEncrypt(publicKey *crypto.Key, inputData []byte) ([]byte, error) {
	if len(inputData) == 0 {
		return nil, nil
	}
	message := crypto.NewPlainMessage(inputData)
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
