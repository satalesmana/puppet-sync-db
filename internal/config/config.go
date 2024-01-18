package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"puppet-sync-db/pkg/env"
)

type Config struct {
	App           App     `yaml:"app"`
	MongoDBRemote MongoDB `yaml:"mongo_db_remote"`
	MongoDBLocal  MongoDB `yaml:"mongo_db_local"`
}

type App struct {
	ENV      string `yaml:"env"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	TimeZone string `yaml:"timezone"`
}

type MongoDB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func New(repoName string) (*Config, error) {
	filename := getConfigFile(repoName, env.ServiceEnv())
	return newWithFile(filename)
}

func getConfigFile(repoName, env string) string {
	var (
		filename = fmt.Sprintf("config.%s.yaml", env)
	)
	return filepath.Join("files/etc", repoName, filename)
}

func newWithFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil

}
