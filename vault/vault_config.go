package vault

import (
	"github.com/hashicorp/vault/api"
)

type VaultConfig struct {
	apiConfig *api.Config
}

func NewVaultConfig(address, tls_ca, tls_cert, tls_key string, tls_insecure bool) (*VaultConfig, error) {
	// https://github.com/hashicorp/vault/blob/master/api/client.go
	config := api.DefaultConfig()
	config.Address = address

	tls := &api.TLSConfig{
		CAPath:     tls_ca,
		ClientCert: tls_cert,
		ClientKey:  tls_key,
		Insecure:   tls_insecure,
	}

	if err := config.ConfigureTLS(tls); err != nil {
		return nil, err
	}

	return &VaultConfig{apiConfig: config}, nil
}
