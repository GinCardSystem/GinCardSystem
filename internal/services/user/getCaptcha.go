package user

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mojocn/base64Captcha"
)

// 包级配置 - 建议在应用启动时初始化
var (
	captchaSecret = []byte("your-strong-secret-key") // JWT签名密钥
	captchaExpire = time.Minute * 2                  // 验证码有效期
	captchaStore  = base64Captcha.DefaultMemStore    // 内存存储（可选备用）
)

// SetCaptchaConfig 初始化配置（密钥、有效期）
func SetCaptchaConfig(secret string, expire time.Duration) {
	captchaSecret = []byte(secret)
	captchaExpire = expire
}

// captchaClaims JWT自定义声明
type captchaClaims struct {
	Answer string `json:"ans"` // 验证码答案
	jwt.RegisteredClaims
}

// GenerateCaptcha 生成验证码（返回base64图片和JWT令牌）
func GenerateCaptcha() (imgBase64, token string, err error) {
	// 创建数字验证码驱动
	driver := base64Captcha.DriverDigit{
		Height:   60,
		Width:    200,
		Length:   5,
		MaxSkew:  0.5,
		DotCount: 80,
	}

	// 生成验证码
	c := base64Captcha.NewCaptcha(&driver, captchaStore)
	id, b64, ans, err := c.Generate()
	if err != nil {
		return "", "", err
	}
	_ = id // 不使用内部存储ID

	// 创建JWT令牌
	claims := captchaClaims{
		Answer: ans,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(captchaExpire)),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString(captchaSecret)
	if err != nil {
		return "", "", err
	}

	return b64, token, nil
}

// VerifyCaptcha 验证验证码（需提供JWT和用户输入）
func VerifyCaptcha(tokenString, userInput string) (bool, error) {
	// 解析JWT
	token, err := jwt.ParseWithClaims(
		tokenString,
		&captchaClaims{},
		func(t *jwt.Token) (interface{}, error) { return captchaSecret, nil },
	)
	if err != nil {
		return false, err
	}

	// 验证声明和有效期
	claims, ok := token.Claims.(*captchaClaims)
	if !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	// 比较答案（不区分大小写）
	if strings.EqualFold(claims.Answer, userInput) {
		return true, nil
	}
	return false, nil
}
