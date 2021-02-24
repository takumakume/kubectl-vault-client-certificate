package cmd

import (
	"github.com/spf13/cobra"
	"github.com/takumakume/kubectl-vault-client-certificate/config"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set",
	Long:  "set",
}

var setServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  "server",
	RunE:  func(cmd *cobra.Command, args []string) error { return setServer() },
}

var setIssuerCmd = &cobra.Command{
	Use:   "issuer",
	Short: "issuer",
	Long:  "issuer",
	RunE:  func(cmd *cobra.Command, args []string) error { return setIssuer() },
}

func init() {
	setIssuerCmd.Flags().StringVarP(&argsIssuerName, "name", "", "", "vault-client-certificate issuer name")
	setIssuerCmd.Flags().StringVarP(&argsVaultServerName, "vault-server-name", "", "", "vault-client-certificate vault server name")
	setIssuerCmd.Flags().StringVarP(&argsPath, "path", "", "", "vault issuer path")
	setIssuerCmd.Flags().StringVarP(&argsCommonName, "common-name", "", "", "common name")
	setIssuerCmd.Flags().StringVarP(&argsTTL, "ttl", "", "", "certificate ttl")

	setServerCmd.Flags().StringVarP(&argsVaultServerName, "name", "", "", "vault-client-certificate issuer name")
	setServerCmd.Flags().StringVarP(&argsVaultAddr, "address", "", "", "vault addr")
	setServerCmd.Flags().StringVarP(&argsVaultNamespace, "namespace", "", "", "vault namespace")
	setServerCmd.Flags().StringVarP(&argsCert, "cert", "", "", "vault tls client certificate path")
	setServerCmd.Flags().StringVarP(&argsKey, "key", "", "", "vault tls client key path")
	setServerCmd.Flags().StringVarP(&argsCACert, "ca", "", "", "vault tls ca path")
	setServerCmd.Flags().BoolVarP(&argsSkipVerify, "insecure", "", false, "vault tls skip verify")

	setCmd.AddCommand(setIssuerCmd)
	setCmd.AddCommand(setServerCmd)
	rootCmd.AddCommand(setCmd)
}

func setIssuer() error {
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}

	if err := cfg.WriteVaultIssuerConfig(config.VaultIssuerConfig{
		Name:            argsIssuerName,
		VaultServerName: argsVaultServerName,
		Path:            argsPath,
		CommonName:      argsCommonName,
		TTL:             argsTTL,
	}); err != nil {
		return err
	}

	return nil
}

func setServer() error {
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}

	if err := cfg.WriteVaultServerConfig(config.VaultServerConfig{
		Name:      argsVaultServerName,
		Address:   argsVaultAddr,
		Namespace: argsVaultNamespace,
		TLS: config.VaultServerTLSConfig{
			Cert:       argsCert,
			Key:        argsKey,
			CACert:     argsCACert,
			SkipVerify: argsSkipVerify,
		},
	}); err != nil {
		return err
	}

	return nil
}
