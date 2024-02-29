package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rcastrejon/multicast-channels/pkg/multicast"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		addrs := map[string]string{
			"foo": "224.0.0.250:9999",
			"bar": "224.0.0.249:9999",
			"baz": "224.0.0.248:9999",
			"qux": "224.0.0.247:9999",
		}

		srv, err := multicast.NewMulticastServer(addrs)
		if err != nil {
			log.Fatal("Error creating multicast server:", err)
		}
		defer srv.Close()

		m := initialModel(srv)
		if _, err := tea.NewProgram(m).Run(); err != nil {
			log.Fatal("Error running program:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
