package flatten

// Flatten accepts a list of possible nested lists and returns the
// flattened list of items, ignoring any nil values.
func Flatten(input interface{}) []interface{} {
	flattened := []interface{}{}

	if input != nil {
		if list, ok := input.([]interface{}); ok {
			for _, v := range list {
				flattened = append(flattened, Flatten(v)...)
			}
		} else {
			flattened = append(flattened, input)
		}
	}
	return flattened
}
