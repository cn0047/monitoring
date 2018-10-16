package Measurement

// CreateVO represents ValueObject to create new measurement entity in DataStore.
type CreateVO struct {
	Project      string
	Took         int
	ResponseCode int
}

// GetName {@inheritdoc}
func (_this CreateVO) GetName() string {
	return "Measurement.CreateVO"
}

// IsValid {@inheritdoc}
func (_this CreateVO) IsValid() bool {
	return true
}
