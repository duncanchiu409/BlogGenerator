package utils

import (
	"os"
	"strconv"
)

func CoordinatorSock() string {
	s := "/var/tmp/blogAI-"
	s += strconv.Itoa(os.Getuid())
	return s
}
