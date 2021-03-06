package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/takumakume/kubectl-vault-client-certificate/config"
	"github.com/takumakume/kubectl-vault-client-certificate/printer"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long:  "get",
	RunE:  func(cmd *cobra.Command, args []string) error { return nil },
}

var getServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  "server",
	RunE:  func(cmd *cobra.Command, args []string) error { return getServerList() },
}

var getIssuerCmd = &cobra.Command{
	Use:   "issuer",
	Short: "issuer",
	Long:  "issuer",
	RunE:  func(cmd *cobra.Command, args []string) error { return getIssuerList() },
}

func init() {
	getCmd.AddCommand(getIssuerCmd)
	getCmd.AddCommand(getServerCmd)
	rootCmd.AddCommand(getCmd)
}

func getIssuerList() error {
	p := printer.NewPrinter(os.Stdout, os.Stderr)
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}
	header := []string{
		"NAME",
		"VAULT_SERVER_NAME",
		"COMMON_NAME",
		"PATH",
		"TTL",
	}
	data := [][]string{}

	for _, c := range cfg.VaultIssuers {
		data = append(data, []string{
			c.Name,
			c.VaultServerName,
			c.CommonName,
			c.Path,
			c.TTL,
		})
	}

	p.Print(header, data, []error{})

	return nil
}

func getServerList() error {
	p := printer.NewPrinter(os.Stdout, os.Stderr)
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}
	header := []string{
		"NAME",
		"ADDRESS",
		"NAMESPACE",
	}
	data := [][]string{}

	for _, c := range cfg.VaultServers {
		data = append(data, []string{
			c.Name,
			c.Address,
			c.Namespace,
		})
	}

	p.Print(header, data, []error{})

	return nil
}
