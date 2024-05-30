package auth

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) {

    // Validate Method
    if r.Method != "POST" {
        http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
        return
    }

    // Parse Form
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Unauthorized", http.StatusMethodNotAllowed)
        return
    }

    //If valid Code Request
    if clientId, ok := validCodeRequest(r); ok {
        grantAccessToken(w, clientId)
        return
    }

    http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func validCodeRequest(r *http.Request) (ClientId, bool) {

    params := r.Form
    grantType := params.Get("grant_type")
    code, err := strconv.Atoi(params.Get("code"))
    redirectUri := params.Get("redirect_uri")

    clientIdStr, clientSecret, ok := r.BasicAuth()
    clientId := ClientId(clientIdStr)
    if !ok {
        return clientId, false
    }

    app, prs := apps[clientId]
    accessCombination, prs := AccessCombinations[clientId]

    if err == nil &&
        prs &&
        grantType == "authorization_code" &&
        redirectUri == app.RedirectUri &&
        code == accessCombination.Code &&
        clientSecret == app.ClientSecret {
        return clientId, true
    }

    return clientId, false
}

func grantAccessToken(w http.ResponseWriter, clientId ClientId) {

    accessToken := strconv.Itoa(AccessCombinations[clientId].AccessToken)

    data := map[string]string{
        "access_token": accessToken,
        "token_type": "example",
    }

    bytes, err := json.Marshal(data)
    if err != nil {
        http.Error(w, "Server Error", http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
    w.WriteHeader(http.StatusOK)
    w.Write(bytes)
}
