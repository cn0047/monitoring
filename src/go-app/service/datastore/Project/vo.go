package Project

type CreateVO struct {
	ID       string
	URL      string
	Method   string
	JSON     string
	Schedule int
}

func (_this CreateVO) GetName() string {
	return "Project.CreateVO"
}

func (_this CreateVO) IsValid() bool {
	return true
}
