// Package twofer contains a function that generates
// a dynamic message.
package twofer

import (
	"fmt"
)

// ShareWith returns a dynamically generated message
// depending on the value of the name parameter.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
