package auth

var apps = map[ClientId]App{
    "printshop": {
        "printshop",
        "alpine",
        "http://localhost:8090/callback",
        []string{"all"},
    },
}

var AccessCombinations = make(map[ClientId]AccessCombination)
