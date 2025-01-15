/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task_tracker",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task_tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Long:  "Adds a task. Usage: ./task_tracker add \"task name\"",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "Lists all task",
	Long:  "Lists all task. Usage: ./task_tracker list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a task",
	Long:  "Updates a task. Usage: ./task_tracker update 1 \"call doctor\"",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Long:  "Deletes a task. Usage: ./task_tracker delete 1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command")
	},
}
