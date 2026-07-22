package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct { 
  UserID int64 `json:"user_id"`
  jwt.RegisteredClaims
}

func GenerateToken (id int64) (string, error) { 
  key := []byte(os.Getenv("JWT_SECRET"))
  claims := MyCustomClaims{ 
    id,
    jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  ss, err := token.SignedString(key)
  if err != nil { 
    return "", nil
  }
  return ss, nil
}

func VerifyToken (token string) error { 
  x, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(x *jwt.Token) (any, error) { 
    return []byte(os.Getenv("JWT_SECRET")), nil
  })
  if err != nil {
     return err
  }
  if !x.Valid {
     return fmt.Errorf("invalid token")
  }
  return nil
}