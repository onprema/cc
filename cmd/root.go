package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "cc",
	Short: "Optimize any project for Claude Code development",
	Long: `CC is a CLI tool that adds Claude Code optimization to any project.
It creates .claude/ configurations, documentation, CI/CD workflows, and examples
to make your project work seamlessly with Claude Code.

Examples:
  cc init                                   # Add Claude Code optimization to current project
  cc init --github=username                # Add GitHub integration
  cc init --description="My awesome project" # Add project description`,
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cc-init.yaml)")
	rootCmd.PersistentFlags().Bool("dry-run", false, "show what would be created without creating")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	viper.BindPFlag("dry-run", rootCmd.PersistentFlags().Lookup("dry-run"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cc")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}