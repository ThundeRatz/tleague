package cmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var VERSION = "DEV"

var (
	cfgFile string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:     "tleague",
	Short:   "tleague has test utilities for ThunderLeague",
	Version: VERSION,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(onInit)

	rootCmd.PersistentFlags().StringP("token", "t", "", "Digital Ocean API token")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (by default searches for .tleague on current dir and home)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

func onInit() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)

		viper.SetConfigType("yaml")
		viper.SetConfigName(".tleague")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	cobra.CheckErr(err)
}

func ifVerbose(o string) {
	if verbose {
		fmt.Println(o)
	}
}
