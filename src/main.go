package main

import (
	"log"
	"net/http"

	"github.com/blacknikka/go-client-mock/router"
)

func main() {
	router.Init()
	log.Fatal(http.ListenAndServe(":5000", nil))

	// resultCh := make(chan string)

	// ctx := context.Background()
	// ctxParent, cancel := context.WithCancel(ctx)
	// usecase.AddAJob(ctxParent, usecase.TimerJob{
	// 	Time: (1000 * time.Millisecond),
	// 	Chan: resultCh,
	// 	Job: func(ch chan string) {
	// 		client := &http.Client{}
	// 		contentUsecase := usecase.NewContentUsecase(client)
	// 		content, err := contentUsecase.GetContent("http://json/posts")
	// 		if err != nil {
	// 			ch <- err.Error()
	// 			return
	// 		}

	// 		ch <- content
	// 	},
	// })

	// go func() {
	// 	time.Sleep(time.Second * 5)

	// 	log.Println("stop")
	// 	cancel()
	// }()

	// for {
	// 	result, ok := <-resultCh
	// 	if ok == false {
	// 		break
	// 	}

	// 	log.Println(result)
	// }

	// log.Println("done")
}
