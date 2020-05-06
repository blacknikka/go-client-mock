package router

import (
	"net/http"
	"io"
	"fmt"

	"github.com/blacknikka/go-client-mock/usecase/session"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!")

	fmt.Println("hello world.")
}

func Init() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/session", session.CheckSession)
}

