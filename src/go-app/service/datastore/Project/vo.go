package Project

// EntityVO represents ValueObject for Project entity in DataStore.
type EntityVO struct {
	Name     string
	URL      string
	Method   string
	JSON     string
	Schedule int // seconds
}
