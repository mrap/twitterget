package api

import (
	"log"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

var MainAuthConfig *AuthConfig

func init() {
	MainAuthConfig = LoadAuthConfig()
}

func LoadAuthConfig() *AuthConfig {
	viper.SetConfigName("secrets")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Could not load auth config! You need a secrets.json file!")
		return nil
	}
	return &AuthConfig{
		ConsumerKey:       viper.GetString("twitter.consumer_key"),
		ConsumerSecret:    viper.GetString("twitter.consumer_secret"),
		AccessToken:       viper.GetString("twitter.access_token"),
		AccessTokenSecret: viper.GetString("twitter.access_token_secret"),
	}
}
