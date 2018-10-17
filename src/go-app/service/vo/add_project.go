package vo

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type AddProjectVO struct {
	Name     string `json:"name" validate:"required"`
	URL      string `json:"url" validate:"required,url"`
	Method   string `json:"method" validate:"required,alpha"`
	JSON     string `json:"json" validate:""`
	Schedule string `json:"schedule" validate:"required,numeric"`
}

// IsValid {@inheritdoc}
func (v AddProjectVO) IsValid() bool {
	validate := validator.New()
	err := validate.Struct(v)
	log.Printf("☢️ %+v", err)
	return false
}
