package vo

// InterfaceVO interface for ValueObject.
type InterfaceVO interface {
	// GetName gets system name for ValueObject.
	GetName() string

	// IsValid returns true in case when ValueObject is valid.
	IsValid() bool
}
