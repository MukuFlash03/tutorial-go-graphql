package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
		"strconv"

	"github.com/MukuFlash03/hackernews/graph/model"
	"github.com/MukuFlash03/hackernews/internal/auth"
	"github.com/MukuFlash03/hackernews/internal/links"
	"github.com/MukuFlash03/hackernews/internal/users"
	"github.com/MukuFlash03/hackernews/pkg/jwt"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	// panic(fmt.Errorf("not implemented: CreateLink - createLink"))
	/*
		mutation {
			createLink(input: {title: "new link", address:"http://address.org"}){
				title,
				user{
					name
				}
				address
			}
		}


		var link model.Link
		var user model.User
		link.Address = input.Address
		link.Title = input.Title
		user.Name = "test123"
		link.User = &user
		return &link, nil

		mutation create{
			createLink(input: {title: "something", address: "somewhere"}){
				title,
				address,
				id,
			}
		}
	*/

	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	link.User = user

	linkID := link.Save()

	graphqlUser := &model.User{
		ID: user.ID,
		Name: user.Username,
	}

	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title:link.Title, Address:link.Address, User:graphqlUser}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	// panic(fmt.Errorf("not implemented: CreateUser - createUser"))

	/*
		mutation {
			createUser(input: {username: "user1", password: "123"})
		}
	*/

	var user users.User
	user.Username = input.Username
	user.Password = input.Password

	user.Create()

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	// panic(fmt.Errorf("not implemented: Login - login"))

	/*

	mutation {
		login(input: {username: "user1", password: "123"})
	}

	*/

	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	
	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	// panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))

	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	// panic(fmt.Errorf("not implemented: Links - links"))
	/*
		query {
			links{
				title
				address,
				user{
					name
				}
			}
		}

		var links []*model.Link
		dummyLink := model.Link {
			Title: "Hello world",
			Address: "https://www.google.com",
			User: &model.User {
				Name: "MukuFlash03",
			},
		}
		links = append(links, &dummyLink)
		return links, nil

		query {
			links {
				id
				title
				address
			}
		}
	*/
	
	var resultLinks []*model.Link
	var dbLinks []links.Link

	dbLinks = links.GetAll()

	for _, link := range dbLinks {
		graphqlUser := &model.User{
			ID: link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: graphqlUser})
	}

	return resultLinks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
