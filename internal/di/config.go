package di

import (
	"io"
	"strings"

	_ "embed"

	"github.com/spf13/viper"
)

//go:embed config.dev.yml
var configDataDev string

type Config struct {
	WorkDir            string `yaml:"workDir"`
	NewProjectKeyCount int    `yaml:"newProjectKeyCount"`
}

func loadConfig() (Config, error) {
	config := Config{}
	viper.SetConfigType("yaml")

	r := io.NopCloser(strings.NewReader(configDataDev))
	err := viper.ReadConfig(r)
	if err != nil {
		return Config{}, err
	}
	defer func() { _ = r.Close() }()

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
