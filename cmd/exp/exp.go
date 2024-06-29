package main

import (
	"fmt"

	"github.com/iamDushu/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "d9b65401f343a1"
	password = "78f05ab331b756"
)

func main() {
	// email := models.Email{
	// 	From:      "test@lenslocked.com",
	// 	To:        "hey.dushyanth@gmail.com",
	// 	Subject:   "This is a test email",
	// 	Plaintext: "Hello, This is the body of the email",
	// 	HTML:      `<h1>Hello there buddy!</h1> <p>This is the email; Hope you enjoy it</p>`,
	// }

	smtpConf := models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}

	es := models.NewEmailService(smtpConf)
	// err := es.Send(email)
	err := es.ForgotPassword("hey.dushyanth@gmail.com", "https://lenslocked.com/reset-pw?token=abc123")

	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")

	//If we want to send multiple emails we can use dialer.Dial
	// sender, err := dailer.Dial()
	// if err != nil {
	// 	panic(err)
	// }
	// defer sender.Close()
	// sender.Send(from, []string{to}, msg)
	// sender.Send(from, []string{to}, msg)
}
