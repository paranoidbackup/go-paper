package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

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
	data, err := os.ReadFile(fmt.Sprintf("%v/projects-public/public-%v.json", r.workDir, id))
	if err != nil {
		return nil, err
	}
	var entity Entity
	err = json.Unmarshal(data, &entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *RepoImpl) SavePublic(project PublicProject) error {
	entity := Entity{
		ID:     project.ProjectID(),
		Public: project.PublicKeys(),
	}
	data, err := json.MarshalIndent(entity, "", "\t")
	if err != nil {
		return err
	}
	err = r.mkdir(fmt.Sprintf("%v/projects-public", r.workDir))
	if err != nil {
		return err
	}
	return os.WriteFile(fmt.Sprintf("%v/projects-public/public-%v.json", r.workDir, entity.ID), data, os.ModePerm)
}

func (r *RepoImpl) SavePrivate(project PrivateProject) error {
	entity := Entity{
		ID:      project.ProjectID(),
		Public:  project.PublicKeys(),
		Private: project.PrivateKeys(),
		Pass:    project.Passphrases(),
	}
	data, err := json.MarshalIndent(entity, "", "\t")
	if err != nil {
		return err
	}
	err = r.mkdir(fmt.Sprintf("%v/projects-private", r.workDir))
	if err != nil {
		return err
	}
	return os.WriteFile(fmt.Sprintf("%v/projects-private/private-%v.json", r.workDir, entity.ID), data, os.ModePerm)
}

func (r *RepoImpl) mkdir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
