package utils

import (
	"os/user"
	"path"
	"strings"
)

func EscapeHomePath(storagePath string) string {
	usr, _ := user.Current()
	if storagePath == "~" {
		storagePath = usr.HomeDir
	} else if strings.HasPrefix(storagePath, "~/") {
		storagePath = path.Join(usr.HomeDir, storagePath[2:])
	}
	return storagePath
}
