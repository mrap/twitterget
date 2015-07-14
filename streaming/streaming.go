package streaming

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func StartUserStream(api *anaconda.TwitterApi) anaconda.Stream {
	v := url.Values{}
	v.Set("with", "user")
	v.Set("replies", "all")
	return api.UserStream(v)
}
