package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"path/filepath"
)

const (
	MAX_BYTES = 52428800 // This number is equal 50mb
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env variables")
	}

	if len(os.Args[1:]) < 1 {
		log.Fatal("invalid file path do sent do Kindle, pls add a valid path")
	}

	validExtensions := []string{".epub", ".pdf", ".mobi", ".doc", ".txt"}

	pathOfArchiveToSend := ""
	for i, v := range os.Args[1:] {
		if len(os.Args[1:])-1 == i {
			pathOfArchiveToSend += v
			continue
		}
		pathOfArchiveToSend += v + " "
	}

	ext := filepath.Ext(pathOfArchiveToSend)
	validFlag := false
	for _, v := range validExtensions {
		if ext == v {
			validFlag = true
		}
	}
	if !validFlag {
		log.Fatal("invalid file format, pls add a valid archive with one os this extensions .epub, .pdf, .mobi or .doc", ext)
	}

	fileInfo, err := os.Stat(pathOfArchiveToSend)
	if err != nil {
		log.Fatal("failed to getting file size, pls try again")
	}

	if fileInfo.Size() > MAX_BYTES {
		log.Fatal("invalid file size, your file is bigger then 50mb, pls try again with a less file")
	}

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", os.Getenv("KINDLE_EMAIL"))
	m.Attach(pathOfArchiveToSend)

	fmt.Println("Processing your file, wait the process finish")
	err = d.DialAndSend(m)
	if err != nil {
		log.Fatal("error sending email")
	}

	fmt.Println("In a few minutes the file is available in your kindle library")
}
