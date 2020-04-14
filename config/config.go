package config

import (
	"time"

	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("GOSHORTENER")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults

	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	// database connfigs
	v.SetDefault("db_driver", "sqlite3")
	v.SetDefault("db_connection_string", "goshortener.db")
	// CAVIAT: don't change this config if you set it once
	v.SetDefault("SHORTENER_BASE", "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// redis configs
	v.SetDefault("REDIS_ADDRESS", "127.0.0.1;6376")
	v.SetDefault("REDIS_POOL_SIZE", 100)
	v.SetDefault("REDIS_MAX_RETRIES", 2)
	v.SetDefault("REDIS_PASSWORD", "")
	v.SetDefault("REDIS_DB_NUMBER", 0)

	return v
}
