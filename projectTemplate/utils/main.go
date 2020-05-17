package utils

import "log"

const (
	RoleAdmin = "admin"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
