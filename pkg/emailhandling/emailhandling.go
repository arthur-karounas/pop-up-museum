package emailhandling

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
)

// EmailService is an email messaging service.
type EmailService struct {
	username string
	password string
	host     string
	port     int
	auth     smtp.Auth
}

// NewEmailService creates a new EmailService instance with the specified credentials.
func NewEmailService(username, password, host string, port int) *EmailService {
	auth := smtp.PlainAuth("", username, password, host)
	return &EmailService{
		username: username,
		password: password,
		host:     host,
		port:     port,
		auth:     auth,
	}
}

// Send sends an email with the specified recipient, subject, and text.
func (es *EmailService) Send(to, subject, body string) error {
	message := fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		body

	// Establish a connection.
	conn, err := smtp.Dial(fmt.Sprintf("%s:%d", es.host, es.port))
	if err != nil {
		return errors.New("error occurred while connection establishing")
	}
	defer conn.Close()

	// Enable TLS for communication.
	if err := conn.StartTLS(&tls.Config{ServerName: es.host}); err != nil {
		return errors.New("error occurred while enabling TLS")
	}

	// Authenticate.
	if err := conn.Auth(es.auth); err != nil {
		return errors.New("error occurred while authentication establishing")
	}

	// Sender designation.
	if err := conn.Mail(es.username); err != nil {
		return errors.New("error occurred while sender designation")
	}

	// Receiver designation.
	if err := conn.Rcpt(to); err != nil {
		return errors.New("error occurred while receiver designation")
	}

	// Data transmission.
	wc, err := conn.Data()
	if err != nil {
		return errors.New("error occurred while email sending")
	}
	defer wc.Close()

	// Write email content.
	_, err = wc.Write([]byte(message))
	if err != nil {
		return errors.New("error occurred while data recording")
	}

	return nil
}
