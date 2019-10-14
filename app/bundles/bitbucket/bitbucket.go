package bitbucket

import (
	"fmt"
)

// HandlePROpen gets the next prime after the given number
func HandlePROpen(payload string) string {
	fmt.Println(payload)
	return "yes"
}
