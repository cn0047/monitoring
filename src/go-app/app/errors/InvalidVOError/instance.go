package InvalidVOError

import (
	"fmt"
)

// Instance represents container for ValueObject errors.
type Instance struct {
	errors map[string]string
}

// New gets new InvalidVOError instance.
func New() *Instance {
	e := Instance{}
	e.errors = make(map[string]string)

	return &e
}

// SetError sets VO error.
func (i *Instance) SetError(key string, err string) {
	i.errors[key] = err
}

// IsEmpty returns true in case when errors map is empty.
func (i Instance) IsEmpty() bool {
	return len(i.errors) == 0
}

// Error gets InvalidVOError error message string.
func (i *Instance) Error() string {
	msg := "Got invalid data. "
	for field, err := range i.errors {
		msg += fmt.Sprintf("%s: %s;", field, err)
	}

	return msg
}

// GetErrors gets map with InvalidVOError's errors.
func (i Instance) GetErrors() map[string]string {
	return i.errors
}
