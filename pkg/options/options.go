package options

import (
	"os"

	postgres "github.com/PandaGoL/api-project/internal/database/postgres/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Options struct {
	//logger settings
	LogLevel string
	Syslog   string

	//API settings
	APIAddr string

	DB *postgres.Options
}

var options *Options

func LoadConfig(configName string) (*Options, error) {
	log.Info("Try to load configuration file...")
	viper.SetConfigName(configName)
	viper.AddConfigPath("./internal/config")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Warnf("Configuration file not found")
			err = nil
		} else {
			return nil, err
		}
	}
	options = &Options{
		LogLevel: viper.GetString("log.level"),
		Syslog:   viper.GetString("log.syslog"),
		APIAddr:  viper.GetString("api.addr"),
		DB: &postgres.Options{
			Login:                 os.Getenv("DB_USER"),
			Password:              os.Getenv("DB_PASSWORD"),
			Host:                  os.Getenv("DB_HOST"),
			Database:              os.Getenv("DB_NAME"),
			QueryTimeout:          viper.GetDuration("storage.query_timeout"),
			MaxOpenConnections:    viper.GetInt("storage.max_open_connections"),
			MaxIdleConnections:    viper.GetInt("storage.max_idle_connections"),
			MaxConnectionLifetime: viper.GetDuration("storage.max_lifetime_connections"),
			MigrationEnable:       viper.GetBool("storage.migration_enable"),
			MigrationVersion:      viper.GetInt64("storage.migration_version"),
		},
	}

	return options, nil
}

func setDefaults() {
	viper.SetDefault("log.level", "DEBUG")
	viper.SetDefault("api.addr", "localhost:8888")
}

// Get - функция возвращает синглтон экземпляра опций
func Get() *Options {
	return options
}
