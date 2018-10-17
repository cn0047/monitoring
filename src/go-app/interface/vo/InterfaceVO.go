package vo

// InterfaceVO interface for ValueObject.
type InterfaceVO interface {
	// IsValid returns true in case when ValueObject is valid.
	IsValid() bool
}
