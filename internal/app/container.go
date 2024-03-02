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
	var c Container
	var err error

	if err != nil {
		return nil, err
	}

	return c.doBootstrap()
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
	err = c.bootstrapServiceLayer()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Container) bootstrapInfraLayer() error {
	c.infraLayer.logger = log.NewLogger(log.InfoLevel)
	return nil
}

func (c *Container) bootstrapServiceLayer() error {
	var err error

	if err != nil {
		return err
	}

	return nil
}
