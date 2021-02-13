package cmd

import "github.com/spf13/cobra"

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long:  "login",
	RunE:  func(cmd *cobra.Command, args []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
