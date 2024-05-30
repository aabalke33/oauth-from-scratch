package resource

import (
    "server/auth"
)

var secretData = map[auth.ClientId]string{
    "printshop": "super secret photos",
}
