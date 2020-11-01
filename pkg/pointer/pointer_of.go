package pointer

// String return pointer of string
func String(str string) *string {
	return &str
}

// Of returns the pointer to the given value
func Of(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return &v
}
