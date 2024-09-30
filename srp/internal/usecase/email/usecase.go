package email

import "github.com/noskov-sergey/solid/srp/internal/model"

type Repository interface {
	Create(model.Email) (int64, error)
}

type Client interface {
	Send(email model.Email) error
}

type useCase struct {
	rep    Repository
	client Client
}

func New(rep Repository, client Client) *useCase {
	return &useCase{
		rep:    rep,
		client: client,
	}
}
