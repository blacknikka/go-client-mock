package main

import (
	"log"
	"net/http"
	"time"

	"github.com/blacknikka/go-client-mock/usecase"
)

func main() {
	// // google
	// client := &http.Client{}
	// contentUsecase := usecase.NewContentUsecase(client)
	// content, err := contentUsecase.GetContent()
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(content)

	// // mock
	// mockContent, err := usecase.GetContentUseMock()
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(mockContent)

	resultCh := make(chan string)
	usecase.AddAJob(usecase.TimerJob{
		Time: (1000 * time.Millisecond),
		Chan: resultCh,
		Job: func(ch chan string) {
			client := &http.Client{}
			contentUsecase := usecase.NewContentUsecase(client)
			content, err := contentUsecase.GetContent()
			if err != nil {
				ch <- err.Error()
			}

			ch <- content
		},
	})

	for {
		result := <-resultCh
		log.Println(result)
	}
}
