package db

import (
	// "GinCardSystem/config"
	"GinCardSystem/internal/model"
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 全局数据库实例
var DB *gorm.DB

func Init() error {
	// 从配置获取数据库信息
	// cfg := config.GetConfig().Database

	// PostgreSQL DSN
	dsn := "host=192.168.10.115 user=GinCardSystem password=kdnb7RkQGpFCytDB dbname=GinCardSystem port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	/*
		*
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			cfg.Host,
			cfg.User,
			cfg.Password,
			cfg.DBName,
			cfg.Port,
			cfg.Location,
		)
	*/

	// 打开数据库连接
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("db connect error", "err", err)
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 获取底层 SQL DB 对象以设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		slog.Error("failed to get sql.DB", "err", err)
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	// sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// 自动迁移（创建表结构）
	err = DB.AutoMigrate(&model.UserModel{})
	if err != nil {
		slog.Error("auto migrate failed", "err", err)
		return fmt.Errorf("auto migrate failed: %w", err)
	}

	slog.Info("db connect success")
	return nil
}
