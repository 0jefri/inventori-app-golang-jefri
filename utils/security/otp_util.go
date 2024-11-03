package security

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

const (
	senderEmail    = "your_email@gmail.com"
	senderPassword = "your_password"
)

var ctx = context.Background()

func SendOTPByEmail(recipient, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Kode OTP untuk registrasi")

	// Isi email dengan OTP
	m.SetBody("text/plain", "Kode OTP Anda: "+otp)

	// Konfigurasi pengiriman email
	d := gomail.NewDialer("smtp.gmail.com", 5432, "postgres", "")

	// Kirim gmail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func GenerateOTP() (string, error) {
	// Generate random 6-digit OTP
	randomBytes := make([]byte, 3)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	otp := hex.EncodeToString(randomBytes)

	// Simpan OTP di Redis dengan sebuah batas waktu pendek (e.g., 5 minutes)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8080",
		Password: "",
		DB:       0,
	})

	err = client.Set(ctx, otp, nil, 5*time.Minute).Err()

	if err != nil {
		return "", nil
	}

	return otp, nil
}

func VerifyOTP(c *gin.Context) {
	userOTP := c.PostForm("otp") // Mengambil otp yang dimasukkan pengguna
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8080",
		Password: "",
		DB:       0,
	})

	storedOTP, err := client.Get(ctx, "user_otp").Result()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if userOTP != storedOTP {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
