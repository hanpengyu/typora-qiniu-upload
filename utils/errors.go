package utils

import "fmt"

func ErrorString(prefix string, msg string) string {
	errStr := fmt.Sprintf("%s\n", prefix)
	errStr += fmt.Sprintf("%s\n", msg)
	return errStr
}

func SuccessString(prefix string, msgs []string) string {
	succStr := fmt.Sprintf("%s\n", prefix)
	for _, msg := range msgs {
		succStr += fmt.Sprintf("%s\n", msg)
	}
	return succStr
}
