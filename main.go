package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	gomail "gopkg.in/mail.v2"
)

func main() {
	email := os.Getenv("GMAIL")
	pass := os.Getenv("PASS_GMAIL")
	if email == "" || pass == "" {
		log.Fatalln("configure sus variables de entorno")
	} else {
		fmt.Println(email)
		fmt.Println(pass)
	}
	m := gomail.NewMessage()

	m.SetHeader("From", "luispfcanales@gmail.com")
	m.SetHeader("To", "lpfunoc@unamad.edu.pe")

	m.SetHeader("Subject", "Gophers GO!")

	//m.Embed("gopherConclusion-min.png")
	//m.SetBody("text/plain", "https://example.com")

	t := template.Must(template.ParseFiles("WellcomeTemplate.html"))
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, "Registrate")
	})

	d := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
