package pmset

import "errors"

// Sentinel errors for pmset operations
var (
	// ErrPermissionDenied indicates the user lacks necessary privileges to modify pmset settings
	ErrPermissionDenied = errors.New("permission denied: administrator privileges required")

	// ErrCommandNotFound indicates the pmset command is not available on the system
	ErrCommandNotFound = errors.New("pmset command not found")

	// ErrParseFailure indicates failure to parse pmset command output
	ErrParseFailure = errors.New("failed to parse pmset output")

	// ErrUserCancelled indicates the user cancelled the password dialog
	ErrUserCancelled = errors.New("user cancelled password dialog")
)
