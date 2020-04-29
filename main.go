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
	stopCh := make(chan struct{})
	usecase.AddAJob(usecase.TimerJob{
		Time:     (1000 * time.Millisecond),
		Chan:     resultCh,
		StopChan: stopCh,
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
		close(stopCh)
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
