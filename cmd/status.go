/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/estalaPaul/timekeeper/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        jsonData := utils.GetCurrentEntry()
        description := jsonData["description"]
        _, timeElapsed := utils.GetElapsedTime(jsonData["start"])

        if description != "" {
            pterm.Println()
            pterm.Info.Printf(
                "Currently tracking %s with %d hours, %d minutes, and %d seconds.\n",
                description,
                timeElapsed["hours"],
                timeElapsed["minutes"],
                timeElapsed["seconds"],
            )
            pterm.Println()
        } else {
            pterm.Println()
            pterm.Info.Printf(
                "Currently tracking with %d hours, %d minutes, and %d seconds.\n",
                timeElapsed["hours"],
                timeElapsed["minutes"],
                timeElapsed["seconds"],
            )
            pterm.Println()
        }
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
