package config

import (
	"fmt"
	"github.com/Unknwon/com"
	consul "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"net/http"
)

var (
	LocalConfig ConfInfo
)

const (
	kAppName       = "APP_NAME"
	kConfigServer  = "CONFIG_SERVER"
	kConfigProfile = "CONFIG_PROFILE"
	kConfigType    = "CONFIG_TYPE"
)

type ConfInfo struct {
	Greetings struct {
		String string `json:"string"`
	} `json:"greetings"`
}

func GetConfig(configServerHost string, serverName string, profile string) ConfInfo {
	var config ConfInfo

	viper.AutomaticEnv()

	viper.SetDefault(kAppName, serverName)
	viper.SetDefault(kConfigServer, getConfigServiceFromConsul(configServerHost))
	viper.SetDefault(kConfigProfile, profile)
	viper.SetDefault(kConfigType, "yaml")

	if err := loadRemoteConfig(); err != nil {
		log.Fatal("Fail to load config", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Fail to parse config", err)
	}

	return config
}

func loadRemoteConfig() (err error) {
	confAddr := fmt.Sprintf("%v/%v-%v.%v",
		viper.Get(kConfigServer), viper.Get(kAppName), viper.Get(kConfigProfile), viper.Get(kConfigType))
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

func getConfigServiceFromConsul(configServerHost string) string {
	client, e := consul.NewClient(consul.DefaultConfig())
	if e != nil {
		log.Println("consul client create failed: ", e)
		return ""
	}
	rsp, _, e := client.Health().Service(configServerHost, "", true, nil)
	if e != nil {
		log.Println("consul request failed: ", e)
		return ""
	}
	if rsp == nil || len(rsp) == 0 {
		log.Println("config service not found in consul: ", configServerHost)
		return ""
	}
	service := rsp[rand.Int()%len(rsp)]
	return "http://" + service.Service.Address + ":" + com.ToStr(service.Service.Port)
}
