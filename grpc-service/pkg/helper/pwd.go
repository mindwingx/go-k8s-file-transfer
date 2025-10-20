package helper

import (
	"os"
)

func Root() (pwd string) {
	pwd, _ = os.Getwd()
	return
}
