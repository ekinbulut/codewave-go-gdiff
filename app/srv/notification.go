package srv

type Notification struct {
	SmtpServer *SmtpServer
}

func NewNotification() *Notification {
	return &Notification{
		SmtpServer: NewSmtpServer("smtp.gmail.com", "587", "", ""),
	}
}

// send email
func (n *Notification) SendEmail(from string, to string, subject string, body string) error {
	return n.SmtpServer.SendEmail(from, to, subject, body)
}
