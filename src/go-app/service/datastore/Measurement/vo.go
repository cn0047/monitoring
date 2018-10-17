package Measurement

// CreateVO represents ValueObject to create new measurement entity in DataStore.
type CreateVO struct {
	Project      string
	Took         int
	ResponseCode int
}

// GetName {@inheritdoc}
func (vo CreateVO) GetName() string {
	return "Measurement.CreateVO"
}

// IsValid {@inheritdoc}
func (vo CreateVO) IsValid() bool {
	return true
}
