package email

import (
	"fmt"
	"net/smtp"

	"github.com/noskov-sergey/solid/srp/internal/model"
)

const (
	smtpHost = "smtp.gmail.com"
)

func (c *Client) Send(email model.Email) error {
	auth := smtp.PlainAuth("", email.From, password, smtpHost)

	err := c.smtp.Auth(auth)

	err = c.smtp.Mail(email.From)
	if err != nil {
		return fmt.Errorf("sendemail: %w", err)
	}

	return nil
}
