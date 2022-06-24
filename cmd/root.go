/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"mackerel-speedtest/speedtest"
	"os"
	"time"

	"github.com/mackerelio/mackerel-client-go"
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
		fmt.Println("Executing speedtest cli")
		var result speedtest.SpeedTestResult
		if err := speedtest.Run(viper.GetString("speedtest_path"), &result); err != nil {
			return err
		}

		time, err2 := time.Parse("2006-01-02T15:04:05Z", result.Timestamp)
		if err2 != nil {
			return err2
		}
		unixTimestamp := time.Unix()

		fmt.Println("Posting metric values to Mackerel")
		client := mackerel.NewClient(viper.GetString("apikey"))
		err3 := client.PostServiceMetricValues("home-network-test", []*mackerel.MetricValue{
			{
				Name:  "speedtest.ping.latency",
				Time:  unixTimestamp,
				Value: result.Ping.Latency,
			},
			{
				Name:  "speedtest.ping.jitter",
				Time:  unixTimestamp,
				Value: result.Ping.Jitter,
			},
			{
				Name:  "speedtest.bandwidth.download",
				Time:  unixTimestamp,
				Value: result.Download.Bandwidth * 8,
			},
			{
				Name:  "speedtest.bandwidth.upload",
				Time:  unixTimestamp,
				Value: result.Upload.Bandwidth * 8,
			},
		})
		if err3 != nil {
			return err3
		}

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
