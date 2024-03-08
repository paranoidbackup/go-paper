package task

import (
	"os"
	"paranoidbackup/go-paper/internal/infra/log"
	"paranoidbackup/go-paper/internal/step/crypt"
	"paranoidbackup/go-paper/internal/step/document"
	"paranoidbackup/go-paper/internal/step/project"
	"paranoidbackup/go-paper/internal/step/qrcode"
	"path"
)

type BackupTask struct {
	logger           *log.Logger
	projectRepo      project.Repo
	encrypter        crypt.Encrypter
	qrCodeGenerator  qrcode.Generator
	documentExporter document.Exporter
}

func NewBackupTask(
	logger *log.Logger,
	projectRepo project.Repo,
	encrypter crypt.Encrypter,
	qrCodeGenerator qrcode.Generator,
	documentExporter document.Exporter,
) (*BackupTask, error) {
	return &BackupTask{
		logger:           logger,
		projectRepo:      projectRepo,
		encrypter:        encrypter,
		qrCodeGenerator:  qrCodeGenerator,
		documentExporter: documentExporter,
	}, nil
}

func (t *BackupTask) BackupWithNewProject(dataPath string) error {
	var input Input
	var err error

	input.docName = path.Base(dataPath)
	input.data, err = os.ReadFile(dataPath)
	if err != nil {
		return err
	}
	doc, err := t.encrypter.EncryptNewProject(&input)
	if err != nil {
		return err
	}
	docWithQRs, err := t.qrCodeGenerator.Generate(doc)
	if err != nil {
		return err
	}
	err = t.documentExporter.Export(docWithQRs)
	if err != nil {
		return err
	}
	err = t.projectRepo.SavePublic(doc)
	if err != nil {
		return err
	}
	err = t.projectRepo.SavePrivate(doc)
	if err != nil {
		return err
	}

	return nil
}

func (t *BackupTask) BackupWithExistingProject(projectId string, dataPath string) error {
	var input Input
	var err error

	p, err := t.projectRepo.LoadPublic(projectId)
	if err != nil {
		return err
	}

	input.projectId = p.ProjectID()
	input.publicKeys = p.PublicKeys()

	input.data, err = os.ReadFile(dataPath)
	if err != nil {
		return err
	}

	doc, err := t.encrypter.Encrypt(&input)
	if err != nil {
		return err
	}
	docWithQRs, err := t.qrCodeGenerator.Generate(doc)
	if err != nil {
		return err
	}
	err = t.documentExporter.Export(docWithQRs)
	if err != nil {
		return err
	}

	return nil
}
