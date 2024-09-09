// models.go
package models

// User represents the user type in the GraphQL schema.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
