package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(-1)
	}

	now := time.Now()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: daily <command> <your text>")
		fmt.Println("Commands:")
		fmt.Println("  did     - Record something you did today")
		fmt.Println("  plan    - Record something you planned or need to plan")
		fmt.Println("  block   - Record a blocking problem")
		fmt.Println("  meeting - Record notes from a meeting or one that needs to be scheduled")
		fmt.Println("  cheat   - Display all recorded notes")
		return
	}

	switch strings.TrimSpace(args[0]) {
	case "did":
		if len(args) < 2 {
			fmt.Println("Please provide something to do.")
			return
		}
		input := strings.Join(args[1:], " ")
		appendToFile(dirname, "did", input, now)
	case "plan":
		if len(args) < 2 {
			fmt.Println("Please provide something to plan.")
			return
		}
		input := strings.Join(args[1:], " ")
		appendToFile(dirname, "plan", input, now)
	case "block":
		if len(args) < 2 {
			fmt.Println("Please provide something that is blocking you.")
			return
		}
		input := strings.Join(args[1:], " ")
		appendToFile(dirname, "block", input, now)
	case "meeting":
		if len(args) < 2 {
			fmt.Println("Please provide something from or for a meeting.")
			return
		}
		input := strings.Join(args[1:], " ")
		appendToFile(dirname, "meeting", input, now)
	case "cheat":
		printFile("did.txt")
		printFile("plan.txt")
		printFile("block.txt")
		printFile("meeting.txt")
	}
}

func appendToFile(dirPath, dataType, input string, date time.Time) {
	fullDirPath := filepath.Join(dirPath, date.Format("2006-01-02"))
	err := os.MkdirAll(fullDirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	p := filepath.Join(fullDirPath, dataType+".txt")
	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(input + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func printFile(dirPath, dataType string, date time.Time) {

}
