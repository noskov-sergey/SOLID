package email

import (
	"fmt"

	"github.com/noskov-sergey/solid/srp/internal/model"
)

func (u *useCase) Send(e model.Email) (int64, error) {
	id, err := u.rep.Create(e)
	if err != nil {
		return 0, fmt.Errorf("create: %w", err)
	}

	err = u.client.Send(e)

	return id, nil
}
