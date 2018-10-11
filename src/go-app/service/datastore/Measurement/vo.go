package Measurement

type CreateVO struct {
	Project      string
	Took         int
	ResponseCode int
}

func (_this CreateVO) GetName() string {
	return "Measurement.CreateVO"
}

func (_this CreateVO) IsValid() bool {
	return true
}
