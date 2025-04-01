package logger

import (
	"os"

	"gopkg.in/yaml.v3"
)

type LogConfig struct {
	LogDir string `yaml:"log_dir"`
	Debug  bool   `yaml:"debug"`

	CommonLog LogFileConfig `yaml:"common_log"`
	ErrorLog  LogFileConfig `yaml:"error_log"`
	OutputLog LogFileConfig `yaml:"output_log"`
	Console   ConsoleConfig `yaml:"console"`
}

type LogFileConfig struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

type ConsoleConfig struct {
	Enabled bool   `yaml:"enabled"`
	Level   string `yaml:"level"`
}

func LoadConfig(path string) (*LogConfig, error) {
	config := &LogConfig{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(config.LogDir, 0755); err != nil {
		return nil, err
	}

	return config, nil
}
