package internal

import "os"

type Notification struct {
	SmtpServer *SmtpServer
}

func NewNotification() *Notification {

	u := os.Getenv("USERNAME")
	p := os.Getenv("PASSWORD")
	return &Notification{
		SmtpServer: NewSmtpServer("smtp.gmail.com", "587", u, p),
	}
}

// send email
func (n *Notification) SendEmail(from string, to string, subject string, body string) error {
	return n.SmtpServer.SendEmail(from, to, subject, body)
}
