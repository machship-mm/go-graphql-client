package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machship-mm/go-graphql-client"
)

func main() {
	client := graphql.NewClient("http://10.200.100.69:8080/graphql", nil)
	var err error

	var query struct {
		QueryUser []struct {
			ID       graphql.GqlID
			Username graphql.GqlString
			Name     graphql.GqlString
			Note     graphql.GqlString
		}
	}

	err = client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}
	fmt.Printf("%+v\n", query.QueryUser)
	fmt.Println()

	input := []AddUserInput{
		{
			UserShared: UserShared{
				Username: graphql.NewStringStruct("abc"),
				Name:     graphql.NewStringStruct("ABC"),
				Note:     graphql.NewStringStruct("note ABC"),
			},
			Password: graphql.NewStringStruct("password1"),
		},
		{
			UserShared: UserShared{
				Username: graphql.NewStringStruct("def"),
				Name:     graphql.NewStringStruct("DEF"),
			},
			Password: graphql.NewStringStruct("password2"),
		},
	}

	var mutation struct {
		AddUser struct {
			User []User
		} `graphql:"addUser(input: $input)"`
	}

	vars := map[string]interface{}{
		"input": input,
	}

	err = client.Mutate(context.Background(), &mutation, vars)
	if err != nil {
		// Handle error.
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", mutation.AddUser)
}

type AddUserInput struct {
	UserShared
	Password graphql.GqlString `json:"password,omitempty"`
}

type User struct {
	ID graphql.GqlID
	UserShared
}

type UserShared struct {
	Username graphql.GqlString `json:"username,omitempty"`
	Name     graphql.GqlString `json:"name,omitempty"`
	Note     graphql.GqlString `json:"note,omitempty"`
}
