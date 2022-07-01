/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"mackerel-speedtest/internal/mackerel"
	"mackerel-speedtest/internal/speedtest"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mackerel-speedtest",
	Short: "Post ookla's Speedtest CLI results to Mackerel as service metrics",
	Long: `"mackerel-speedtest" is a command to post ookla's speedtest CLI results
to Mackerel as service metrics.
"Speedtest CLI" (https://www.speedtest.net/apps/cli) is required.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		s := speedtest.SpeedTest{
			Path:     viper.GetString("speedtest_path"),
			ServerId: viper.GetUint64("speedtest_server_id"),
		}
		if err := s.IsInstalled(); err != nil {
			return err
		}

		fmt.Println("Executing speedtest cli")
		if err := s.Run(); err != nil {
			return err
		}

		fmt.Println("Posting metric values to Mackerel")
		m := mackerel.NewMackerelClient(viper.GetString("apikey"), viper.GetString("service_name"))
		m.PostSpeedtestMetric(s.Result)
		fmt.Println("Complete!")

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mackerel-speedtest.conf)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".mackerel-speedtest" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".mackerel-speedtest.conf")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
