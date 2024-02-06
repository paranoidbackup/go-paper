package crypt

type KeySplitter interface {
	Split(keyArmored string, passphrase string, shards int) ([]string, []string, error)
}

type keySplitterImpl struct{}

func newKeySplitter() *keySplitterImpl {
	return &keySplitterImpl{}
}

func (s *keySplitterImpl) Split(keyArmored string, passphrase string, shards int) ([]string, []string, error) {
	return nil, nil, nil // TODO
}
