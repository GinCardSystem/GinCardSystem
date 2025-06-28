package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

// AppConfig 全局配置结构体（根据实际配置项定义）
type AppConfig struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"database"`
}

type DatabaseConfig struct {
	Driver          string `mapstructure:"driver"` // 添加驱动类型字段
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	Location        string `mapstructure:"location"`          // 时区
	Timeout         int    `mapstructure:"timeout"`           // 连接超时（秒）
	MaxOpenConns    int    `mapstructure:"max_open_conns"`    // 最大打开连接数
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`    // 最大空闲连接数
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` // 连接最大生命周期（秒）
}

var (
	appConfig  *AppConfig
	configOnce sync.Once
)

// LoadConfigFile 加载配置文件（线程安全）
func LoadConfigFile() {
	configOnce.Do(func() {
		viper.SetConfigName("config")   // 文件名(不带后缀)
		viper.SetConfigType("yaml")     // 文件类型
		viper.AddConfigPath("./config") // 搜索路径

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Fatal error config file: %s \n", err)
		}

		// 解析配置到结构体
		cfg := &AppConfig{}
		if err := viper.Unmarshal(cfg); err != nil {
			log.Fatalf("Unable to decode config: %v \n", err)
		}

		appConfig = cfg
	})
}

// GetConfig 获取配置实例（线程安全）
func GetConfig() *AppConfig {
	if appConfig == nil {
		LoadConfigFile()
	}
	return appConfig
}
