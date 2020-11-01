package errors

// ValidationError represents validation error returns from the service layer
type ValidationError struct {
	Err error
}

// Unwrap returns the underlying error
func (e ValidationError) Unwrap() error { return e.Err }
func (e ValidationError) Error() string { return e.Err.Error() }

// DBError represents database error returns from the service layer
type DBError struct {
	Err error
}

// Unwrap returns the underlying error
func (e DBError) Unwrap() error { return e.Err }
func (e DBError) Error() string { return e.Err.Error() }
