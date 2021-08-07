package server

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type pingHandler struct{}

func (pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

type authorizeHandler struct{}

func (authorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Authentication form</title>
</head>
<body>
    <h1>Authentication form</h1>
    <form action="/authenticate" method="post">
        <div>
            <label for="email">Email</label>
            <input type="email" name="email" id="email" required>
        </div>
        <div>
            <label for="password">Password</label>
            <input type="password" name="password" id="password" required>
        </div>
        <div>
            <input type="submit" value="Submit">
        </div>
    </form>
</body>
</html>
	`
	t, err := template.New("").Parse(html)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err)
	}
}

type authenticateHandler struct{}

func (authenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ping", http.StatusFound)
}

func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/ping", pingHandler{})
	mux.Handle("/authorize", authorizeHandler{})
	mux.Handle("/authenticate", authenticateHandler{})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}
