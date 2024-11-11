package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)

	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// remove entry at index

}

// separate concerns
func (j *Journal) Save(filename string) {
	err := os.WriteFile(filename, []byte(j.String()), 0644)
	if err != nil {
		fmt.Printf("Error saving journal to file: %v\n", err)
	}
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) Load(filename string) {
	// load journal from file

}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// process url
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	err := os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
	if err != nil {
		fmt.Printf("Error saving journal to file: %v\n", err)
	}
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	err := os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
	if err != nil {
		fmt.Printf("Error saving journal to file: %v\n", err)
	}
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	fmt.Println(j.entries)
	fmt.Println(j.String())

	// save to file
	filename := "journal.txt"

	SaveToFile(&j, filename)

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, filename)

}
