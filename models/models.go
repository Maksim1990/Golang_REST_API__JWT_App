package models


type JwtToken struct {
	Token string `json:"token"`
}
type UserBase struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type User struct {
	UserBase
	Password string `json:"-"`
}
type UserSecure struct {
	UserBase
	Password string `json:"password"`
}

type Response struct {
	StatusCode  int       `json:"statusCode"`
	Headers     map[string]string  `json:"headers"`
	Message        string    `json:"message"`
	Error        string    `json:"error"`
}

type Exception struct {
	Message string `json:"message"`
}
