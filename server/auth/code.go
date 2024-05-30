package auth

import (
	"fmt"
	"net/http"
)

func RequestApproval(w http.ResponseWriter, r *http.Request) {

    // Validate Method
    if r.Method != "POST" {
        http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
        return
    }

    //Parse Form
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Unauthorized", http.StatusMethodNotAllowed)
        return
    }

    // Validate Request
    if clientId, ok := validRequestApproval(r); ok {
        sendCode(w, r, clientId)
        return
    }

    http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func validRequestApproval(r *http.Request) (ClientId, bool) {

    form := r.Form
    action := form.Get("action")
    state := form.Get("state")
    clientId := ClientId(form.Get("client_id"))

    storedState := AccessCombinations[clientId].State

    if action == "allow" &&
        state == storedState {
            return clientId, true
    }

    return clientId, false
}

func sendCode(w http.ResponseWriter, r *http.Request, clientId ClientId) {

    code := AccessCombinations[clientId].Code
    redirectUrl := apps[clientId].RedirectUri
    state := r.Form.Get("state")

    url := fmt.Sprintf("%s?code=%d&state=%s", redirectUrl, code, state)

    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
