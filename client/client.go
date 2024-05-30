package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const (
	Port      = ":8090"
	ClientUrl = "http://localhost:8090/callback"
	ServerUrl = "http://localhost:8080"
)

var (
    clientId = ""
    clientSecret = ""
    scope = ""
)

type Token struct {
    AccessToken string `json:"access_token"`
    TokenType string `json:"token_type"`
}

func main() { 
    http.HandleFunc("/", root)
    http.HandleFunc("/submit", submit)
    http.HandleFunc("/callback", callback)

    log.Fatal(http.ListenAndServe(Port, nil))
}

func root(w http.ResponseWriter, r *http.Request) {

    state := rand.Int()

    html := fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
        <title>Print Shop</title>
    </head>
    <body>
        <h1>Print Shop</h1>
        <p>Share Photo data with Print Shop</p>
        <form method="get" action="/submit">
            <input type="hidden" name="response_type" value="code">
            <input type="hidden" name="state" value="%d">
            <input type="hidden" name="redirect_uri" value="%s">
            <label for="client_id">Client Id:</label>
            <input type="text" name="client_id" value="">
            <label for="client_secret">Client Secret:</label>
            <input type="text" name="client_secret" value="">
            <label for="scope">Scope:</label>
            <input type="text" name="scope" value="">
            <button type="submit" name="action" value="allow">Allow</button>
        </form>
    </body>
    </hmtl>
    `, state, ClientUrl)

    fmt.Fprintf(w, html)
}

func submit(w http.ResponseWriter, r *http.Request) {

    // Validate Method
    if r.Method != "GET" {
        http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
        return
    }

    //Pull Params we want to keep
    params := r.URL.Query()
    responseType := params.Get("response_type")
    clientId = params.Get("client_id")
    clientSecret = params.Get("client_secret")
    redirectUri := params.Get("redirect_uri")
    scope = params.Get("scope")
    state := params.Get("state")

    //Redirect to server
    url := fmt.Sprintf(
        "%s/authorize?response_type=%s&client_id=%s&redirect_uri=%s&scope=%s&state=%s",
        ServerUrl, responseType, clientId, redirectUri, scope, state)

    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callback(w http.ResponseWriter, r *http.Request) {

    params := r.URL.Query()

    token := getAccessToken(w, params)
    secret := getProtectedData(w, token)

    html := fmt.Sprintf(
    `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Print Shop</title>
    </head>
    <body>
        <h1>Super Secret</h1>
        <p>%s</p>
    </body>
    </html>
    `, secret)

    fmt.Fprintf(w, html)
}

func getAccessToken(w http.ResponseWriter, params url.Values) Token {

    form := url.Values{
        "grant_type": []string{"authorization_code"},
        "code": []string{params.Get("code")},
        "redirect_uri": []string{ClientUrl},
    }

    uri := fmt.Sprintf("%s/token", ServerUrl)

    req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))

    req.SetBasicAuth(clientId, clientSecret)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
    }
    defer res.Body.Close()

    data := ""

    scanner := bufio.NewScanner(res.Body)
    for scanner.Scan() {
        data += scanner.Text()
    }

    token := Token{}
    json.Unmarshal([]byte(data), &token)

    return token
}

func getProtectedData(w http.ResponseWriter, token Token) string {

    url := fmt.Sprintf("%s/access?access_token=%s", ServerUrl, token.AccessToken)
    res, err := http.Get(url)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
    }
    defer res.Body.Close()

    data := ""
    scanner := bufio.NewScanner(res.Body)
    for scanner.Scan() {
        data += scanner.Text()
    }

    return data
}

