// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Account struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Mutation struct {
}

type Query struct {
}

type Todo struct {
	ID   string `json:"id"`
	Done bool   `json:"done"`
	Text string `json:"text"`
}
