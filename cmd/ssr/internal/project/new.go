package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCfg string

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "create a config if not exist",
	Long:  `create a config if not exist`,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		impCfg()
	},
}

func init() {
	ConfigCmd.PersistentFlags().StringVarP(&newCfg, "config", "c", "", "create a config file at")
}

func impCfg() {
	fmt.Println(newCfg)
}
