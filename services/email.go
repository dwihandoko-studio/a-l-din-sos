package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

func SendEmail(email, token, code, subject, content string) error {
	smtpUsername := "noreplay.dinsoslampungtengah@gmail.com"
	smtpPassword := "wzguklotayzumaeg"
	sender := NewGmailSender("Si-lastri Dinsos Kab. Lampung Tengah", smtpUsername, smtpPassword)

	to := []string{email}

	return sender.SendEmail(subject, content, to, nil, nil, nil)

}

// func SendEmailNew(email, token, code string) error {
// 	// Konfigurasi SMTP
// 	smtpHost := os.Getenv("SMTP_HOST")
// 	smtpPort := os.Getenv("SMTP_PORT")
// 	smtpUsername := os.Getenv("SMTP_USER")
// 	smtpPassword := os.Getenv("SMTP_PASSWORD")
// 	// smtpCrypto := ""

// 	// Membuat pesan email
// 	from := os.Getenv("SMTP_USER")
// 	to := email
// 	// subject := "Reset Password"
// 	// body := "Silakan klik tautan berikut untuk mereset kata sandi Anda: https://example.com/reset"

// 	// message := "From: " + from + "\n" +
// 	// 	"To: " + to + "\n" +
// 	// 	"Subject: " + subject + "\n\n" +
// 	// 	body

// 	// // Mengirim email melalui SMTP server
// 	// auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
// 	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// return nil

// 	data := struct {
// 		Email  string
// 		Urlweb string
// 		Token  string
// 		Code   string
// 	}{
// 		Email:  email,
// 		Urlweb: os.Getenv("WEB_URL"),
// 		Token:  token,
// 		Code:   code,
// 	}

// 	// Memuat template reset password dari file\
// 	// dir, errd := os.Getwd()
// 	// if errd != nil {
// 	// 	return fmt.Errorf("error directory: %w", errd)
// 	// }
// 	// fmt.Println("Working Directory:", dir)
// 	tmpl, err := template.ParseFiles("template/reset_password_email.html")
// 	if err != nil {
// 		return fmt.Errorf("error parsing template: %w", err)
// 	}

// 	// Membuat buffer untuk menampung hasil eksekusi template
// 	body := new(strings.Builder)

// 	// Mengeksekusi template ke dalam buffer
// 	err = tmpl.Execute(body, data)
// 	if err != nil {
// 		return fmt.Errorf("error executing template: %w", err)
// 	}

// 	// subject := []byte(body.String())
// 	// subject := body.String()

// 	// msg := "From: " + from + "\n" +
// 	// 	"To: " + to + "\n" +
// 	// 	"Subject: " + *subject + "\n" +
// 	// 	"\n" +
// 	// 	body

// 	// // Establish a secure connection using TLS
// 	// tlsConfig := &tls.Config{
// 	// 	InsecureSkipVerify: true, // Only use this for testing purposes
// 	// 	ServerName:         smtpHost,
// 	// }

// 	// // Connect to the SMTP server using TLS
// 	// client, err := smtp.Dial(fmt.Sprintf("%s:%s", smtpHost, smtpPort))
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to connect to the server: %w", err)
// 	// }
// 	// defer client.Close()

// 	// err = client.StartTLS(tlsConfig)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to start TLS: %w", err)
// 	// }

// 	// // Authenticate to the SMTP server
// 	// auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
// 	// if err := client.Auth(auth); err != nil {
// 	// 	return fmt.Errorf("failed to authenticate: %w", err)
// 	// }

// 	// // Set the sender and recipient
// 	// if err := client.Mail(from); err != nil {
// 	// 	return fmt.Errorf("failed to set sender: %w", err)
// 	// }
// 	// if err := client.Rcpt(to); err != nil {
// 	// 	return fmt.Errorf("failed to set recipient: %w", err)
// 	// }

// 	// // Send the email message
// 	// w, err := client.Data()
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to open data writer: %w", err)
// 	// }
// 	// _, err = w.Write([]byte(body.String()))
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to write email data: %w", err)
// 	// }
// 	// err = w.Close()
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to close data writer: %w", err)
// 	// }

// 	// return nil

// 	// // Mengirim email menggunakan protokol SMTP
// 	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
// 	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body.String()))
// 	if err != nil {
// 		return fmt.Errorf("error sending email: %w", err)
// 	}

// 	return nil
// }

func SendEmailOld(email, token, code string) error {
	// Konfigurasi SMTP
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	// smtpCrypto := ""

	// Membuat pesan email
	from := os.Getenv("SMTP_USER")
	to := email
	// subject := "Reset Password"
	// body := "Silakan klik tautan berikut untuk mereset kata sandi Anda: https://example.com/reset"

	// message := "From: " + from + "\n" +
	// 	"To: " + to + "\n" +
	// 	"Subject: " + subject + "\n\n" +
	// 	body

	// // Mengirim email melalui SMTP server
	// auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	// if err != nil {
	// 	return err
	// }

	// return nil

	data := struct {
		Email  string
		Urlweb string
		Token  string
		Code   string
	}{
		Email:  email,
		Urlweb: os.Getenv("WEB_URL"),
		Token:  token,
		Code:   code,
	}

	// Memuat template reset password dari file\
	// dir, errd := os.Getwd()
	// if errd != nil {
	// 	return fmt.Errorf("error directory: %w", errd)
	// }
	// fmt.Println("Working Directory:", dir)
	tmpl, err := template.ParseFiles("template/reset_password_email.html")
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Membuat buffer untuk menampung hasil eksekusi template
	body := new(strings.Builder)

	// Mengeksekusi template ke dalam buffer
	err = tmpl.Execute(body, data)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// subject := []byte(body.String())
	// subject := body.String()

	// msg := "From: " + from + "\n" +
	// 	"To: " + to + "\n" +
	// 	"Subject: " + *subject + "\n" +
	// 	"\n" +
	// 	body

	// Establish a secure connection using TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Only use this for testing purposes
		ServerName:         smtpHost,
	}

	// Connect to the SMTP server using TLS
	client, err := smtp.Dial(fmt.Sprintf("%s:%s", smtpHost, smtpPort))
	if err != nil {
		return fmt.Errorf("failed to connect to the server: %w", err)
	}
	defer client.Close()

	err = client.StartTLS(tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to start TLS: %w", err)
	}

	// Authenticate to the SMTP server
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	// Set the sender and recipient
	if err := client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err := client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	// Send the email message
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to open data writer: %w", err)
	}
	_, err = w.Write([]byte(body.String()))
	if err != nil {
		return fmt.Errorf("failed to write email data: %w", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	return nil

	// // Mengirim email menggunakan protokol SMTP
	// auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	// err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body.String()))
	// if err != nil {
	// 	return fmt.Errorf("error sending email: %w", err)
	// }

	// return nil
}
