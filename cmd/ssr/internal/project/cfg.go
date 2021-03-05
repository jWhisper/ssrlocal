package project

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCfg string

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "create a config if not exist",
	Long:  `create a config if not exist`,
	//Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		initCfg()
	},
}

func init() {
	ConfigCmd.PersistentFlags().StringVarP(&newCfg, "config", "c", "", "create a config file at")
}

func initCfg() {
	viper.SetConfigType("yaml")
	if newCfg == "" {
		newCfg = "./config.yaml"
	}
	viper.SetConfigFile(newCfg)

	// Set default value
	viper.SetDefault("server", []string{"ss_provider_ip1"})
	viper.SetDefault("server_port", ":1020")
	viper.SetDefault("local_port", ":1080")
	viper.SetDefault("password", "pass")
	viper.SetDefault("protocol", "auth_aes128_md5")
	viper.SetDefault("protocol_param", "xxxx")
	viper.SetDefault("obfs", "plain")
	viper.SetDefault("obfs_param", "cdn.sharepointonline.com")
	viper.SetDefault("fast_open", false)
	viper.SetDefault("method", "chacha20-ietf")
	viper.SetDefault("timeout", 60)
	viper.SetDefault("udp_timeout", 60)
	viper.SetDefault("log", map[string]string{"level": "0", "log_file": ""})
	viper.SetDefault("type", "ssr")

	// uncomplete
	viper.SetDefault("speed_limit_per_con", 0)
	viper.SetDefault("speed_limit_per_user", 0)
	viper.SetDefault("additional_ports", map[string]int{})
	viper.SetDefault("additional_ports_only", false)
	viper.SetDefault("dns_ipv6", false)
	viper.SetDefault("connect_verbose_info", 0)
	viper.SetDefault("redirect", "")

	// Write as file if not exist
	viper.SafeWriteConfigAs(newCfg)
	fmt.Printf("has .yaml conf at: %s\n", viper.ConfigFileUsed())
}
