// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Prefecture   int    `json:"prefecture"`
	Introduction string `json:"introduction"`
	BloodType    int    `json:"bloodType"`
}

type UpdateUserPasswordInput struct {
	ID                   string `json:"id"`
	ExPassword           string `json:"exPassword"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdateUserProfileInput struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Prefecture   int    `json:"prefecture"`
	Introduction string `json:"introduction"`
	BloodType    int    `json:"bloodType"`
}

type UpdateUserStatusInput struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Status       string `json:"status"`
	Prefecture   string `json:"prefecture"`
	Introduction string `json:"introduction"`
	BloodType    string `json:"bloodType"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
