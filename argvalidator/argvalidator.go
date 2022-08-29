// Package for checking argument values
package argvalidator

import (
	"fmt"
)

func Validate(args []string) error {
	if len(args) == 1 {
		return fmt.Errorf("No argument provided")
	} else if len(args) > 2 {
		return fmt.Errorf("Only one argument accepted")
	}
	return nil
}
