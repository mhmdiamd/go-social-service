package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App   AppConfig   `yaml:"app"`
	Kafka KafkaConfig `yaml:"kafka"`
	Db    DBConfig    `yaml:"db"`
}

type KafkaConfig struct {
	Host string `yaml:"host"`
}

type AppConfig struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Encryption EncryptionConfig `yaml:"encryption"`
	External   ExternalConfig   `yaml:"external"`
}

type ExternalConfig struct {
	Google GoogleConfig `yaml:"google"`
}

type GoogleConfig struct {
	Smtp_password     string            `yaml:"smtp_password"`
	Smtp_sender_email string            `yaml:"smtp_sender_email"`
	Smtp_sender_name  string            `yaml:"smtp_sender_name"`
	Drive             GoogleDriveConfig `yaml:"drive"`
}

type GoogleDriveConfig struct {
	Type                    string `yaml:"type"`
	ProjectId               string `yaml:"project_id"`
	PrivateKeyId            string `yaml:"private_key_id"`
	PrivateKey              string `yaml:"private_key"`
	ClientEmail             string `yaml:"client_email"`
	ClientId                string `yaml:"client_id"`
	AuthUri                 string `yaml:"auth_uri"`
	TokenUri                string `yaml:"token_uri"`
	AuthProviderX509CertUrl string `yaml:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `yaml:"client_x509_cert_url"`
	UniverseDomain          string `yaml:"universe_domain"`
}

type EncryptionConfig struct {
	Salt      uint8  `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

type DBConfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	User           string                 `yaml:"user"`
	Password       string                 `yaml:"password"`
	Name           string                 `yaml:"name"`
	ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
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
