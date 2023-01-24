package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/diogo-costa98/GoGraphQL/auth"
	"github.com/diogo-costa98/GoGraphQL/graph/model"
	"github.com/diogo-costa98/GoGraphQL/internal/options"
	"github.com/diogo-costa98/GoGraphQL/internal/questions"
	"github.com/diogo-costa98/GoGraphQL/internal/users"
	"github.com/diogo-costa98/GoGraphQL/jwt"
)

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	//Auth user
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	//Fill Question
	question := questions.Question{
		User: user,
		Body: strings.Replace(input.Body, "'", "''", -1),
	}
	question.ID = strconv.FormatInt(question.Save(), 10)

	//Create Options
	var createdOptions []*model.Option
	for _, input_option := range input.Options {
		//Create Single Option
		option := options.Option{
			Question: &question,
			Body:     strings.Replace(input_option.Body, "'", "''", -1),
			Correct:  input_option.Correct,
		}
		optionID := option.Save()

		//Add option to options list
		createdOptions = append(createdOptions, &model.Option{ID: strconv.FormatInt(optionID, 10), Body: strings.Replace(option.Body, "''", "'", -1), Correct: option.Correct})
	}
	return &model.Question{ID: question.ID, Body: strings.Replace(question.Body, "''", "'", -1), Options: createdOptions}, nil
}

// UpdateQuestion is the resolver for the updateQuestion field.
func (r *mutationResolver) UpdateQuestion(ctx context.Context, id string, input model.NewQuestion) (*model.Question, error) {
	//Auth user
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	//Get question and options to update
	question, optionsList := questions.GetQuestionById(id)

	//Check user access to question
	if question.User.ID != user.ID {
		return nil, fmt.Errorf("access to resource unauthorized")
	}

	//Delete all options tied to the question
	options.Delete(optionsList)

	//Update question and create new options
	question.Body = strings.Replace(input.Body, "'", "''", -1)
	question.ID = strconv.FormatInt(question.Update(), 10)

	//Create Options
	var createdOptions []*model.Option
	for _, input_option := range input.Options {
		//Create Single Option
		option := options.Option{
			Question: &question,
			Body:     strings.Replace(input_option.Body, "'", "''", -1),
			Correct:  input_option.Correct,
		}
		optionID := option.Save()

		//Add option to options list
		createdOptions = append(createdOptions, &model.Option{ID: strconv.FormatInt(optionID, 10), Body: strings.Replace(option.Body, "''", "'", -1), Correct: option.Correct})
	}
	return &model.Question{ID: question.ID, Body: strings.Replace(question.Body, "''", "'", -1), Options: createdOptions}, nil
}

// DeleteQuestion is the resolver for the deleteQuestion field.
func (r *mutationResolver) DeleteQuestion(ctx context.Context, id string) (*model.Question, error) {
	var question questions.Question
	var optionsList []questions.Option

	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	//Get Question using its ID
	question, optionsList = questions.GetQuestionById(id)

	if question.User.ID != user.ID {
		return nil, fmt.Errorf("access to resource unauthorized")
	}

	//Delete all options tied to the question
	options.Delete(optionsList)

	//Delete question
	question.Delete()

	return &model.Question{ID: question.ID, Body: strings.Replace(question.Body, "''", "'", -1)}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	user.Create()

	//Generates user token
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}

	//Checks users' credentials
	if !user.Authenticate() {
		return "", &users.WrongUsernameOrPasswordError{}
	}

	//Generates user token
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Questions(ctx context.Context, page *string, pageSize *string) ([]*model.Question, error) {
	//Auth user
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	//Retrieve information from db
	dbQuestions, dbOptions := questions.GetAll(page, pageSize, &user.ID)

	//Creates questions list with options list inside
	var resultQuestions []*model.Question
	for _, question := range dbQuestions {
		var resultOptions []*model.Option

		for _, option := range dbOptions {
			if option.Question.ID == question.ID {
				resultOptions = append(resultOptions, &model.Option{ID: option.ID, Body: strings.Replace(option.Body, "''", "'", -1), Correct: option.Correct})
			}
		}
		resultQuestions = append(resultQuestions, &model.Question{ID: question.ID, Body: strings.Replace(question.Body, "''", "'", -1), Options: resultOptions})
	}
	return resultQuestions, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
