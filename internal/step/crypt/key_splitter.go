package crypt

type KeySplitter interface {
	Split(keyArmored string, shards int) ([]string, error)
}

type keySplitterImpl struct{}

func newKeySplitter() *keySplitterImpl {
	return &keySplitterImpl{}
}

func (s *keySplitterImpl) Split(keyArmored string, shards int) ([]string, error) {
	return nil, nil // TODO
}
