package config

import (
	"fmt"
	"os"
	"path/filepath"
	model "puppet-sync-db/internal/model/config"
	"puppet-sync-db/pkg/env"

	"gopkg.in/yaml.v3"
)

func New(repoName string) (*model.Config, error) {
	filename := getConfigFile(repoName, env.ServiceEnv())
	return newWithFile(filename)
}

func getConfigFile(repoName, env string) string {
	var (
		filename = fmt.Sprintf("config.%s.yaml", env)
	)
	return filepath.Join("files/etc", repoName, filename)
}

func newWithFile(filename string) (*model.Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg model.Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil

}
