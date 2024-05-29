/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"tfcli/internal"

	"github.com/hashicorp/jsonapi"
	"github.com/spf13/cobra"
)

// var client, err = tfcli.NewClient(&tfcli.Config{
// 	Token: "YOUR_TOKEN",
// })

// var client *tfcli.Client

// projectCmd represents the project command
var organization string
var jsonFormat bool
var name string
var description string

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
		projects, err := project.List(organization, "", 10)
		if err != nil {
			log.Fatal(err)
		}
		if jsonFormat {
			err = jsonapi.MarshalPayloadWithoutIncluded(&buff, projects.Items)
			if err != nil {
				fmt.Println(err)
			}
			internal.PrettyPrintJSON(buff.String())
		} else {
			for _, p := range projects.Items {
				fmt.Printf("%s\n", p.Name)
			}
		}
	},
}

var prCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Long:  `Create a project`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := project.Create(organization, name, description)
		if err != nil {
			log.Fatal(err)
		}
		if jsonFormat {
			var buff bytes.Buffer
			err = jsonapi.MarshalPayloadWithoutIncluded(&buff, p)
			if err != nil {
				log.Fatal(err)
			}
			internal.PrettyPrintJSON(buff.String())
		} else {
			fmt.Printf("%s\n", p.Name)
		}

	},
}

var prShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a project",
	Long:  `Show a project`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := project.Show(organization, name)
		if err != nil {
			log.Fatal(err)
		}
		var buff bytes.Buffer
		err = jsonapi.MarshalPayloadWithoutIncluded(&buff, p)
		if err != nil {
			log.Fatal(err)
		}
		internal.PrettyPrintJSON(buff.String())

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
		err := project.Delete(organization, name)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(prListCmd)
	projectCmd.AddCommand(prCreateCmd)
	projectCmd.AddCommand(prShowCmd)
	projectCmd.AddCommand(prUpdateCmd)
	projectCmd.AddCommand(prDeleteCmd)

	projectCmd.PersistentFlags().StringVarP(&organization, "organization", "", "", "The organization to use")
	projectCmd.PersistentFlags().BoolVarP(&jsonFormat, "json", "", false, "Output in JSON format")
	projectCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "The name of the project")
	projectCmd.PersistentFlags().StringVarP(&description, "description", "", "", "The description of the project")
}
