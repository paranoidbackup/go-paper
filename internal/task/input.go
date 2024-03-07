package task

type Input struct {
	docName    string
	projectId  string
	publicKeys []string
	data       []byte
}

func (i *Input) DocName() string {
	return i.docName
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
