package db

// UnknownModelError type
type UnknownModelError struct{}

func (e UnknownModelError) Error() string {
	return "unknown model type passed"
}

// UnknownDatabaseNameTypeError type
type UnknownDatabaseNameTypeError struct{}

func (e UnknownDatabaseNameTypeError) Error() string {
	return "unknown database name type passed"
}
