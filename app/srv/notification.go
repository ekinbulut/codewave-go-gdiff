package srv

type Notification struct {
	Message    string
	SmtpServer *SmtpServer
}

func NewNotification(message string) *Notification {
	return &Notification{
		Message:    message,
		SmtpServer: NewSmtpServer("smtp.gmail.com", "587", "", ""),
	}
}

// send email
func (n *Notification) SendEmail() {
	n.SmtpServer.SendEmail("ekinbulut@gmail.com", "ekinbulut@gmail.com", "Notification Email From Listener", n.Message)
}
