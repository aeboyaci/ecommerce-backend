package env

import "os"

var JWT_SECRET = os.Getenv("JWT_SECRET")
