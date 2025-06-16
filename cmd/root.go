package cmd

import (
	"fmt"

	"github.com/sholt0r/todo/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "Todo is a simple todo list app",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("list", "l", "default", "todo list name")
	viper.BindPFlag("ActiveList", rootCmd.PersistentFlags().Lookup("list"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		confPath, err := internal.GetConfigPath()
		cobra.CheckErr(err)

		viper.AddConfigPath(confPath)
		viper.SetConfigFile("yaml")
		viper.SetConfigName("config")

		viper.SetDefault("ActiveList", "default")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config: %w", err))
	}
}
