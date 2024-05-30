package resource

import (
	"fmt"
	"net/http"
	"server/auth"
	"strconv"
)

func GrantAccess(w http.ResponseWriter, r *http.Request) {

    params := r.URL.Query()
    token, err := strconv.Atoi(params.Get("access_token"))
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    }

    for clientId, combo := range auth.AccessCombinations {
        if combo.AccessToken == token {
            data := secretData[clientId]
            fmt.Fprintf(w, data)
            return
        }
    }

    http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
