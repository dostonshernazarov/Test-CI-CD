package models

import (
	// "github/Test_CI-CD/api/handlers/models"
	pbu "github/test-ci-cd-github/genproto/user"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// UserResponse ...
type UserResponse struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"udpated_at"`
}

// UserDetail ...
type UserDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// ResponseMessage ...
type ResponseMessage struct {
	Content string `json:"content"`
}

func (rm *UserDetail) Validate() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
		validation.Field(
			&rm.Password,
			validation.Required,
			validation.Length(8, 30),
			validation.Match(regexp.MustCompile("[a-z]|[A-Z][1-9]")),
		),
	)
}

func ParseStruct(req *pbu.User, access string) *User {
	return &User{
		Id:          req.Id,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
		Bio:         req.Bio,
		Age:         int(req.Age),
		Gender:      req.Gender,
		Email:       req.Email,
		Password:    req.Password,
		AccessToken: access,
	}
}
