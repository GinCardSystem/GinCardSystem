package db

import (
	"GinCardSystem/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func main() {
	dsn := "root:123456@tcp(192.168.48.138:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		slog.Error("db connect error", err)
	}
	db.Create(model.UserModel{})
	slog.Info("db connect success")
}
