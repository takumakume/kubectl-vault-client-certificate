package vault

import (
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/helper/certutil"
)

type Vault struct {
	client *api.Client
	token  string
}

func NewVault(config *VaultConfig) (*Vault, error) {
	client, err := api.NewClient(config.apiConfig)
	if err != nil {
		return nil, err
	}
	return &Vault{client: client}, nil
}

func (v *Vault) SetToken(token string) {
	v.token = token
}

func (v *Vault) CreateCertBundle(path string, certConfig *VaultCertConfig) (*certutil.CertBundle, error) {
	data := map[string]interface{}{
		"common_name":  certConfig.CommonName,
		"organization": certConfig.Organization,
		"alt_names":    certConfig.AltNames,
		"ip_sans":      certConfig.IPSans,
		"ttl":          certConfig.TTL,
	}

	secret, err := v.create(path, data)
	if err != nil {
		return nil, err
	}

	cert, err := certutil.ParsePKIMap(secret.Data)
	if err != nil {
		return nil, err
	}

	return cert.ToCertBundle()
}

func (v *Vault) create(path string, data map[string]interface{}) (*api.Secret, error) {
	return v.client.Logical().Write(path, data)
}
