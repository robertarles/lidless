package ui

import (
	"log"

	"github.com/getlantern/systray"
	"github.com/robertarles/lidless/internal/assets"
	"github.com/robertarles/lidless/internal/state"
)

var manager *state.Manager

// OnReady is called when the systray is ready
func OnReady() {
	// Create state manager
	manager = state.NewManager()

	// Sync with system on startup (FR-005)
	if err := manager.SyncWithSystem(); err != nil {
		log.Printf("Warning: Failed to sync with system state: %v", err)
		// Continue with default state
	}

	// Set initial icon and tooltip based on current state
	updateIcon(manager.IsSleepDisabled())

	// Register for state change notifications
	manager.OnStateChange(updateIcon)

	// Add menu items
	mToggle := systray.AddMenuItem("Toggle Sleep", "Enable or disable system sleep")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Exit Lidless")

	// Handle menu clicks in a goroutine
	go func() {
		for {
			select {
			case <-mToggle.ClickedCh:
				if err := manager.Toggle(); err != nil {
					log.Printf("Error toggling sleep state: %v", err)
					// TODO: Show user notification in Task 8
				}

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

// OnExit is called when the systray is exiting
func OnExit() {
	// Cleanup: No special cleanup needed for now
	// The state manager will be garbage collected
	// Sleep state is preserved (not reset on exit)
	log.Println("Lidless shutting down")
}

// updateIcon updates the menubar icon and tooltip based on sleep state
func updateIcon(sleepDisabled bool) {
	if sleepDisabled {
		// System is awake - show sun icon
		systray.SetTitle(assets.IconAwake)
		systray.SetTooltip("Sleep Disabled - Your Mac will stay awake")
	} else {
		// System can sleep - show moon icon
		systray.SetTitle(assets.IconSleep)
		systray.SetTooltip("Sleep Enabled - Normal sleep behavior")
	}
}
