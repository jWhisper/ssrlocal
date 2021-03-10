package project

import (
	"log"

	"github.com/jWhisper/ssrlocal/configs"
	"github.com/jWhisper/ssrlocal/internal/proxy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCfg string
var _ configs.Cnf = (*viper.Viper)(nil)

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
	viper.SetConfigFile(startCfg)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read conf: %s\n", err)
	}
	opt := proxy.GetCnfOption()
	server, err := proxy.NewServer(opt...)
	if err != nil {
		log.Fatalf("failed to get a server, check config", err)
		return
	}
	server.ListenTCP()
}
