package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
  App AppConfig `yaml:"app"`
  Db DBConfig `yaml:"db"`
}

type AppConfig struct {
  Name string `yaml:"name"`
  Port string `yaml:"port"`
  Encryption EncryptionConfig `yaml:"encryption"`
  External ExternalConfig `yaml:"external"`
}

type ExternalConfig struct {
  Google GoogleConfig `yaml:"google"`
}

type GoogleConfig struct {
  Smtp_password string `yaml:"smtp_password"`
  Smtp_sender_email string `yaml:"smtp_sender_email"`
  Smtp_sender_name string `yaml:"smtp_sender_name"`
}

type EncryptionConfig struct {
  Salt uint8 `yaml:"salt"`
  JWTSecret string `yaml:"jwt_secret"`
}

type DBConfig struct {
  Host string `yaml:"host"`
  Port string `yaml:"port"`
  User string `yaml:"user"`
  Password string `yaml:"password"`
  Name string `yaml:"name"`
  ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
  MaxIdleConnection uint8 `yaml:"max_idle_connection"`
  MaxOpenConnection uint8 `yaml:"max_open_connection"`
  MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
  MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
  config, err := os.ReadFile(filename)

  if err != nil {
    return
  }

  return yaml.Unmarshal(config, &Cfg)
}






