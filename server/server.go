// server.go
package server

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"graphql-demo/models"
)

// ... setup GraphQL schema ...

func resolveUser(p graphql.ResolveParams) (interface{}, error) {
	// Simulate fetching a user from the database
	user := models.User{
		ID:    "1",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	return user, nil
}

func Start() error {
	fields := graphql.Fields{
		"user": &graphql.Field{
			Type:    graphql.NewObject(graphql.ObjectConfig{Name: "User", Fields: DefaultUserFields()}),
			Resolve: resolveUser,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String},
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return err
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	return http.ListenAndServe(":8080", nil)
}

// DefaultUserFields returns the default fields for the User type.
func DefaultUserFields() graphql.Fields {
	return graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	}
}
