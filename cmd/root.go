package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/internal/http/rest"
	"github.com/CienciaArgentina/go-profiles/utils"
	"gopkg.in/yaml.v2"

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
				"DB.Hostname": "localhost",
				"DB.Port":     "27017",
				"DB.Database": "profile",
				"DB.Username": "profile",
				"DB.Password": "profile",
			}

			utils.InitConfiguration(&cfg, "user-profiles", "PROFILES", cfgFile, defaults)

			if printConfigExample {
				sample, _ := yaml.Marshal(cfg)
				fmt.Println(string(sample))
				os.Exit(0)
			}

		},
		Run: func(cmd *cobra.Command, args []string) {
			address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
			router, finish := rest.InitRouter(&cfg)
			defer finish()

			log.Fatal(router.Run(address))
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
