package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/ojiry/goth/internal/service"
)

type authorizeHandler struct{}

type authorizeErrorResponse struct {
	Error            string  `json:"error"`
	ErrorDescription *string `json:"error_description"`
	ErrorUri         *string `json:"error_uri"`
	State            *string `json:"state"`
}

var authenticateHtml = `
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

func (authorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		resp := authorizeErrorResponse{Error: "invalid_request"}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var ar service.AuthorizeRequest
	if err := mapFormToAuthorizeRequest(r.Form, &ar); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	as := service.NewAuthorizeService(ar)

	if err := as.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ed := err.Error()
		resp := authorizeErrorResponse{
			Error:            "invalid_request",
			ErrorDescription: &ed,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	t, err := template.New("").Parse(authenticateHtml)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err)
	}
}

func mapFormToAuthorizeRequest(f url.Values, ar *service.AuthorizeRequest) error {
	ar.Scope = f.Get("scope")
	ar.ResponseType = f.Get("response_type")
	ar.ClientID = f.Get("client_id")
	ar.RedirectUri = f.Get("redirect_uri")
	state := f.Get("state")
	ar.State = &state
	return nil
}
