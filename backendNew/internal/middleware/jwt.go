package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wofiporia/folium-backend-new/internal/util"
)

// JWTMiddleware JWT认证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "缺少Authorization头",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization格式错误，应为Bearer <token>",
			})
			c.Abort()
			return
		}

		// 验证token
		username, err := util.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "无效的token",
			})
			c.Abort()
			return
		}

		// 将用户名存储到上下文中，供后续处理器使用
		c.Set("username", username)
		c.Next()
	}
}

// OptionalJWTMiddleware 可选的JWT认证中间件（不强制要求token）
func OptionalJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.Next()
			return
		}

		username, err := util.ParseToken(tokenString)
		if err == nil {
			c.Set("username", username)
		}

		c.Next()
	}
}

