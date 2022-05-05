package srv

import (
	"os"
	"reflect"
	"testing"
)

// read environmnet variables
// go test -v -run TestNewSmtpServer -count=1
// go test -v -run TestSmtpServer_SendEmail -count=1

func TestNewSmtpServer(t *testing.T) {

	// read environmnet variables
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASS")

	print("email: ", email)

	type args struct {
		host string
		port string
		user string
		pass string
	}
	tests := []struct {
		name string
		args args
		want *SmtpServer
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{host: "smtp.gmail.com", port: "587", user: email, pass: password},
			want: &SmtpServer{
				Host: "smtp.gmail.com",
				Port: "587",
				User: email,    // environment variable
				Pass: password, // environment variable
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSmtpServer(tt.args.host, tt.args.port, tt.args.user, tt.args.pass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSmtpServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmtpServer_SendEmail(t *testing.T) {
	type args struct {
		from    string
		to      string
		subject string
		body    string
	}
	tests := []struct {
		name    string
		s       *SmtpServer
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "testing mail",
			s: &SmtpServer{
				Host: "smtp.gmail.com",
				Port: "587",
				User: "",
				Pass: "",
			},
			args: args{
				from:    "",
				to:      "",
				subject: "test mail",
				body:    "Test message",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SendEmail(tt.args.from, tt.args.to, tt.args.subject, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("SmtpServer.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
