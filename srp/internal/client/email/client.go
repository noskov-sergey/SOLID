package email

import (
	"net/smtp"
)

type Client struct {
	smtp smtp.Client
}

func New(smtp smtp.Client) *Client {
	return &Client{
		smtp: smtp,
	}
}
