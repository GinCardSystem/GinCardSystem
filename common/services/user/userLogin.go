package user

import (
	"GinCardSystem/common/response"
	db "GinCardSystem/internal/db/user"
	"GinCardSystem/internal/model"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JWT 密钥 - 实际应用中应从安全配置中获取
const jwtSecret = "your_strong_secret_key_here" // 替换为强密钥
const tokenExpiration = 24 * time.Hour          // Token 有效期 24 小时

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 自定义 JWT 声明
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		var login LoginUser
		if err := c.ShouldBindJSON(&login); err != nil {
			response.StatusUnauthorized(c, "Invalid input")
			return
		}

		slog.Info("Login attempt", "email", login.Email)

		// 验证用户凭证
		authenticated, user, err := db.VerifyUserCredentials(login.Email, login.Password)
		if err != nil {
			slog.Error("Database error during login", "error", err)
			response.StatusUnauthorized(c, "Server error")
			return
		}

		if !authenticated || user == nil {
			slog.Info("Authentication failed", "email", login.Email)
			response.StatusUnauthorized(c, "Invalid email or password, or account is disabled")
			return
		}

		// 生成 JWT Token
		token, err := generateJWTToken(user)
		if err != nil {
			slog.Error("Failed to generate JWT token", "error", err)
			response.StatusUnauthorized(c, "Failed to generate authentication token")
			return
		}

		slog.Info("Login successful", "email", login.Email, "user_id", user.ID)
		response.StatusSuccess(c, gin.H{
			"token":   token,
			"user_id": user.ID,
			"email":   user.Email,
		})
		return
	}

	if c.Request.Method == "GET" {
		// 示例响应，实际应用中可能需要其他逻辑
		response.StatusSuccess(c, gin.H{
			"message": "Login endpoint - use POST to authenticate",
		})
		return
	}
}

// 生成 JWT Token
func generateJWTToken(user *model.UserModel) (string, error) {
	// 设置声明
	claims := CustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "GinCardSystem",
			Subject:   "user_auth",
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名 token
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

// 验证 JWT Token 的中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.StatusUnauthorized(c, "Authorization header is required")
			c.Abort()
			return
		}

		// 检查格式 "Bearer <token>"
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			response.StatusUnauthorized(c, "Invalid authorization format")
			c.Abort()
			return
		}

		tokenString := authHeader[7:]

		// 解析和验证 token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			slog.Warn("Invalid token", "error", err)
			response.StatusUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		// 验证声明
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// 将用户信息添加到上下文中
			c.Set("userID", claims.UserID)
			c.Set("userEmail", claims.Email)
			c.Set("username", claims.Username)
			slog.Debug("User authenticated via JWT", "user_id", claims.UserID)
			c.Next()
		} else {
			response.StatusUnauthorized(c, "Invalid token claims")
			c.Abort()
		}
	}
}
