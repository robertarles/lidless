// Package pmset provides a Go wrapper for macOS pmset commands
// to query and modify sleep state.
package pmset

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// disableSleepRegex matches the disablesleep/SleepDisabled setting in pmset output.
// macOS may display it as either "disablesleep" or "SleepDisabled"
var disableSleepRegex = regexp.MustCompile(`(?i)(disablesleep|sleepdisabled)\s+(\d+)`)

// GetSleepDisabled queries the current sleep disabled state via pmset -g.
// Returns true if sleep is disabled (disablesleep=1), false if enabled (disablesleep=0).
// If the disablesleep setting is not found in the output (older macOS), defaults to false.
func GetSleepDisabled() (bool, error) {
	cmd := exec.Command("pmset", "-g")
	output, err := cmd.Output()
	if err != nil {
		// Check if command not found
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// Check stderr for specific error types
			stderr := string(exitErr.Stderr)
			if strings.Contains(stderr, "not found") || strings.Contains(stderr, "No such file") {
				return false, fmt.Errorf("%w: %v", ErrCommandNotFound, err)
			}
		}
		return false, fmt.Errorf("pmset query failed: %w", err)
	}

	matches := disableSleepRegex.FindStringSubmatch(string(output))
	if len(matches) < 3 {
		// disablesleep not found in output, default to sleep enabled
		return false, nil
	}

	// matches[1] is the setting name, matches[2] is the value
	val, err := strconv.Atoi(matches[2])
	if err != nil {
		return false, fmt.Errorf("%w: invalid value %q", ErrParseFailure, matches[2])
	}
	return val == 1, nil
}

// SetSleepDisabled sets the sleep disabled state via pmset with privilege escalation.
// Sets disablesleep to 1 if disabled=true, 0 if disabled=false.
// Uses sudo to trigger macOS authentication (Touch ID if configured, or password).
// Returns any error from command execution, including ErrUserCancelled if auth is cancelled.
func SetSleepDisabled(disabled bool) error {
	val := "0"
	if disabled {
		val = "1"
	}

	// Use sudo to execute pmset with privileges
	// This will trigger Touch ID if configured via pam_tid.so, otherwise password prompt
	cmd := exec.Command("sudo", "pmset", "-a", "disablesleep", val)

	output, err := cmd.CombinedOutput()
	if err != nil {
		outputStr := string(output)

		// Check for user cancellation (cancelled Touch ID or password prompt)
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// Exit code 1 typically means auth was cancelled or failed
			if exitErr.ExitCode() == 1 {
				return fmt.Errorf("%w", ErrUserCancelled)
			}
		}

		// Check for command not found
		if strings.Contains(outputStr, "not found") ||
		   strings.Contains(outputStr, "No such file") {
			return fmt.Errorf("%w: %v", ErrCommandNotFound, err)
		}

		// Check for permission denied
		if strings.Contains(outputStr, "must be root") ||
		   strings.Contains(outputStr, "Operation not permitted") ||
		   strings.Contains(outputStr, "Permission denied") {
			return fmt.Errorf("%w: %v", ErrPermissionDenied, err)
		}

		// Generic error
		return fmt.Errorf("pmset command failed: %w, output: %s", err, outputStr)
	}

	return nil
}
