package src

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func FatalError(mes string) {
	fmt.Println(mes)
	os.Exit(1)
}

func ExpandTilde(dirPath string) (string, error) {
	if strings.HasPrefix(dirPath, "~") {
		// Get the current user
		currUser, err := user.Current()
		if err != nil {
			return "", err
		}
		// Replace the ~ character with the home directory
		dirPath = strings.Replace(dirPath, "~", currUser.HomeDir, 1)
	}
	return dirPath, nil
}
