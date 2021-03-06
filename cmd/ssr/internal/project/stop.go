package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop ssrclient",
	Long:  `stop ssrclient`,

	Run: func(cmd *cobra.Command, args []string) {
		impStop()
	},
}

func impStop() {
	fmt.Println("Hang in the air, you can use supervisor also.")
}
