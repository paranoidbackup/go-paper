package crypt

import (
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
	newProjectKeyCount int
	keyGen             KeyGenerator
	idGen              IDGenerator
	passGen            PassphraseGenerator
}

type keyGenRoundResult struct {
	publicKey  string
	privateKey string
	passphrase string
}

type encryptResult struct {
	docId string
	data  string
}

type createProjectResult struct {
	projectId   string
	publicKeys  []string
	privateKeys []string
	passphrases []string
}

func NewEncrypter(newProjectKeyCount int) (*EncrypterImpl, error) {
	passGen, err := newPassphraseGenerator()
	if err != nil {
		return nil, err
	}
	return &EncrypterImpl{
		newProjectKeyCount: newProjectKeyCount,
		keyGen:             newKeyGenerator(),
		idGen:              newIDGenerator(),
		passGen:            passGen,
	}, nil
}

func (e *EncrypterImpl) Encrypt(input EncryptInput) (*EncryptOutput, error) {
	result, err := e.encrypt(input.PublicKeys(), input.Data())
	if err != nil {
		return nil, err
	}
	return &EncryptOutput{
		projectID:  input.ProjectID(),
		docID:      result.docId,
		publicKeys: input.PublicKeys(),
		data:       result.data,
	}, nil
}

func (e *EncrypterImpl) EncryptNewProject(input EncryptNewProjectInput) (*EncryptOutput, error) {
	project, err := e.createProject(e.newProjectKeyCount)
	if err != nil {
		return nil, err
	}
	encryptRes, err := e.encrypt(project.publicKeys, input.Data())
	if err != nil {
		return nil, err
	}

	return &EncryptOutput{
		projectID:   project.projectId,
		docID:       encryptRes.docId,
		publicKeys:  project.publicKeys,
		privateKeys: project.privateKeys,
		passphrases: project.passphrases,
		data:        encryptRes.data,
	}, nil
}

func (e *EncrypterImpl) createProject(keyCount int) (*createProjectResult, error) {
	if keyCount <= 0 {
		return nil, fmt.Errorf("privateKey count should be >=1, got %v", keyCount)
	}

	projectId, err := e.idGen.GenerateID()
	if err != nil {
		return nil, err
	}

	publicKeys := make([]string, 0)
	privateKeys := make([]string, 0)
	passphrases := make([]string, 0)

	for round := 0; round < keyCount; round++ {
		result, err := e.keyGenRound(round, projectId)
		if err != nil {
			return nil, err
		}
		publicKeys = append(publicKeys, result.publicKey)
		privateKeys = append(privateKeys, result.privateKey)
		passphrases = append(passphrases, result.passphrase)
	}

	return &createProjectResult{
		projectId:   projectId,
		publicKeys:  publicKeys,
		privateKeys: privateKeys,
		passphrases: passphrases,
	}, nil
}

func (e *EncrypterImpl) encrypt(publicKeys []string, inputData []byte) (*encryptResult, error) {
	docId, err := e.idGen.GenerateID()
	if err != nil {
		return nil, err
	}

	data := inputData
	for _, publicKey := range publicKeys {
		data, err = e.encryptRound(publicKey, data)
		if err != nil {
			return nil, err
		}
	}
	dataArmored, err := armor.ArmorWithType(data, constants.PGPMessageHeader)
	if err != nil {
		return nil, err
	}

	return &encryptResult{
		docId: docId,
		data:  dataArmored,
	}, nil
}

func (e *EncrypterImpl) encryptRound(publicKeyArmored string, inputData []byte) ([]byte, error) {
	if len(inputData) == 0 {
		return nil, nil
	}
	publicKey, err := crypto.NewKeyFromArmored(publicKeyArmored)
	if err != nil {
		return nil, err
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

func (e *EncrypterImpl) keyGenRound(roundId int, projectId string) (*keyGenRoundResult, error) {
	passphrase, err := e.passGen.GeneratePassphrase(12)
	if err != nil {
		return nil, err
	}
	key, err := e.keyGen.GenerateKey(fmt.Sprintf("%v_%v", projectId, roundId), passphrase)
	if err != nil {
		return nil, err
	}
	privateKeyArmored, err := key.Armor()
	if err != nil {
		return nil, err
	}
	publicKeyArmored, err := key.GetArmoredPublicKey()
	if err != nil {
		return nil, err
	}
	return &keyGenRoundResult{
		publicKey:  publicKeyArmored,
		privateKey: privateKeyArmored,
		passphrase: passphrase,
	}, nil
}
