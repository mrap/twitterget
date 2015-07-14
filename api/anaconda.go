package api

import "github.com/ChimeraCoder/anaconda"

func NewApi(config AuthConfig) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	return anaconda.NewTwitterApi(config.AccessToken, config.AccessTokenSecret)
}
