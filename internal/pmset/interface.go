package pmset

// Controller defines the interface for interacting with macOS pmset commands.
// This interface allows for easy testing by providing a mock implementation.
type Controller interface {
	// GetSleepDisabled queries the current sleep disabled state
	GetSleepDisabled() (bool, error)

	// SetSleepDisabled sets the sleep disabled state
	SetSleepDisabled(disabled bool) error
}

// RealController implements the Controller interface using actual pmset commands
type RealController struct{}

// NewController returns a new RealController
func NewController() Controller {
	return &RealController{}
}

// GetSleepDisabled implements Controller.GetSleepDisabled
func (r *RealController) GetSleepDisabled() (bool, error) {
	return GetSleepDisabled()
}

// SetSleepDisabled implements Controller.SetSleepDisabled
func (r *RealController) SetSleepDisabled(disabled bool) error {
	return SetSleepDisabled(disabled)
}

// DefaultController is the package-level default controller
var DefaultController Controller = NewController()
