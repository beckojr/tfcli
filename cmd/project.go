/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"tfcli/tfe"

	"github.com/hashicorp/jsonapi"
	"github.com/spf13/cobra"
)

var p tfe.Project

// var client, err = tfe.NewClient(&tfe.Config{
// 	Token: "YOUR_TOKEN",
// })

// var client *tfe.Client

// projectCmd represents the project command
var organization string
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage a Terraform Enterprise project",
	Long:  `Manage a Terraform Enterprise project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Long:  `List projects`,
	Run: func(cmd *cobra.Command, args []string) {
		var buff bytes.Buffer
		p = tfe.NewProject(client)
		projects, err := p.List(organization, "", 10)
		if err != nil {
			fmt.Println(err)
		}
		err = jsonapi.MarshalPayload(&buff, projects.Items)
		if err != nil {
			fmt.Println(err)
		}
		prettyPrintJSON(buff.String())
	},
}

var prCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Long:  `Create a project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create project called")
	},
}

var prShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a project",
	Long:  `Show a project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Show project called")
	},
}

var prUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project",
	Long:  "Update a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Update project called")
	},
}

var prDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	Long:  `Delete a project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete project called")
	},
}

func init() {
	// p = tfe.NewProject(client)
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(prListCmd)
	projectCmd.AddCommand(prCreateCmd)
	projectCmd.AddCommand(prShowCmd)
	projectCmd.AddCommand(prUpdateCmd)
	projectCmd.AddCommand(prDeleteCmd)

	projectCmd.PersistentFlags().StringVarP(&organization, "organization", "o", "", "The organization to use")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
