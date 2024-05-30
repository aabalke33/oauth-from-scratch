package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
)

func Authorize(w http.ResponseWriter, r *http.Request) {

    // Validate Method
    if r.Method != "GET" {
        http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
        return
    }

    //Validate the Request
    if params := r.URL.Query(); validAuthRequest(params) {

        //Create Access Combo
        clientId := params.Get("client_id")
        state := params.Get("state")

        combo := NewAccessCombination(state)
        AccessCombinations[ClientId(clientId)] = *combo

        //request user permission
        requestUserPermission(w, r)
        return
    }

    http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func validAuthRequest(params url.Values) bool {

    responseType := params.Get("response_type")
    clientId := ClientId(params.Get("client_id"))
    redirectUri := params.Get("redirect_uri")
    scope := params.Get("scope")

    if app, prs := apps[clientId];
        prs &&
        app.RedirectUri == redirectUri &&
        responseType == "code" &&
        slices.Contains(app.Scope, scope) {
        return true
    }

    return false
}

func requestUserPermission(w http.ResponseWriter, r *http.Request) {

    params := r.URL.Query()
    clientId := params.Get("client_id")
    redirectUri := params.Get("redirect_uri")
    scope := params.Get("scope")
    state := params.Get("state")

    url := fmt.Sprintf(
        "%s/login?client_id=%s&redirect_uri=%s&scope=%s&state=%s",
        ServerUrl, clientId, redirectUri, scope, state)

    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
