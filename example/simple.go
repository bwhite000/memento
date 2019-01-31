package main

import (
	"log"

	"github.com/bwhite000/memento"
)

func main() {
	// Create a new Memento simple data store.
	mem, err := memento.NewMemento("user_prefs", "./memento")
	if err != nil {
		log.Fatal(err)
	}

	// Read the value.
	isFirstRun := mem.GetBool("isFirstRun", true)
	log.Printf("isFirstRun: %v", isFirstRun) // isFirstRun: true

	// Set the value.
	err = mem.SetBool("isFirstRun", false)
	if err != nil {
		log.Fatal(err)
	}

	// Read the previously set value.
	isFirstRun = mem.GetBool("isFirstRun", true)
	log.Printf("isFirstRun: %v", isFirstRun) // isFirstRun: false
}
