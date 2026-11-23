package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Boalq/CLI-Obsidian-Logger/internal/logtext"
)

// Your dir path to your vault
const note_dir = "/home/boalq/Documents/BoranPersonal/09 Calender/"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid Args\nUsage: ./x string")
		return
	}

	year, month, day := time.Now().Date()

	// Change this if you format the Date different
	// Or use something other than MarkDown
	// This has to match the file name pattern
	date := fmt.Sprintf("%02d-%02d-%d.md", day, month, year)

	notes, err := os.ReadDir(note_dir)
	if err != nil {
		fmt.Print(err)
		return
	}

	found := false

	for _, note := range notes {
		if note.Name() == date {
			found = true
			break
		}
	}

	// Carefully opening the File with needed Permissions
	// Precise Permissions for security (a bit verbosity)
	var f *os.File

	if found == false {
		f, err = os.Create(fmt.Sprintf("%s/%s", note_dir, date))
	} else {
		f, err = os.OpenFile(fmt.Sprintf("%s/%s", note_dir, date), os.O_WRONLY|os.O_APPEND, 0)
	}

	if err != nil {
		fmt.Printf("%e File couldn't be found neither created.", err)
		return
	}

	// Building the Actual Log
	var logtxt logtext.LogText

	LogTime := fmt.Sprintf("%02d:%02d", time.Now().Hour(), time.Now().Minute())

	// Header
	logtxt.Header = fmt.Sprintf("\n--- \t New Log at: %s \t ---\n\n", LogTime)

	// Body
	logtxt.Body = fmt.Sprintf("%s\n", os.Args[1])

	// Footer
	//logtxt.Footer = fmt.Sprintf("\n--- \t Log End \t ---")

	msg := logtxt.Header + logtxt.Body + logtxt.Footer

	// Writing to the Daily Note File
	fmt.Fprint(f, msg)

	// Feedback of successful transcription
	fmt.Printf("'%s' successfully written to %s", os.Args[1], date)
}
