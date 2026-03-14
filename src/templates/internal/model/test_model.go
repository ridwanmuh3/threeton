package model

type SayHelloRequest struct {
	Name string `json:"name" validate:"required,alphaspace"`
}
