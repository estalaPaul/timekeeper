/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/estalaPaul/timekeeper/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// weekCmd represents the week command
var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Get time data for a specific week.",
	Long: `Get time data for a specific week. The first argument should be
    the year and the week number that you would like to see data for, in a
    format like "2024-5". If no argument is given, the current year and week
    is used.`,
	Run: func(cmd *cobra.Command, args []string) {
        week := ""
        if len(args) > 0 {
            matched, _ := regexp.MatchString("^[0-9]{4}-[0-9]{1,2}$", args[0])
            if ! matched {
                pterm.Println()
                pterm.Error.Println("Error: Invalid week format.")
                pterm.Println()
                return  
            }

            week = args[0]
        } else {
            currentYear, currentWeek := time.Now().ISOWeek()
            week = fmt.Sprintf("%d-%d", currentYear, currentWeek)
        }

        if ! utils.Exists(fmt.Sprintf("entries/week-%s", week)) {
            tableData := pterm.TableData{
                {"Entry", "Hours", "Minutes", "Seconds"},
                {"TOTAL", "0", "0", "0"},
            }

            pterm.Println()
            pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
            pterm.Println()
            return
        }


        entries, err := os.ReadDir(fmt.Sprintf("entries/week-%s", week))
        if err != nil {
            pterm.Error.Printf("Error reading %s directory: %s\n", week, err)
        }

        totals := map[string]int64{
            "hours": 0,
            "minutes": 0,
            "seconds": 0,
        }

        tableData := pterm.TableData{
            {"Entry", "Date", "Hours", "Minutes", "Seconds"},
        }
        for _, entry := range entries {
            if entry.IsDir() {
                continue
            }

            entryData := decodeFile(fmt.Sprintf("entries/week-%s/%s", week, entry.Name()))

            hours, err := strconv.ParseInt(entryData["hours"], 10, 64)
            if err != nil {
                pterm.Error.Printf("Error parsing hours: %s\n", err)
                os.Exit(1)
            }
            minutes, err := strconv.ParseInt(entryData["minutes"], 10, 64)
            if err != nil {
                pterm.Error.Printf("Error parsing minutes: %s\n", err)
                os.Exit(1)
            }
            seconds, err := strconv.ParseInt(entryData["seconds"], 10, 64)
            if err != nil {
                pterm.Error.Printf("Error parsing seconds: %s\n", err)
                os.Exit(1)
            }

            totals["hours"] += hours
            totals["minutes"] += minutes
            totals["seconds"] += seconds
            tableData = append(tableData, []string{
                entryData["description"], entryData["date"], entryData["hours"], entryData["minutes"], entryData["seconds"],
            })
        }

        tableData = append(tableData, []string{
            "", "", "", "", "",
        })
        tableData = append(tableData, []string {
            "TOTAL", week, strconv.Itoa(int(totals["hours"])), strconv.Itoa(int(totals["minutes"])), strconv.Itoa(int(totals["seconds"])),
        })
        pterm.Println()
        pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithData(tableData).Render()
        pterm.Println()
	},
}

func decodeFile(path string) map[string]string {
    data, err := os.ReadFile(path)
    if err != nil {
        pterm.Error.Printf("Error reading data: %s\n", err)
        os.Exit(1)
    }

    var jsonData map[string]string
    err = json.Unmarshal([]byte(data), &jsonData)
    if err != nil {
        pterm.Error.Printf("Error unmarshalling data: %s\n", err)
        os.Exit(1)
    }

    return jsonData
}

func init() {
	rootCmd.AddCommand(weekCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weekCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weekCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
