package session

import (
	"net/http"
	"fmt"
	"time"
	"github.com/gorilla/sessions"
)

const (
	sessionName = "session-aaaaa"
	sessionKey = "session"
)

var (
	store = sessions.NewCookieStore([]byte(sessionKey))
)

func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Fprintf(w, "something wrong: %s", err.Error())
		return
	}

	before := session.Values["time"]
	after := time.Now().Format(time.RFC3339)

	session.Values["time"] = after
	if err := session.Save(r, w); err != nil {
		fmt.Fprintf(w, "error was occurred when saving the session: %s", err.Error())
		return
	}

	fmt.Fprint(w, "session is ok\n")
	fmt.Fprintf(w, "time before: %s\n", before)
	fmt.Fprintf(w, "time after: %s\n", after)
}
