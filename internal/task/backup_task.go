package task

import (
	"paranoidbackup/go-paper/internal/infra/log"
	"paranoidbackup/go-paper/internal/step/crypt"
	"paranoidbackup/go-paper/internal/step/document"
	"paranoidbackup/go-paper/internal/step/project"
	"paranoidbackup/go-paper/internal/step/qrcode"
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

func (t *BackupTask) BackupToNewProject(dataPath string) error {
	return nil
}

func (t *BackupTask) BackupToExistingProject(projectId string, dataPath string) error {
	return nil
}
