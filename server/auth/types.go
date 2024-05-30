package auth

import "math/rand"

const (
	Port      = ":8080"
	ServerUrl = "http://localhost:8080"
)

type ClientId string

type App struct {
	ClientId     ClientId
	ClientSecret string
	RedirectUri  string
	Scope        []string
}

type AccessCombination struct {
	State       string
	Code        int
	AccessToken int
}

func NewAccessCombination(state string) *AccessCombination {
	return &AccessCombination{
		State: state,
		Code:  rand.Int(),
		AccessToken:  rand.Int(),
	}
}
