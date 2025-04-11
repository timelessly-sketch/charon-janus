package token

import (
	"charon-janus/internal/model"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	key                    = []byte("charon-janus")
	codeGenerateTokenError = gerror.New("生成Token失败")
)

type CustomClaims struct {
	model.Identity
	jwt.RegisteredClaims
}

func GenerateJWT(ctx context.Context, user *model.Identity) (token string, err error) {
	var (
		now      = time.Now()
		issuer   = "charon-janus"
		duration = time.Second * gconv.Duration(7*24*time.Hour)
	)

	claims := CustomClaims{
		Identity: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    issuer,
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		g.Log().Error(ctx, err)
		return "", codeGenerateTokenError
	}
	return
}

func ValidateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
