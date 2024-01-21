package token

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/fx"

	"github.com/takeuchi-shogo/ticket-api/auth"
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/pkg/uuid"
)

var Module = fx.Options(
	fx.Provide(NewJwtMaker),
)

type JwtMakerInterface interface {
	GenerateJWT(userID string) (*TokenPairs, error)
	VerifyJwtToken(token string) (*CustomClaims, error)
	NewNumericDate(hour int) *jwt.NumericDate
}

type JwtMaker struct {
	ApplicationName string
	TokenExpireAt   time.Duration
	PublicKey       []byte
	PrivateKey      []byte
}

type CustomClaims struct {
	UserID string `json:"uid"`
	*jwt.RegisteredClaims
}

type TokenPairs struct {
	AccessToken          string `json:"access_token"`
	RefreshToken         string `json:"refresh_token"`
	TokenExpireAt        int64  `json:"token_expire_at"`
	RefreshTokenExpireAt int64  `json:"refresh_token_expire_at"`
}

func NewJwtMaker(c config.ServerConfig) JwtMakerInterface {
	return &JwtMaker{
		TokenExpireAt: 15 * time.Minute,
		PublicKey:     auth.RawPublicKey,
		PrivateKey:    auth.RawPrivateKey,
	}
}

func (jm *JwtMaker) GenerateJWT(userID string) (*TokenPairs, error) {

	claims := &CustomClaims{
		userID,
		&jwt.RegisteredClaims{
			ExpiresAt: jm.NewNumericDate(24),
			IssuedAt:  jm.NewNumericDate(0),
			NotBefore: jm.NewNumericDate(0),
			Issuer:    "admin",
			Subject:   "tacketmaster",
			ID:        uuid.NewRandom(),
			Audience:  []string{jm.ApplicationName},
		},
	}

	claims.UserID = userID

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(jm.PrivateKey)
	if err != nil {
		return &TokenPairs{}, err
	}

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return &TokenPairs{}, err
	}

	refreshClaims := &CustomClaims{
		userID,
		&jwt.RegisteredClaims{
			ExpiresAt: jm.NewNumericDate(24 * 30),
			IssuedAt:  jm.NewNumericDate(0),
			NotBefore: jm.NewNumericDate(0),
			Issuer:    "admin",
			Subject:   "tacketmaster",
			ID:        uuid.NewRandom(),
			Audience:  []string{jm.ApplicationName},
		},
	}

	refreshClaims.UserID = userID

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)

	refreshTokenString, err := refreshToken.SignedString(signKey)
	if err != nil {
		return &TokenPairs{}, err
	}

	pairs := &TokenPairs{
		AccessToken:          tokenString,
		RefreshToken:         refreshTokenString,
		TokenExpireAt:        time.Now().Add(24 * time.Hour).Unix(),
		RefreshTokenExpireAt: time.Now().Add(24 * 30 * time.Hour).Unix(),
	}

	return pairs, nil
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
		key, err := jwt.ParseRSAPublicKeyFromPEM(jm.PublicKey)
		if err != nil {
			return nil, err
		}
		return key, nil
	})

	claims := token.Claims.(*CustomClaims)
	return claims, nil
}

func (c *CustomClaims) Validate() error {
	n := c.RegisteredClaims.ExpiresAt.Time
	t := time.Now()
	if t.After(n) {
		return errors.New("有効期限が切れています")
	}
	return nil
}

func (jm *JwtMaker) NewNumericDate(hour int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(hour) * time.Hour))
}
