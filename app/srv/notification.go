package srv

import "fmt"

type Notification struct {
	Message string
	SmtpServer *SmtpServer
}

func NewNotification(message string, smtpServer *SmtpServer) *Notification {
	return &Notification{
		Message: message,
		SmtpServer: smtpServer,
	}
}

// send email
func (n *Notification) SendEmail() {
	n.SmtpServer.SendEmail("", "", "", n.Message)
}
