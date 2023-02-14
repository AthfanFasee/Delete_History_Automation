package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	// Global mouse delay
	robotgo.MouseSleep = 100

	// // Open edge
	robotgo.MoveClick(126, 273, "left", true)
	robotgo.MilliSleep(1000)

	// Get ProcessID of Edge and max window
	EdgePID := robotgo.GetPID()
	robotgo.MaxWindow(EdgePID)
	robotgo.MilliSleep(500)

	// Open history tab
	robotgo.KeyTap("delete", "control", "shift")
	robotgo.MilliSleep(2200)

	// Choose 24 hours & Delete History
	robotgo.MoveClick(708, 335, "left", false)
	robotgo.MilliSleep(1200)
	robotgo.MoveClick(680, 414, "left", false)
	robotgo.MilliSleep(200)
	robotgo.MoveClick(672, 673, "left", false)
	robotgo.MilliSleep(3700)

	// Close Edge and End Programme
	robotgo.CloseWindow()
}
