package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	Server  ServerConf
}

type LoggerConf struct {
	Level string `yaml:"level"`
}

type StorageConf struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig(configPath string) (config *Config, err error) {
	conf := new(Config)
	file, err := os.Open(configPath)
	if err != nil {
		return conf, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return conf, err
	}
	return
}

// TODO
