package app

import (
	"fmt"
	"os"
	"paranoidbackup/go-paper/internal/infra/log"
	"runtime/debug"
	"time"
)

type Container struct {
	infraLayer struct {
		logger *log.Logger
	}

	serviceLayer struct {
	}
}

func Bootstrap() (*Container, error) {
	var self Container
	var err error

	if err != nil {
		return nil, err
	}

	return self.doBootstrap()
}

func (self *Container) Halt() {
	os.Exit(0)
}

func (self *Container) HandlePanic() {
	fmt.Printf("PANIC: %s\n\n%s\n", recover(), string(debug.Stack()))
	time.Sleep(time.Second * 10)
	os.Exit(1)
}

func (self *Container) doBootstrap() (*Container, error) {
	var err error

	err = self.bootstrapInfraLayer()
	if err != nil {
		return nil, err
	}
	err = self.bootstrapServiceLayer()
	if err != nil {
		return nil, err
	}

	return self, nil
}

func (self *Container) bootstrapInfraLayer() error {
	self.infraLayer.logger = log.NewLogger(log.InfoLevel)
	return nil
}

func (self *Container) bootstrapServiceLayer() error {
	var err error

	if err != nil {
		return err
	}

	return nil
}
