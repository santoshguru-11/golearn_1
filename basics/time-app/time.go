package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		// Create a new window
		window := ui.NewWindow("Time Display", 200, 100, false)

		// Create a new label to display the time
		label := ui.NewLabel("")

		// Set up the window layout
		window.SetChild(label)

		// Start a goroutine to continuously update the time
		go func() {
			for {
				// Get the current time
				currentTime := time.Now().Format("15:04:05")

				// Update the label with the current time
				ui.QueueMain(func() {
					label.SetText(currentTime)
				})

				// Wait for 1 second before updating again
				time.Sleep(time.Second)
			}
		}()

		// Set up the function to close the window when the user closes it
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		// Show the window
		window.Show()
	})

	// Handle any errors
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
