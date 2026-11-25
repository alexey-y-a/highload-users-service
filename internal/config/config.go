package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

const envFileName = ".env"

type Config struct {
	env *EnvSetting
}

type EnvSetting struct {
	ServerHost string `env:"SERVER_HOST" env-default:"0.0.0.0"`
	ServerPort int    `env:"SERVER_PORT" env-default:"8080"`

	LogLevel  string `env:"LOG_LEVEL" env-default:"info"`
	LogFormat string `env:"LOG_FORMAT" env-default:"text"`
}

func Load() (*Config, error) {
	env := &EnvSetting{}

	_, err := os.Stat(envFileName)
	switch {
	case err == nil:
		readErr := cleanenv.ReadConfig(envFileName, env)
		if readErr != nil {
			return nil, fmt.Errorf("read config file %s: %w", envFileName, readErr)
		}
	case errors.Is(err, os.ErrNotExist):
		readErr := cleanenv.ReadEnv(env)
		if readErr != nil {
			return nil, fmt.Errorf("read environment: %w", readErr)
		}
	default:
		return nil, fmt.Errorf("stat config file %s: %w", envFileName, err)
	}

	return &Config{env: env}, nil
}

func (c *Config) GetServerAddr() string {
	return net.JoinHostPort(c.env.ServerHost, strconv.Itoa(c.env.ServerPort))
}

func (c *Config) GetLogLevel() string {
	return c.env.LogLevel
}

func (c *Config) GetLogFormat() string {
	return c.env.LogFormat
}
