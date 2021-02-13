package vault

type VaultCertConfig struct {
	CommonName   string
	Organization string
	AltNames     string
	IPSans       string
	TTL          string
}
