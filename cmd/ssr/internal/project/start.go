package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCfg string

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start ssrclient as a proxy",
	Long:  `start ssrclient as a proxy`,

	Run: func(cmd *cobra.Command, args []string) {
		impStart()
	},
}

func init() {
	StartCmd.Flags().StringVarP(&startCfg, "config", "c", "", "start ssr with config")
}

func impStart() {
	fmt.Println(startCfg)
}
