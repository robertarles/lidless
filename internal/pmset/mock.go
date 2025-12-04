package pmset

// MockController is a mock implementation of Controller for testing
type MockController struct {
	// GetSleepDisabledFunc allows configuring the return value of GetSleepDisabled
	GetSleepDisabledFunc func() (bool, error)

	// SetSleepDisabledFunc allows configuring the behavior of SetSleepDisabled
	SetSleepDisabledFunc func(disabled bool) error

	// CallLog tracks all calls made to the mock
	CallLog []string

	// CurrentState tracks the mock's internal state for simple use cases
	CurrentState bool
}

// NewMockController creates a new MockController with default behavior
func NewMockController() *MockController {
	return &MockController{
		CallLog:      make([]string, 0),
		CurrentState: false,
	}
}

// GetSleepDisabled implements Controller.GetSleepDisabled
func (m *MockController) GetSleepDisabled() (bool, error) {
	m.CallLog = append(m.CallLog, "GetSleepDisabled")

	if m.GetSleepDisabledFunc != nil {
		return m.GetSleepDisabledFunc()
	}

	// Default behavior: return current state with no error
	return m.CurrentState, nil
}

// SetSleepDisabled implements Controller.SetSleepDisabled
func (m *MockController) SetSleepDisabled(disabled bool) error {
	m.CallLog = append(m.CallLog, "SetSleepDisabled")

	if m.SetSleepDisabledFunc != nil {
		return m.SetSleepDisabledFunc(disabled)
	}

	// Default behavior: update internal state and return no error
	m.CurrentState = disabled
	return nil
}

// Reset clears the call log and resets state
func (m *MockController) Reset() {
	m.CallLog = make([]string, 0)
	m.CurrentState = false
	m.GetSleepDisabledFunc = nil
	m.SetSleepDisabledFunc = nil
}
