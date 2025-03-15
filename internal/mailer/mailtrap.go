package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"gopkg.in/gomail.v2"
)

type mailTrapClient struct {
	host      string
	port      int
	userName  string
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(host, userName, apikey, fromEmail string, port int) (mailTrapClient, error) {
	if apikey == "" {
		return mailTrapClient{}, errors.New("api key is required")
	}

	return mailTrapClient{
		host:      host,
		userName:  userName,
		port:      port,
		fromEmail: fromEmail,
		apiKey:    apikey,
	}, nil
}

func (m *mailTrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {

	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)

	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	messgae := gomail.NewMessage()

	fmt.Println("From email", m.fromEmail)
	fmt.Println("To email", email)

	messgae.SetHeader("From", m.fromEmail)
	messgae.SetHeader("To", email)
	messgae.SetHeader("Subject", subject.String())

	messgae.AddAlternative("text/html", body.String())

	dialer := gomail.NewDialer(m.host, m.port, m.userName, m.apiKey)

	if err := dialer.DialAndSend(messgae); err != nil {
		fmt.Println("Error sending email", err)
		return -1, err
	}

	return 200, nil
}
