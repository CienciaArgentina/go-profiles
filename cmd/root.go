package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/internal/http/rest"
	"github.com/CienciaArgentina/go-profiles/utils"

	"github.com/spf13/cobra"
)

var (
	cfgFile            string
	printConfigExample bool

	cfg config.Configuration

	rootCmd = &cobra.Command{
		Use:   "profiles",
		Short: "Profiles resource",
		PreRun: func(cmd *cobra.Command, args []string) {
			defaults := map[string]string{
				"Server.Port": "8080",
				"Server.Host": "localhost",
			}

			utils.InitConfiguration(&cfg, "user-profiles", "PROFILES", cfgFile, defaults)
		},
		Run: func(cmd *cobra.Command, args []string) {
			address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
			log.Fatal(rest.InitRouter(&cfg).Run(address))
		},
	}
)

// Execute is the CLI entry point
func Execute() {
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}

func init() {
	// init global flags
	rootCmd.
		PersistentFlags().
		BoolVarP(
			&cfg.Verbose,
			"verbose",
			"v",
			false,
			"Print verbose info")

	rootCmd.
		PersistentFlags().
		StringVarP(
			&cfgFile,
			"config",
			"c",
			"",
			`Configure resource using the provided yaml file`)

	rootCmd.
		PersistentFlags().
		BoolVarP(
			&printConfigExample,
			"printConfigExample",
			"p",
			false,
			`Print out a configuration template`)
}
