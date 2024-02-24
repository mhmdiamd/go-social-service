package helper

import (
	"math/rand"
	"strconv"
)

func GenerateOTP() string {
  maxNumber := 9999
  minNumber := 1000
  randNumber := rand.Intn(maxNumber - minNumber + 1) + minNumber
  return strconv.Itoa(randNumber)
}

