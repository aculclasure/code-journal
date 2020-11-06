package erratum

import (
	"fmt"
)

// Use opens a resource and calls Frob(input) on that resource,
// returning any errors encountered that are not a TransientError.
func Use(o ResourceOpener, input string) (e error) {
	var resource Resource

	for resource, e = o(); e != nil; {
		switch e.(type) {
		case TransientError:
			resource, e = o()
		default:
			return e
		}
	}
	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				resource.Defrob(r.(FrobError).defrobTag)
				e = r.(FrobError).inner
			case error:
				e = r.(error)
			default:
				e = fmt.Errorf("got unexpected return value from call to Frob(input): %v", r)
			}
		}
	}()
	resource.Frob(input)
	return nil
}
