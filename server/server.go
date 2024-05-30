package main

import (
	"fmt"
	"log"
	"net/http"
	"server/auth"
	"server/resource"
)

const (
	Port      = ":8080"
	ServerUrl = "http://localhost:8080"
)

func main() { 
    http.HandleFunc("/authorize", auth.Authorize)
    http.HandleFunc("/login", login)
    http.HandleFunc("/consent", auth.RequestApproval)
    http.HandleFunc("/token", auth.GetAccessToken)
    http.HandleFunc("/access", resource.GrantAccess)

    log.Fatal(http.ListenAndServe(Port, nil))
}

func login(w http.ResponseWriter, r *http.Request) {
    
    params := r.URL.Query()
    clientId := params.Get("client_id")
    redirectUri := params.Get("redirect_uri")
    scope := params.Get("scope")
    state := params.Get("state")

    html := fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
        <title>Photo Gallery</title>
    </head>
    <body style="background-color: blue;">
        <h1>Photo Gallery</h1>
        <p>Print Service is requesting access to your photos.</p>
        <form method="post" action="/consent">
            <input type="hidden" name="client_id" value="%s">
            <input type="hidden" name="redirect_uri" value="%s">
            <input type="hidden" name="scope" value="%s">
            <input type="hidden" name="state" value="%s">
            <button type="submit" name="action" value="allow">Allow</button>
        </form>
    </body>
    </html>
    `, clientId, redirectUri, scope, state)

    fmt.Fprintf(w, html)
}
