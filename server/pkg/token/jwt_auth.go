package token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/fx"

	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/pkg/uuid"
)

type JwtMakerInterface interface {
	GenerateJWT(userID string) (string, error)
	VerifyJwtToken(token string) (*CustomClaims, error)
	NewNumericDate(hour int) *jwt.NumericDate
}

type JwtMaker struct {
	ApplicationName string
	TokenExpireAt   time.Duration
	PublicKey       string
	PrivateKey      string
}

var Module = fx.Options(
	fx.Provide(NewJwtMaker),
)

type CustomClaims struct {
	UserID string `json:"uid"`
	*jwt.RegisteredClaims
}

func NewJwtMaker(c config.ServerConfig) JwtMakerInterface {
	return &JwtMaker{
		PrivateKey: "test",
	}
}

func (jm *JwtMaker) GenerateJWT(userID string) (string, error) {

	claims := &CustomClaims{
		userID,
		&jwt.RegisteredClaims{
			ExpiresAt: jm.NewNumericDate(24),
			IssuedAt:  jm.NewNumericDate(0),
			NotBefore: jm.NewNumericDate(0),
			Issuer:    "test test",
			Subject:   "tacketmaster",
			ID:        uuid.NewRandom(),
			Audience:  []string{jm.ApplicationName},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jm.PrivateKey))
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (jm *JwtMaker) VerifyJwtToken(token string) (*CustomClaims, error) {
	if _, err := extractBearerToken(token); err != nil {
		return nil, err
	}
	payload, err := jm.parseToken(token)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, ".")
	if len(jwtToken) != 3 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return header, nil
}

func (jm *JwtMaker) parseToken(jwtToken string) (*CustomClaims, error) {
	token, _ := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jm.PrivateKey), nil
	})

	claims := token.Claims.(*CustomClaims)
	return claims, nil
}

func (c *CustomClaims) Validate() error {
	n := c.RegisteredClaims.ExpiresAt.Time
	t := time.Now()
	if n.After(t) {
		return errors.New("有効期限が切れています")
	}
	return nil
}

func (jm *JwtMaker) NewNumericDate(hour int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(hour) * time.Hour))
}
