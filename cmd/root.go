package cmd

import (
	"fmt"
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
	kubeconfig string
)

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "c", kubeconfigPath(), "kubeconfig path (default: '~/.kube/config')Overwrite with the 'KUBECONFIG' env")
}

func kubeconfigPath() string {
	kubeconfig := "~/.kube/config"
	kubeconfigFromEnv := os.Getenv("KUBECONFIG")

	if kubeconfigFromEnv != "" {
		kubeconfig = kubeconfigFromEnv
	}

	return kubeconfig
}
