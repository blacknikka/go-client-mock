package main

import (
	"log"
	"net/http"

	"github.com/blacknikka/go-client-mock/usecase"
)

func main() {
	// google
	client := &http.Client{}
	contentUsecase := usecase.NewContentUsecase(client)
	content, err := contentUsecase.GetContent()
	if err != nil {
		log.Println(err)
	}

	log.Println(content)

	// mock
	mockContent, err := usecase.GetContentUseMock()
	if err != nil {
		log.Println(err)
	}
	log.Println(mockContent)
}
