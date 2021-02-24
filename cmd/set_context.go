package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/takumakume/kubectl-vault-client-certificate/config"
	"github.com/takumakume/kubectl-vault-client-certificate/kubeconfig"
	"github.com/takumakume/kubectl-vault-client-certificate/vault"
)

var setContextCmd = &cobra.Command{
	Use:   "set-context",
	Short: "set-context",
	Long:  "set-context",
	RunE: func(cmd *cobra.Command, args []string) error {
		if argsContextName == "" {
			return errors.New("'context' option was empty")
		}

		if argsIssuerName == "" {
			return errors.New("'issuer' option was empty")
		}

		return setContext(argsContextName, argsIssuerName)
	},
}

func init() {
	setContextCmd.Flags().StringVarP(&argsContextName, "context", "c", "", "kubeconfig context name")
	setContextCmd.Flags().StringVarP(&argsIssuerName, "issuer", "i", "", "vault-client-certificate issuer name")

	rootCmd.AddCommand(setContextCmd)
}

func setContext(contextName, pkiName string) error {
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}

	issuerConfig, err := cfg.GetVaultPKIConfig(pkiName)
	if err != nil {
		return err
	}

	vaultServerConfig, err := cfg.GetVaultServerConfig(issuerConfig.VaultServerName)
	if err != nil {
		return err
	}

	kcfg, err := kubeconfig.New()
	if err != nil {
		return err
	}

	context, err := kcfg.ReadContext(contextName)
	if err != nil {
		return err
	}

	cluster, err := kcfg.ReadCluster(context.Cluster)
	if err != nil {
		return err
	}

	user, err := kcfg.ReadUser(context.AuthInfo)
	if err != nil {
		return err
	}

	vaultConfig, err := vault.NewVaultConfig(vault.VaultConfigInput{
		Address:  vaultServerConfig.Address,
		CertPath: vaultServerConfig.TLS.Cert,
		KeyPath:  vaultServerConfig.TLS.Key,
		CAPath:   vaultServerConfig.TLS.CACert,
		Insecure: vaultServerConfig.TLS.SkipVerify,
	})
	if err != nil {
		return err
	}

	vaultClient, err := vault.New(vaultConfig)
	if err != nil {
		return err
	}

	vaultToken, err := vault.GetAuthenticatedToken()
	if err != nil {
		return err
	}
	vaultClient.SetToken(vaultToken)

	secret, err := vaultClient.IssueNewCertificate(issuerConfig.Path, issuerConfig.CommonName, issuerConfig.TTL)
	if err != nil {
		return err
	}

	certBundle, err := vault.SecretToCertificateBundle(secret)
	if err != nil {
		return err
	}

	cluster.CertificateAuthorityData = []byte(certBundle.IssuingCA)
	user.ClientCertificateData = []byte(certBundle.Certificate)
	user.ClientKeyData = []byte(certBundle.PrivateKey)

	kcfg.WriteCluster(context.Cluster, cluster)
	kcfg.WriteUser(context.AuthInfo, user)

	return nil
}
