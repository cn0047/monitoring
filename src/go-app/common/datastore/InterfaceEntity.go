package datastore

// InterfaceEntity interface for datastore Entity.
type InterfaceEntity interface {
	// GetKind method which returns string name (kind) for datastore Entity.
	GetKind() string
}
