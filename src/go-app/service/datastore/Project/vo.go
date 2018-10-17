package Project

// CreateVO represents ValueObject to create new project entity in DataStore.
type CreateVO struct {
	ID       string
	URL      string
	Method   string
	JSON     string
	Schedule int
}

// IsValid {@inheritdoc}
func (vo CreateVO) IsValid() bool {
	return true
}
