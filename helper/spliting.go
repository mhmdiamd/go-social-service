package helper

import (
	"strings"
)

func GetFileExtention(name string) string {
  splitedName := strings.Split(name, ".")
  return splitedName[len(splitedName) - 1]
}
