package configs

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Cfg struct {
	HTTPServer *HTTPServer
	MongoDB    *MongoDB
}

type HTTPServer struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MongoDB struct {
	Protocol       string
	Host           string
	Port           int
	Username       string
	Password       string
	WriteConcern   string
	MaxPoolSize    int
	ConnectTimeout int
}

const (
	configName = "wajve.conf"
	configPath = "configs"
)

func Init() (*Cfg, error) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath + "/dev") // for local development
	viper.AddConfigPath(configPath)          // base config file directory

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper.ReadInConfig.error: %w", err)
	}

	cfg := &Cfg{
		HTTPServer: &HTTPServer{
			Host:         viper.GetString("http_server.host"),
			Port:         viper.GetString("http_server.port"),
			ReadTimeout:  viper.GetDuration("http_server.read_timeout"),
			WriteTimeout: viper.GetDuration("http_server.write_timeout"),
		},
		MongoDB: &MongoDB{
			Protocol:       viper.GetString("mongo.protocol"),
			Host:           viper.GetString("mongo.host"),
			Port:           viper.GetInt("mongo.port"),
			Username:       viper.GetString("mongo.username"),
			Password:       viper.GetString("mongo.password"),
			WriteConcern:   viper.GetString("mongo.write_concern"),
			MaxPoolSize:    viper.GetInt("mongo.max_pool_size"),
			ConnectTimeout: viper.GetInt("mongo.connect_timeout"),
		},
	}

	return cfg, nil
}
