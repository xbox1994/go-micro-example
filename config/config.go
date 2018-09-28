package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	kAppName       = "APP_NAME"
	kConfigServer  = "CONFIG_SERVER"
	kConfigLabel   = "CONFIG_LABEL"
	kConfigProfile = "CONFIG_PROFILE"
	kConfigType    = "CONFIG_TYPE"
)

var (
	App AppConfig
)

type AppConfig struct {
	Name      string
	SearchUrl string `mapstructure:"search_url"`
}

func init() {
	viper.AutomaticEnv()
	initDefault()

	if err := loadRemoteConfig(); err != nil {
		log.Fatal("Fail to load config", err)
	}

	if err := sub("app", &App); err != nil {
		log.Fatal("Fail to parse config", err)
	}
}

func initDefault() {
	viper.SetDefault(kAppName, "greeter")
	viper.SetDefault(kConfigServer, "http://config:8080")
	viper.SetDefault(kConfigLabel, "config")
	viper.SetDefault(kConfigProfile, "default")
	viper.SetDefault(kConfigType, "yaml")
}

func loadRemoteConfig() (err error) {
	confAddr := fmt.Sprintf("%v/%v/%v-%v.%v",
		viper.Get(kConfigServer), viper.Get(kConfigLabel),
		viper.Get(kAppName), viper.Get(kConfigProfile),
		viper.Get(kConfigType))
	resp, err := http.Get(confAddr)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	viper.SetConfigType(viper.GetString(kConfigType))
	if err = viper.ReadConfig(resp.Body); err != nil {
		return
	}
	log.Println("Load config from: ", confAddr)
	return
}

func sub(key string, value interface{}) error {
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}
