package cmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/takumakume/kubectl-vault-client-certificate/config"
	"github.com/takumakume/kubectl-vault-client-certificate/printer"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe",
	Long:  "describe",
	RunE:  func(cmd *cobra.Command, args []string) error { return nil },
}

var describeServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Invalid argument")
		}
		return describeServerList(args[0])
	},
}

var describeIssuerCmd = &cobra.Command{
	Use:   "issuer",
	Short: "issuer",
	Long:  "issuer",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Invalid argument")
		}
		return describeIssuerList(args[0])
	},
}

func init() {
	describeCmd.AddCommand(describeIssuerCmd)
	describeCmd.AddCommand(describeServerCmd)
	rootCmd.AddCommand(describeCmd)
}

func describeIssuerList(name string) error {
	p := printer.NewPrinter(os.Stdout, os.Stderr)
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}

	c, err := cfg.GetVaultIssuerConfig(name)
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{
		"NAME",
		c.Name,
	})

	data = append(data, []string{
		"VAULT SERVER NAME",
		c.VaultServerName,
	})

	data = append(data, []string{
		"COMMON NAME",
		c.CommonName,
	})

	data = append(data, []string{
		"PATH",
		c.Path,
	})

	data = append(data, []string{
		"TTL",
		c.TTL,
	})

	p.PrintNoHeader(data, []error{})

	return nil
}

func describeServerList(name string) error {
	p := printer.NewPrinter(os.Stdout, os.Stderr)
	cfg, err := config.NewConfig(config.DefaultConfigPath())
	if err != nil {
		return err
	}

	c, err := cfg.GetVaultServerConfig(name)
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{
		"NAME",
		c.Name,
	})

	data = append(data, []string{
		"ADDRESS",
		c.Address,
	})

	data = append(data, []string{
		"NAMESPACE",
		c.Namespace,
	})

	data = append(data, []string{
		"TLS CA CERT",
		c.TLS.CACert,
	})

	data = append(data, []string{
		"TLS CERT",
		c.TLS.Cert,
	})

	data = append(data, []string{
		"TLS Key",
		c.TLS.Key,
	})

	data = append(data, []string{
		"TLS SKIP VERIFY",
		strconv.FormatBool(c.TLS.SkipVerify),
	})

	p.PrintNoHeader(data, []error{})

	return nil
}
