package main

import (
	"os"

	"github.com/jWhisper/ssrlocal/cmd/ssr/internal/project"
	"github.com/jWhisper/ssrlocal/pkg/log"
	"github.com/spf13/cobra"
)

var logger log.Logger

var (
	version string = "v20210306.0.1"
	rootCmd        = &cobra.Command{
		Use:     "ssr",
		Short:   "ssr: a command line tool for ssr",
		Long:    `ssr: a command line tool for ssr`,
		Version: version,
	}
)

func init() {
	logger = log.WithLevelAndMeta(log.DefaultLogger, log.LvInfo, "ssrlocal:")
	rootCmd.AddCommand(project.ConfigCmd)
	rootCmd.AddCommand(project.StartCmd)
	rootCmd.AddCommand(project.StopCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Print(err)
		os.Exit(1)
	}
}
