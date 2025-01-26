package cmd

import (
	"OutQuestor/cli/cmd/commands"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "OutQuestor",
	Short: "Tool to listen outgoing network requests",
	Long:  "Tool to listen outgoing network requests",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(commands.ListenCmd)
}
