package crypt

import "github.com/google/uuid"

type IDGenerator interface {
	GenerateID() (string, error)
}

type idGeneratorImpl struct{}

func newIDGenerator() *idGeneratorImpl {
	return &idGeneratorImpl{}
}

func (g *idGeneratorImpl) GenerateID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
