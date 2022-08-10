package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger  LoggerConf  `yaml:"logger"`
	Storage StorageConf `yaml:"storage"`
	Server  ServerConf  `yaml:"server"`
}

type LoggerConf struct {
	Level string `yaml:"level"`
}

type StorageConf struct {
	Driver    string `yaml:"driver"`
	Dsn       string `yaml:"dsn"`
	Migration string `yaml:"migration"`
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

func (c *Config) GetDriverName() string {
	return c.Storage.Driver
}

func (c *Config) GetDataSourceName() string {
	return c.Storage.Dsn
}

func (c *Config) GetMigrationDir() string {
	return c.Storage.Migration
}
