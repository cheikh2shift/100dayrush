package events

import "log"

// This package knows NOTHING about the repository.

// Option A: Define a type
type SaveHandler func(data string) error

// The method accepts the function as a parameter
func Trigger(data string, onSave SaveHandler) error {
	// Logic...
	log.Println("Event triggered with data:", data)
	// Execute the callback
	return onSave(data)
}
