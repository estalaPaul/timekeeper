package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

func GetCurrentEntry() map[string]string {
    if ! Exists(fmt.Sprintf("%s/current.json", GetDataDir())) {
        pterm.Println()
        pterm.Error.Println("There is no active time entry.")
        pterm.Println()
        os.Exit(1)
    }

    data, err := os.ReadFile(fmt.Sprintf("%s/current.json", GetDataDir()))
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

func Exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }

    if ! errors.Is(err, os.ErrNotExist) {
        pterm.Error.Printf("Error checking if %s exists: %s\n", path, err)
        os.Exit(1)
    }

    return false
}

func GetElapsedTime(timestamp string) (endTime time.Time, times map[string]int64) {
    startTime, err := strconv.ParseInt(timestamp, 10, 64)
    endTime = time.Now()
    if err != nil {
        pterm.Error.Printf("Error parsing start time: %s\n", err)
        os.Exit(1)
    }

    secondsPassed := endTime.Unix() - startTime
    if secondsPassed < 60 {
        return endTime, map[string]int64{
            "seconds": secondsPassed,
            "minutes": 0,
            "hours": 0,
        }
    }

    minutesPassed := secondsPassed / 60
    secondsPassed = secondsPassed - (minutesPassed * 60)
    if minutesPassed < 60 {
        return endTime, map[string]int64{
            "seconds": secondsPassed,
            "minutes": minutesPassed,
            "hours": 0,
        }
    }

    hoursPassed := minutesPassed / 60
    minutesPassed = minutesPassed - (hoursPassed * 60)
    return endTime, map[string]int64{
        "seconds": secondsPassed,
        "minutes": minutesPassed,
        "hours": hoursPassed,
    }
}

func GetDataDir() string {
    return "/home/pulpo/Documents/personal/repos/timekeeper/"
}
