/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/upsaurav12/taskido-cli/storage"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "command to list all the tasks that were created.",
	Long:  `command to list all the tasks that were created.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error while loading", err)
		}

		for _, task := range tasks {
			status := " "
			if task.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
