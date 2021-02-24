package vault

import (
	"github.com/hashicorp/vault/api"
)

type VaultConfig struct {
	apiConfig *api.Config
}

type VaultConfigInput struct {
	Address  string
	CAPath   string
	CertPath string
	KeyPath  string
	Insecure bool
}

func NewVaultConfig(cfg VaultConfigInput) (*VaultConfig, error) {
	// https://github.com/hashicorp/vault/blob/master/api/client.go
	config := api.DefaultConfig()
	config.Address = cfg.Address

	tls := &api.TLSConfig{
		CAPath:     cfg.CAPath,
		ClientCert: cfg.CertPath,
		ClientKey:  cfg.KeyPath,
		Insecure:   cfg.Insecure,
	}

	if err := config.ConfigureTLS(tls); err != nil {
		return nil, err
	}

	return &VaultConfig{apiConfig: config}, nil
}
