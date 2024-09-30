package model

type Email struct {
	From    string `db:"from"`
	To      string `db:"to"`
	Subject string `db:"subject"`
	Message string `db:"message"`
}
