package vault

import (
	"crypto/tls"

	"github.com/hashicorp/vault/api"
)

type Vault struct {
	client *api.Client
}

func New(cfg *VaultConfig) (*Vault, error) {
	client, err := api.NewClient(cfg.apiConfig)
	if err != nil {
		return nil, err
	}
	return &Vault{client: client}, nil
}

func (v *Vault) SetToken(token string) {
	v.client.SetToken(token)
}

func (v *Vault) IssueNewCertificate(path, commonName, ttl string) (*api.Secret, error) {
	data := map[string]interface{}{
		"common_name": commonName,
		"ttl":         ttl,
	}

	secret, err := v.client.Logical().Write(path, data)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

type CertificateBundle struct {
	Certificate string
	PrivateKey  string
	IssuingCA   string
}

func SecretToCertificateBundle(secret *api.Secret) (CertificateBundle, error) {
	var bundle CertificateBundle
	cert := secret.Data["certificate"].(string)
	key := secret.Data["private_key"].(string)
	ca := secret.Data["issuing_ca"].(string)

	_, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		return bundle, err
	}

	bundle.Certificate = cert
	bundle.PrivateKey = key
	bundle.IssuingCA = ca

	return bundle, nil
}
