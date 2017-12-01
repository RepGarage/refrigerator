package db

// UnknownModelError type
type UnknownModelError struct{}

func (e UnknownModelError) Error() string {
	return "unknown model type passed"
}
