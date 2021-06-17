package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Project struct {
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Version string `yaml:"version"`
}

type Grpc struct {
	Address string `yaml:"address"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
	SslMode  string `yaml:"sslmode"`
	Driver   string `yaml:"driver"`
}

type Metrics struct {
	Address string `yaml:"address"`
	Path    string `yaml:"path"`
}

type Kafka struct {
	Topic string `yaml:"topic"`
}

type Getway struct {
	Address string `yaml:"address"`
}

type Config struct {
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Metrics  Metrics  `yaml:"metrics"`
	Kafka    Kafka    `yaml:"kafka"`
	Getway   Getway   `yaml:"getway"`
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
