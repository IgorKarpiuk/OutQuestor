package commands

import (
	"OutQuestor/core"
	"fmt"
	"github.com/spf13/cobra"
)

var listenArgs core.ListenArgs = core.ListenArgs{}

var ListenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for all outgoing requests",
	Long:  "Listen for TCP/UPD outgoing requests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listening on channel %s\n", listenArgs.Protocol)
		core.StartListening(listenArgs)
	},
}

func init() {
	ListenCmd.Flags().StringVarP(&listenArgs.Protocol, "protocol", "p", "", "Protocol to listen (tcp or upd)")
	ListenCmd.Flags().StringVarP(&listenArgs.IpLayer, "ipLayer", "i", "", "Ip layer to listen (v4 or v6)")
	ListenCmd.Flags().BoolVar(&listenArgs.HttpOnly, "httpOnly", false, "Ip layer to listen (v4 or v6)")
	//greetCmd.MarkFlagRequired("name")
}
