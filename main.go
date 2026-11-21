package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	// Your dir path to your vault
	note_dir := "/home/boalq/Documents/BoranPersonal/09 Calender/"

	if len(os.Args) != 2 {
		fmt.Println("Invalid Args\nUsage: ./x string")
		return
	}

	year, month, day := time.Now().Date()

	// Change this if you format the Date different
	// Or use something other than MarkDown
	// This has to match the file name pattern
	date := fmt.Sprintf("%d-%d-%d.md", day, month, year)

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
	var text bytes.Buffer

	LogTime := fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute())

	// Header
	fmt.Fprintf(&text, "\n--- \t New Log at: %s \t ---\n\n", LogTime)

	// Body
	fmt.Fprintf(&text, "%s\n", os.Args[1])

	// Footer
	//fmt.Fprintf(&text, "\n--- \t Log End \t ---")

	// Writing to the Daily Note File
	fmt.Fprint(f, text.String())

	// Feedback of successful transcription
	fmt.Printf("'%s' successfully written to %s", os.Args[1], date)
}
