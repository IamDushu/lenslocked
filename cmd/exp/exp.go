package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/iamDushu/lenslocked/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err = es.ForgotPassword("hey.dushyanth@gmail.com", "https://lenslocked.com/reset-pw?token=abc123")
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
