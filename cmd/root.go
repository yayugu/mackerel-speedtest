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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing speedtest cli")
		var result speedtest.SpeedTestResult
		if err := speedtest.Run(&result); err != nil {
			return err
		}

		now := time.Now()
		unixTimestamp := now.Unix()

		fmt.Println("Posting metric values to Mackerel")
		client := mackerel.NewClient(viper.GetString("apikey"))
		err2 := client.PostServiceMetricValues("home-network", []*mackerel.MetricValue{
			{
				Name:  "internet-speed.ping.latency",
				Time:  unixTimestamp,
				Value: result.Latency.Seconds() * 1000,
			},
			{
				Name:  "internet-speed.bandwidth.download",
				Time:  unixTimestamp,
				Value: result.DLSpeed * 1024 * 1024,
			},
			{
				Name:  "internet-speed.bandwidth.upload",
				Time:  unixTimestamp,
				Value: result.ULSpeed * 1024 * 1024,
			},
		})
		if err2 != nil {
			return err2
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
