package cmd

import (
	"fmt"
	"log"

	"github.com/rcastrejon/multicast-one-way-msg/pkg/multicast"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := multicast.NewMulticastClient(args[0])
		if err != nil {
			log.Fatal("Error creating the client: ", err)
		}

		fmt.Printf("Listening to multicast socket: %s\n\n", args[0])
		for {
			msg := c.Receive()
			fmt.Printf("Server> %s\n", msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
