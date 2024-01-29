/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/estalaPaul/timekeeper/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current time tracking session.",
	Long: `Stop the current time tracking session.`,
	Run: func(cmd *cobra.Command, args []string) {
        jsonData := utils.GetCurrentEntry()
        description := jsonData["description"]
        endTime, timeElapsed := utils.GetElapsedTime(jsonData["start"])
        saveEntry(endTime, timeElapsed["seconds"], timeElapsed["minutes"], timeElapsed["hours"], description)
	},
}

func saveEntry(endTime time.Time, seconds int64, minutes int64, hours int64, description string) {
    data := map[string]string{
        "date": endTime.Format("2006-01-02"),
        "description": description,
        "hours": strconv.Itoa(int(hours)),
        "minutes": strconv.Itoa(int(minutes)),
        "seconds": strconv.Itoa(int(seconds)),
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        pterm.Error.Printf("Error marshalling data: %s\n", err)
        return
    }

    year, week := endTime.ISOWeek()
    weekDirectory := fmt.Sprintf("week-%d-%d", year, week)
    err = os.MkdirAll(fmt.Sprintf("%s/entries/%s", utils.GetDataDir(), weekDirectory), os.ModePerm)
    if err != nil {
        pterm.Error.Printf("Error creating week directory: %s\n", err)
        return
    }

    err = os.WriteFile(fmt.Sprintf("%s/entries/%s/entry-%s.json", utils.GetDataDir(), weekDirectory, endTime.Format(time.RFC3339)), jsonData, 0666)
    if err != nil {
        pterm.Error.Printf("Error writing data: %s\n", err)
        return
    }

    e := os.Remove(fmt.Sprintf("%s/current.json", utils.GetDataDir())) 
    if e != nil { 
        pterm.Error.Printf("Error deleting current entry file: %s\n", err)
        return
    }

    if description != "" {
        pterm.Println()
        pterm.Info.Printf("Saved %s with %d hours, %d minutes, and %d seconds.\n", description, hours, minutes, seconds)
        pterm.Println()
    } else {
        pterm.Println()
        pterm.Info.Printf("Saved entry with %d hours, %d minutes, and %d seconds.\n", hours, minutes, seconds)
        pterm.Println()
    }
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
