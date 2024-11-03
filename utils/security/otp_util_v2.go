package security

import (
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func GenerateOTPV2(count int) string {
	rand.Seed(time.Now().UnixNano())

	otp := ""

	for i := 0; i < count; i++ {
		otp += strconv.Itoa(rand.Intn(10))
	}

	return otp
}

func SendEmail(recipient, otp string) error {
	username := "littleeyes17@gmail.com"
	password := "@Ency1705clopaedia"

	// Server SMTP Gmail
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	// Pesan Email
	to := recipient

	subject := "Your OTP"
	body := "Your OTP Is: " + otp

	// Konfigurasi SMTP
	auth := smtp.PlainAuth("", username, password, smtpServer)

	// Format pesan email
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Kirim Email melalui server smtp
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, username, []string{to}, message)

	if err != nil {
		return err
	}

	return nil
}
