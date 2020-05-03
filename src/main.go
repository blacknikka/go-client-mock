package main

import (
	"context"
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

	ctx := context.Background()
	ctxParent, cancel := context.WithCancel(ctx)
	usecase.AddAJob(ctxParent, usecase.TimerJob{
		Time: (1000 * time.Millisecond),
		Chan: resultCh,
		Job: func(ch chan string) {
			client := &http.Client{}
			contentUsecase := usecase.NewContentUsecase(client)
			content, err := contentUsecase.GetContent()
			if err != nil {
				ch <- err.Error()
				return
			}

			ch <- content
		},
	})

	go func() {
		time.Sleep(time.Second * 5)

		log.Println("stop")
		cancel()
	}()

	for {
		result, ok := <-resultCh
		if ok == false {
			break
		}

		log.Println(result)
	}

	log.Println("done")
}
