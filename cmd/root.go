package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "vault-client-certificate",
	Short:   "This command helpers TLS client authentication of kubernetes API endpoints using vault PKI.",
	Long:    "This command helpers TLS client authentication of kubernetes API endpoints using vault PKI.",
	Version: "0.0.1",
}

var (
	argsContextName     string
	argsIssuerName      string
	argsVaultServerName string
	argsVaultAddr       string
	argsPath            string
	argsCommonName      string
	argsTTL             string
)

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}

func kubeconfigPath() string {
	kubeconfig := "~/.kube/config"
	kubeconfigFromEnv := os.Getenv("KUBECONFIG")

	if kubeconfigFromEnv != "" {
		kubeconfig = kubeconfigFromEnv
	}

	return kubeconfig
}
