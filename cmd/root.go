/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task_tracker/task"
	"time"

	"github.com/spf13/cobra"
)

var (
	currentTime time.Time
	tasks       task.Tasks
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
	// rootCmd.AddCommand(deleteCmd)
	currentTime = time.Now()

	for _, tk := range tasks {
		fmt.Println(tk)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Long:  "Adds a task. Usage: ./task_tracker add \"task name\"",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task.NewTask(args[0], currentTime)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "Lists all task",
	Long:  "Lists all task. Usage: ./task_tracker list",
	Run: func(cmd *cobra.Command, args []string) {
		for index, tk := range task.ParseTasksFile() {
			fmt.Printf("%d - %s - %s\n", index, tk.Status, tk.Description)
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Args:  cobra.ExactArgs(2),
	Short: "Updates a task",
	Long:  "Updates a task. Usage: ./task_tracker update 1 \"call doctor\"",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.ParseTasksFile()
		task_index, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error parsing index argument: err")
		}

		// check task index
		if task_index+1 > len(tasks) {
			log.Fatalln("Index out of range")
		}

		// update description
		updated_task := &tasks[task_index]
		updated_task.Description = args[1]
		updated_task.UpdatedAt = currentTime
		fmt.Printf("Description updated to: %s\n", updated_task.Description)
		fmt.Printf("Description updated on: %s\n", updated_task.UpdatedAt)

		task.WriteTasks(tasks)
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
