package test

import (
	"os"
	"reflect"
	"testing"

	service "github.com/ekinbulut/go-http-crawler/internal"
)

// read environmnet variables
// go test -v -run TestNewSmtpServer -count=1
// go test -v -run TestSmtpServer_SendEmail -count=1

func TestNewSmtpServer(t *testing.T) {

	// read environmnet variables
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASS")

	type args struct {
		host string
		port string
		user string
		pass string
	}
	tests := []struct {
		name string
		args args
		want *service.SmtpServer
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{host: "smtp.gmail.com", port: "587", user: email, pass: password},
			want: &service.SmtpServer{
				Host: "smtp.gmail.com",
				Port: "587",
				User: email,    // environment variable
				Pass: password, // environment variable
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewSmtpServer(tt.args.host, tt.args.port, tt.args.user, tt.args.pass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSmtpServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmtpServer_SendEmail(t *testing.T) {

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASS")

	type args struct {
		from    string
		to      string
		subject string
		body    string
	}
	tests := []struct {
		name    string
		s       *service.SmtpServer
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "testing mail",
			s: &service.SmtpServer{
				Host: "smtp.gmail.com",
				Port: "587",
				User: email,
				Pass: password,
			},
			args: args{
				from:    email,
				to:      email,
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
