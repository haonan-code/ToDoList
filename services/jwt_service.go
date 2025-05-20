package services

import (
	"bubble/models"
	"errors"
	"golang.org/x/oauth2/jwt"
	"time"
)

// 定义JWT密钥，生产环境请使用更安全的配置方式 (如环境变量)
var jwtSecret = []byte("your_super_secret_jwt_key")

// Claims 定义 JWT 的有效载荷
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 为用户生成 JWT
func GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 令牌有效期24小时

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),      // 签发时间
			Issuer: "your-api-issuer", // 签发者
			Subject: user.Username,     // 主题
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", ErrGenerateTokenFailed
	}
	return tokenString, nil
}

// ParseToken 解析和验证 JWT
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
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