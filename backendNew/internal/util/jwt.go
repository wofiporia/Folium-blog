package util

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// 使用更复杂的密钥，建议从环境变量获取
	defaultSecret = "WfP@2025!Blog#Secure*Key&Salt123456789"
	// Token过期时间：24小时
	tokenExpiration = 24 * time.Hour
)

// Claims 定义JWT claims结构
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetSecret 获取JWT密钥
func GetSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return defaultSecret
	}
	return secret
}

// GenerateToken 生成JWT token
func GenerateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetSecret()))
}

// ParseToken 解析JWT token，返回用户名
func ParseToken(tokenString string) (string, error) {
	// 基本格式检查
	if tokenString == "" {
		return "", errors.New("token为空")
	}
	
	if len(tokenString) < 50 {
		return "", errors.New("token格式不正确，长度过短")
	}

	// 检查token是否包含正确的JWT格式（三个部分用.分隔）
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", errors.New("token格式不正确，应包含三个部分")
	}

	// 严格验证签名算法和解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否为HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		}
		// 返回密钥
		secret := GetSecret()
		return []byte(secret), nil
	})

	// 如果解析出错，直接返回错误
	if err != nil {
		return "", errors.New("token解析失败: " + err.Error())
	}

	// 检查token是否有效
	if !token.Valid {
		return "", errors.New("token无效")
	}

	// 提取claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("无法提取token claims")
	}

	// 验证必要字段
	if claims.Username == "" {
		return "", errors.New("token中缺少用户名")
	}
	
	if claims.Subject == "" {
		return "", errors.New("token中缺少subject")
	}

	// 验证过期时间
	if claims.ExpiresAt == nil {
		return "", errors.New("token缺少过期时间")
	}
	
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("token已过期")
	}

	// 验证用户名和subject是否一致
	if claims.Username != claims.Subject {
		return "", errors.New("token中用户名和subject不匹配")
	}

	return claims.Username, nil
}

// ValidateToken 验证token是否有效
func ValidateToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
