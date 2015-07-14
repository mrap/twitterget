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

func LoadAuthConfig(name string) *AuthConfig {
	viper.SetConfigName(name)
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Could not load auth config! Could not find file named:", name)
		return nil
	}
	return &AuthConfig{
		ConsumerKey:       viper.GetString("twitter.consumer_key"),
		ConsumerSecret:    viper.GetString("twitter.consumer_secret"),
		AccessToken:       viper.GetString("twitter.access_token"),
		AccessTokenSecret: viper.GetString("twitter.access_token_secret"),
	}
}
