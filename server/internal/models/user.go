package models

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
}

type UserResponse struct {
	ID        string
	Username  string
	CreatedAt string
	UpdatedAt string
}
