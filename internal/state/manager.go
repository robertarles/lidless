// Package state provides sleep state management with thread-safe operations
// and callback notifications for UI updates.
package state

import (
	"fmt"
	"sync"

	"github.com/robertarles/lidless/internal/pmset"
)

// StateChangeCallback is called when the sleep state changes
type StateChangeCallback func(sleepDisabled bool)

// Manager manages the sleep disabled state with thread-safe operations
type Manager struct {
	mu            sync.RWMutex
	sleepDisabled bool
	callbacks     []StateChangeCallback
	controller    pmset.Controller
}

// NewManager creates a new state manager
func NewManager() *Manager {
	return &Manager{
		callbacks:  make([]StateChangeCallback, 0),
		controller: pmset.NewController(),
	}
}

// NewManagerWithController creates a manager with a custom pmset controller
// (useful for testing with mock controller)
func NewManagerWithController(controller pmset.Controller) *Manager {
	return &Manager{
		callbacks:  make([]StateChangeCallback, 0),
		controller: controller,
	}
}

// SyncWithSystem queries the actual system sleep state via pmset and updates
// the internal state. Only notifies callbacks if the state actually changed.
func (m *Manager) SyncWithSystem() error {
	systemState, err := m.controller.GetSleepDisabled()
	if err != nil {
		return fmt.Errorf("failed to sync with system: %w", err)
	}

	m.mu.Lock()
	changed := m.sleepDisabled != systemState
	m.sleepDisabled = systemState
	m.mu.Unlock()

	if changed {
		m.notifyCallbacks(systemState)
	}

	return nil
}

// notifyCallbacks calls all registered callbacks with the new state.
// Note: This is called while holding the lock, so callbacks must not
// call back into the Manager (would deadlock).
func (m *Manager) notifyCallbacks(sleepDisabled bool) {
	for _, cb := range m.callbacks {
		cb(sleepDisabled)
	}
}

// Toggle inverts the current sleep state, applies the change via pmset,
// and notifies all registered callbacks. Returns any error from pmset.
func (m *Manager) Toggle() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	newState := !m.sleepDisabled
	if err := m.controller.SetSleepDisabled(newState); err != nil {
		return fmt.Errorf("failed to toggle sleep: %w", err)
	}

	m.sleepDisabled = newState
	m.notifyCallbacks(newState)
	return nil
}

// OnStateChange registers a callback to be called when the sleep state changes
func (m *Manager) OnStateChange(cb StateChangeCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.callbacks = append(m.callbacks, cb)
}

// IsSleepDisabled returns the current sleep disabled state
func (m *Manager) IsSleepDisabled() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.sleepDisabled
}
