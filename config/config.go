package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-playground/validator"
	"github.com/goccy/go-yaml"
)

type Config struct {
	path         string
	VaultServers []VaultServerConfig `yaml:"vault"`
	VaultPKIs    []VaultPKIConfig    `yaml:"vault_pki"`
}

type VaultServerConfig struct {
	Name      string               `yaml:"name" validate:"required"`
	Address   string               `yaml:"address" validate:"url,required"`
	Namespace string               `yaml:"namespace,omitempty"`
	TLS       VaultServerTLSConfig `yaml:"tls,omitempty"`
}

type VaultServerTLSConfig struct {
	Cert       string `yaml:"cert,omitempty"`
	Key        string `yaml:"key,omitempty"`
	CACert     string `yaml:"ca_cert,omitempty"`
	SkipVerify bool   `yaml:"skip_verify,omitempty"`
}

type VaultPKIConfig struct {
	Name            string `yaml:"name" validate:"required"`
	VaultServerName string `yaml:"vault_server_name" validate:"required"`
	Path            string `yaml:"path" validate:"required"`
	CommonName      string `yaml:"common_name" validate:"required"`
	TTL             string `yaml:"ttl" validate:"required"`
}

func DefaultConfigPath() string {
	return fmt.Sprintf("%s/.kube/kubectl-vault-client-certificate/config", os.Getenv("HOME"))
}

func initConfig(path string) error {
	dirName := filepath.Dir(path)
	_, err := os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dirName, 0777)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(path)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func NewConfig(path string) (*Config, error) {
	var cfg Config
	if err := initConfig(path); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		return nil, err
	}

	cfg.path = path

	return &cfg, nil
}

func (cfg *Config) GetVaultServerConfig(name string) (*VaultServerConfig, error) {
	for _, a := range cfg.VaultServers {
		if a.Name == name {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("'%s' vault server was not found", name)
}

func (cfg *Config) validateVaultServerConfig(vaultServerConfig VaultServerConfig) error {
	v := validator.New()
	if err := v.Struct(vaultServerConfig); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) WriteVaultServerConfig(vaultServerConfig VaultServerConfig) error {
	if err := cfg.validateVaultServerConfig(vaultServerConfig); err != nil {
		return err
	}

	overwrite := false
	for i, a := range cfg.VaultServers {
		if a.Name == vaultServerConfig.Name {
			cfg.VaultServers[i] = vaultServerConfig
			overwrite = true
		}
	}

	if !overwrite {
		cfg.VaultServers = append(cfg.VaultServers, vaultServerConfig)
	}

	if err := cfg.writeSelf(); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) GetVaultPKIConfig(name string) (*VaultPKIConfig, error) {
	for _, a := range cfg.VaultPKIs {
		if a.Name == name {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("'%s' vault pki was not found", name)
}

func (cfg *Config) validateVaultPKIConfig(vaultPKIConfig VaultPKIConfig) error {
	v := validator.New()
	if err := v.Struct(vaultPKIConfig); err != nil {
		return err
	}

	_, err := cfg.GetVaultServerConfig(vaultPKIConfig.VaultServerName)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) WriteVaultPKIConfig(vaultPKIConfig VaultPKIConfig) error {
	if err := cfg.validateVaultPKIConfig(vaultPKIConfig); err != nil {
		return err
	}

	overwrite := false
	for i, a := range cfg.VaultPKIs {
		if a.Name == vaultPKIConfig.Name {
			cfg.VaultPKIs[i] = vaultPKIConfig
			overwrite = true
		}
	}

	if !overwrite {
		cfg.VaultPKIs = append(cfg.VaultPKIs, vaultPKIConfig)
	}

	if err := cfg.writeSelf(); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) writeSelf() error {
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(cfg.path, buf, os.ModeExclusive)
	if err != nil {
		return err
	}

	return nil
}
