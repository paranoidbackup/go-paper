package project

type PublicProject interface {
	ProjectID() string
	PublicKeys() []string
}

type PrivateProject interface {
	ProjectID() string
	PublicKeys() []string
	PrivateKeys() []string
	Passphrases() []string
}

type Entity struct {
	ID      string   `json:"id,omitempty"`
	Public  []string `json:"public,omitempty"`
	Private []string `json:"private,omitempty"`
	Pass    []string `json:"pass,omitempty"`
}

func NewPublicProject(
	projectId string,
	publicKeys []string,
) Entity {
	return Entity{
		ID:     projectId,
		Public: publicKeys,
	}
}

func NewPrivateProject(
	projectId string,
	publicKeys []string,
	privateKeys []string,
	passphrases []string,
) Entity {
	return Entity{
		ID:      projectId,
		Public:  publicKeys,
		Private: privateKeys,
		Pass:    passphrases,
	}
}

func (e *Entity) ProjectID() string {
	return e.ID
}
func (e *Entity) PublicKeys() []string {
	return e.Public
}
func (e *Entity) PrivateKeys() []string {
	return e.Private
}
func (e *Entity) Passphrases() []string {
	return e.Pass
}
