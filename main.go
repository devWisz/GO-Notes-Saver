package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const notesFile = "notes.json"

type Note struct {
	Text string `json:"text"`
}

func loadNotes() ([]Note, error) {
	file, err := os.ReadFile(notesFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Note{}, nil
		}
		return nil, err
	}

	var notes []Note
	err = json.Unmarshal(file, &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func saveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(notesFile, data, 0644)
}

func showMenu() {
	fmt.Println("\nWelcome to GO Notes Saver ")
	fmt.Println("Choose an option:")
	fmt.Println("1. Add Note")
	fmt.Println("2. List Notes")
	fmt.Println("3. Delete Note")
	fmt.Println("4. Clear All Notes")
	fmt.Println("5. Exit")
}
func addNote() {
	var text string
	fmt.Print("Enter your note: ")
	fmt.Scanln(&text)

	notes, _ := loadNotes()
	notes = append(notes, Note{Text: text})
	err := saveNotes(notes)
	if err != nil {
		fmt.Println("Failed to save note:", err)
		return
	}
	fmt.Println(" Note added successfully!")
}

func listNotes() {
	notes, _ := loadNotes()
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}
	fmt.Println("\n Your Notes:")
	for i, note := range notes {
		fmt.Printf("%d. %s\n", i+1, note.Text)
	}
}

func deleteNote() {
	listNotes()

	var input string
	fmt.Print("Enter note number to delete: ")
	fmt.Scanln(&input)

	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index <= 0 {
		fmt.Println(" Invalid input.")
		return
	}

	notes, _ := loadNotes()
	if index > len(notes) {
		fmt.Println(" No note with that number.")
		return
	}

	notes = append(notes[:index-1], notes[index:]...)
	err = saveNotes(notes)
	if err != nil {
		fmt.Println("Failed to delete note:", err)
		return
	}
	fmt.Println(" Note deleted successfully!")
}

func clearNotes() {
	err := os.Remove(notesFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Failed to clear notes:", err)
		return
	}
	fmt.Println(" All notes cleared successfully!")
}

func main() {
	for {
		showMenu()

		var choice string
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addNote()
		case "2":
			listNotes()
		case "3":
			deleteNote()
		case "4":
			clearNotes()
		case "5":
			fmt.Println("Thanks for using GO Notes Saver .Developed by Sarjak Khanal!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
