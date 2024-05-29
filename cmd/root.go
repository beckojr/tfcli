/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"tfcli/tfcli"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var client *tfe.Client
var project tfcli.Project

var rootCmd = &cobra.Command{
	Use:   "tfcli",
	Short: "A CLI client for Terraform Enterprise",
	Long: `tfcli is a CLI client to interact with your Terraform Enterprise deployment.
It allows you to interact with Terraform enterprise constructs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tfcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	tfc, err := tfe.NewClient(tfe.DefaultConfig())
	if err != nil {
		fmt.Println(err)
	}
	client = tfc
	project = tfcli.NewProject(client)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
