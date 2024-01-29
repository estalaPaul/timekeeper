/*
Copyright © 2024 Paul Estala pestala495@gmail.com
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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking time.",
	Long: `Start tracking time. If given
    the first argument will be used as a description.`,
	Run: func(cmd *cobra.Command, args []string) {
        if utils.Exists(fmt.Sprintf("%s/current.json", utils.GetDataDir())) {
            pterm.Println()
            pterm.Error.Println("You've already started time keeping, stop current time before starting again.")
            pterm.Println()
            return
        }

        description := ""
        if len(args) > 0 {
            description = args[0]
        }

        start := time.Now()
        data := map[string]string{
            "start": strconv.FormatInt(start.Unix(), 10),
            "description": description,
        }

        jsonData, err := json.Marshal(data)
        if err != nil {
            pterm.Error.Printf("Error marshalling data: %s\n", err)
            return
        }

        err = os.WriteFile(fmt.Sprintf("%s/current.json", utils.GetDataDir()), jsonData, 0666)
        if err != nil {
            pterm.Error.Printf("Error writing data: %s\n", err)
            return
        }

        if description != "" {
            pterm.Println()
            pterm.Success.Printf("Started %s at %s\n", description, start.Format(time.RFC822))
            pterm.Println()
        } else {
            pterm.Println()
            pterm.Success.Printf("Started tracking time at %s\n", start.Format(time.RFC822))
            pterm.Println()
        }
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
