package flatten

// Flatten accepts a list of possible nested lists and returns the
// flattened list of items, ignoring any nil values.
func Flatten(input interface{}) []interface{} {
	if list, ok := input.([]interface{}); ok {
		return flattenList(list)
	}
	return nil
}

func flattenList(input []interface{}) []interface{} {
	flattened := []interface{}{}

	for _, v := range input {
		if v != nil {
			if nestedList, ok := v.([]interface{}); ok {
				flattened = append(flattened, flattenList(nestedList)...)
			} else {
				flattened = append(flattened, v)
			}
		}
	}
	return flattened
}
