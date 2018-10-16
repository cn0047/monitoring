package Project

// CreateVO represents ValueObject to create new project entity in DataStore.
type CreateVO struct {
	ID       string
	URL      string
	Method   string
	JSON     string
	Schedule int
}

// GetName {@inheritdoc}
func (_this CreateVO) GetName() string {
	return "Project.CreateVO"
}

// IsValid {@inheritdoc}
func (_this CreateVO) IsValid() bool {
	return true
}
