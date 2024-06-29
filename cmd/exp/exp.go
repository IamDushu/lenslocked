package main

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	host = "sandbox.smtp.mailtrap.io"
	port = 2525
	username = "d9b65401f343a1"
	password = "78f05ab331b756"
)

func main() {
	from := "test@lenslocked.com"
	to := "hey.dushyanth@gmail.com"
	subject := "This is a test email"
	plainText := "Hello, This is the body of the email"
	html := `<h1>Hello there buddy!</h1> <p>This is the email; Hope you enjoy it</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plainText)
	msg.AddAlternative("text/html", html)
	// msg.WriteTo(os.Stdout)

    dailer := mail.NewDialer(host, port, username, password)

	err := dailer.DialAndSend(msg)
	if err != nil {
			panic(err)
	}
	fmt.Println("message sent")

	//If we want to send multiple emails we can use dialer.Dial
	// sender, err := dailer.Dial()
	// if err != nil {
	// 	panic(err)
	// }
	// defer sender.Close()
	// sender.Send(from, []string{to}, msg)
	// sender.Send(from, []string{to}, msg)
}