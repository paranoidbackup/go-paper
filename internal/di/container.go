package di

import (
	"fmt"
	"os"
	"paranoidbackup/go-paper/internal/infra/log"
	"paranoidbackup/go-paper/internal/step/crypt"
	"paranoidbackup/go-paper/internal/step/document"
	"paranoidbackup/go-paper/internal/step/project"
	"paranoidbackup/go-paper/internal/step/qrcode"
	"paranoidbackup/go-paper/internal/task"
	"runtime/debug"
	"time"
)

type Container struct {
	config Config

	infraLayer struct {
		logger *log.Logger
	}

	stepLayer struct {
		projectRepo      project.Repo
		encrypter        crypt.Encrypter
		qrCodeGenerator  qrcode.Generator
		documentExporter document.Exporter
	}

	taskLayer struct {
		backupTask *task.BackupTask
	}
}

func Bootstrap(workDir string) (*Container, error) {
	var c Container
	var err error

	c.config, err = loadConfig()
	if err != nil {
		return nil, err
	}

	if workDir != "" {
		c.config.WorkDir = workDir
	}

	return c.doBootstrap()
}

func (c *Container) BackupTask() *task.BackupTask {
	return c.taskLayer.backupTask
}

func (c *Container) Halt() {
	os.Exit(0)
}

func (c *Container) HandlePanic() {
	fmt.Printf("PANIC: %s\n\n%s\n", recover(), string(debug.Stack()))
	time.Sleep(time.Second * 10)
	os.Exit(1)
}

func (c *Container) doBootstrap() (*Container, error) {
	var err error

	err = c.bootstrapInfraLayer()
	if err != nil {
		return nil, err
	}
	err = c.bootstrapStepLayer()
	if err != nil {
		return nil, err
	}
	err = c.bootstrapTaskLayer()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Container) bootstrapInfraLayer() error {
	c.infraLayer.logger = log.NewLogger(log.InfoLevel)
	return nil
}

func (c *Container) bootstrapStepLayer() error {
	var err error

	c.stepLayer.projectRepo, err = project.NewRepo(c.config.WorkDir)
	if err != nil {
		return err
	}
	c.stepLayer.encrypter, err = crypt.NewEncrypter()
	if err != nil {
		return err
	}
	c.stepLayer.qrCodeGenerator, err = qrcode.NewGenerator()
	if err != nil {
		return err
	}
	c.stepLayer.documentExporter, err = document.NewHtmlExporter(c.config.WorkDir)
	if err != nil {
		return err
	}

	return nil
}

func (c *Container) bootstrapTaskLayer() error {
	var err error

	c.taskLayer.backupTask, err = task.NewBackupTask()
	if err != nil {
		return err
	}

	return nil
}
