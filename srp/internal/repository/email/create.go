package email

import (
	"fmt"

	"github.com/noskov-sergey/solid/srp/internal/model"
)

func (r *Repository) Create(e model.Email) (int64, error) {
	ct, err := r.db.Exec(`
		insert into email(from, to, subject, message)
		values ($1, $2, $3, $4)`,
		e.From, e.To,
		e.Subject, e.Message,
	)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return 0, fmt.Errorf("no rows are affected: %w", err)
	}

	return ct.LastInsertId(), nil
}
