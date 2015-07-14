package streaming

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	. "github.com/mrap/twitterget/api"
)

func NewApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(MainAuthConfig.ConsumerKey)
	anaconda.SetConsumerSecret(MainAuthConfig.ConsumerSecret)
	return anaconda.NewTwitterApi(MainAuthConfig.AccessToken, MainAuthConfig.AccessTokenSecret)
}

func StartUserStream() anaconda.Stream {
	api := NewApi()
	v := url.Values{}
	v.Set("with", "user")
	v.Set("replies", "all")
	return api.UserStream(v)
}
