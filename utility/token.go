package utility

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(public_id, secret string) (tokenString string, err error) {

  jwtClaim := jwt.MapClaims{
    "id" : public_id,
  }

  tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

  tokenString, err = tokenClaim.SignedString([]byte(secret))
  if err != nil {
    return "", err
  }

  return 
}

func ValidateToken(tokenString, secret string) (id string, err error) {

  tokens, err := jwt.Parse(tokenString, func (j *jwt.Token) (interface{}, error) {
    if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method")
    }

    return []byte(secret), nil
  })


  if err != nil {
    return 
  }

  claims, ok := tokens.Claims.(jwt.MapClaims)
  if ok && tokens.Valid {
    id = fmt.Sprintf("%v", claims["id"])

    return 
  }

  err = fmt.Errorf("unable to extract claims")

  return
}
