package utils

import "os"

func GetGHRedirectURL() string {
	base := os.Getenv("BACKEND_URI")
	return base + "/v1/auth/github"
}
