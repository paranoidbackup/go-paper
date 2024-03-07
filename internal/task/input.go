package task

type Input struct {
	projectId  string
	publicKeys []string
	data       []byte
}

func (i *Input) Data() []byte {
	return i.data
}
func (i *Input) PublicKeys() []string {
	return i.publicKeys
}
func (i *Input) ProjectID() string {
	return i.projectId
}
