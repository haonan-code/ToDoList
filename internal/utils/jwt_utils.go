package utils

import (
	"bubble/configs"
	"bubble/internal/model"
	service "bubble/pkg/serve/service/account"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Claims 定义 JWT 的有效载荷
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 为用户生成 JWT
func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 令牌有效期24小时

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 签发时间
			Issuer:    "localhost",                        // 签发者
			Subject:   user.Username,                      // 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return "", service.ErrGenerateTokenFailed
	}
	return tokenString, nil
}

// ParseToken 解析和验证 JWT
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("无效的令牌签名")
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("令牌已过期")
		}
		return nil, errors.New("无效的令牌")
	}

	if !token.Valid {
		return nil, errors.New("无效的令牌")
	}

	return claims, nil
}
