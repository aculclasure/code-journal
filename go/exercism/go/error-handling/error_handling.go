package erratum

// Use opens a resource and calls Frob(input) on that resource,
// returning any errors encountered that are not a TransientError.
func Use(o ResourceOpener, input string) (e error) {
	resource, e := o()
	if e != nil {
		if _, ok := e.(TransientError); ok {
			return Use(o, input)
		}
		return e
	}

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				resource.Defrob(r.(FrobError).defrobTag)
				e = r.(FrobError).inner
			case error:
				e = r.(error)
			default:
				panic(r)
			}
		}
		resource.Close()
	}()
	resource.Frob(input)
	return nil
}
