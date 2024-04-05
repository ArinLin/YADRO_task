package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	SourceURL   string `yaml:"source_url"`
	DBFile      string `yaml:"db_file"`
	StopWords   string `yaml:"stop_words"`
	ComicsCount int    `yaml:"comics_count"`
}

func ReadConfig() (*Config, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil

}
