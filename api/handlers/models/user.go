package models

// User ...
type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Bio         string `json:"bio"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"acces_token"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Pong struct {
	Pong string `json:"pong"`
}