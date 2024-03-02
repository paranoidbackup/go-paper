package project

type Repo interface {
	LoadPublic(id string) (PublicProject, error)
	SavePublic(project PublicProject) error
	SavePrivate(project PrivateProject) error
}

type RepoImpl struct {
	workDir string
}

func NewRepo(workDir string) (*RepoImpl, error) {
	return &RepoImpl{
		workDir: workDir,
	}, nil
}

func (r *RepoImpl) LoadPublic(id string) (PublicProject, error) {
	return nil, nil // TODO
}

func (r *RepoImpl) SavePublic(project PublicProject) error {
	return nil // TODO
}

func (r *RepoImpl) SavePrivate(project PrivateProject) error {
	return nil // TODO
}
