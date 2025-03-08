package internal

import (
	"fmt"
	"log"
	"os"
)

func AddIgnoreTemplate(fileName string, templateName string) {
	colorReset := "\033[0m"

	tr := NewTemplateRegistry()
	if !tr.HasTemplate(templateName) {
		log.Fatalf("template '%s' does not exist", templateName)
	}

	pathFile := "./" + fileName
	_, err := os.Stat(pathFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(fileName)
			if err != nil {
				log.Fatal("An error has occurred")
			}
		} else {
			log.Fatal("Error:", err)
		}
	}

	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tr.WriteTemplate(templateName, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("√  ", string(colorReset))
	fmt.Printf("%s %s created successfully\n", templateName, fileName)
}
