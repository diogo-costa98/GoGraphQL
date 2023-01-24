// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewOption struct {
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

type NewQuestion struct {
	Body    string       `json:"body"`
	Options []*NewOption `json:"options"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Option struct {
	ID      string `json:"id"`
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

type Question struct {
	ID      string    `json:"id"`
	Body    string    `json:"body"`
	Options []*Option `json:"options"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
