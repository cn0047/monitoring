package Measurement

// CreateVO represents ValueObject to create new measurement entity in DataStore.
type CreateVO struct {
	Project      string
	Took         int
	ResponseCode int
}

// IsValid {@inheritdoc}
func (vo CreateVO) IsValid() bool {
	return true
}
