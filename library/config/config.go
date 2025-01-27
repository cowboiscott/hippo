package config

import (
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type Config struct {
	Database struct {
		Path string
	}
}

func NewConfig(path string) (c Config, err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := toml.NewDecoder(file)
	if err = decoder.Decode(&c); err != nil {
		return
	}

	return
}
