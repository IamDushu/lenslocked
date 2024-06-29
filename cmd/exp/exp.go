package main

import (
	"os"

	"github.com/go-mail/mail/v2"
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

	msg.WriteTo(os.Stdout)
}

// MIME-Version: 1.0
// Date: Sat, 29 Jun 2024 15:08:41 +0530
// To: hey.dushyanth@gmail.com
// From: test@lenslocked.com
// Subject: This is a test email
// Content-Type: multipart/alternative;
//  boundary=c1ef3f377cb6ed61bd8e00fe2b8ad70ace16cf3a5881e6a6d40c55e915bb

// --c1ef3f377cb6ed61bd8e00fe2b8ad70ace16cf3a5881e6a6d40c55e915bb
// Content-Transfer-Encoding: quoted-printable
// Content-Type: text/plain; charset=UTF-8

// Hello, This is the body of the email
// --c1ef3f377cb6ed61bd8e00fe2b8ad70ace16cf3a5881e6a6d40c55e915bb
// Content-Transfer-Encoding: quoted-printable
// Content-Type: text/html; charset=UTF-8

// <h1>Hello there buddy!</h1> <p>This is the email; Hope you enjoy it</p>
// --c1ef3f377cb6ed61bd8e00fe2b8ad70ace16cf3a5881e6a6d40c55e915bb--