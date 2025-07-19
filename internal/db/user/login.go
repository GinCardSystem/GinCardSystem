package db

import (
	"GinCardSystem/internal/db"
	"GinCardSystem/internal/model"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 验证用户凭据并输出加密哈希
func VerifyUserCredentials(email, password string) (bool, *model.UserModel, error) {
	// 1. 动态生成密码的 bcrypt 哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to generate password hash", "error", err)
		return false, nil, err
	}

	// 输出加密的哈希值
	slog.Info("Generated bcrypt hash",
		"password", password,
		"hash", string(hashedPassword),
	)

	// 2. 查找用户
	var user model.UserModel
	result := db.DB.Where("email = ?", email).
		Where("account_status = ?", true).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			slog.Info("User not found or account disabled", "email", email)
			return false, nil, nil
		}
		slog.Error("Database query error", "error", result.Error)
		return false, nil, result.Error
	}

	// 3. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		slog.Info("Password mismatch", "email", email)
		return false, nil, nil
	}

	slog.Info("User authentication successful", "email", email)
	return true, &user, nil
}

// 单独生成并输出密码哈希的函数
func GenerateAndPrintHash(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to generate password hash", "error", err)
		return
	}

	slog.Info("Generated bcrypt hash",
		"password", password,
		"hash", string(hashedPassword),
	)
}
