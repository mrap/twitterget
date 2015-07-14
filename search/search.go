package search

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/mrap/twitterget/api"
	"github.com/mrap/twitterget/tweet"
	"github.com/mrap/twitterget/user"
	_url "github.com/mrap/util/net/url"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	. "github.com/mrap/util/builtin"
)

var (
	SearchURL url.URL
)

func init() {
	SearchURL = api.BaseURL
	SearchURL.Path = "search"
}

func TweetsToUser(username string) []tweet.Tweet {
	u := SearchURL
	_url.SetQueryParams(&u, map[string]string{
		"q": "to:" + user.ToUsername(username),
		"f": "tweets",
	})

	res, err := http.Get(u.String())
	PanicIf(err)
	root, err := html.Parse(res.Body)
	PanicIf(err)

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.P {
			return strings.HasSuffix(scrape.Attr(n, "class"), "tweet-text")
		}
		return false
	}

	tweetNodes := scrape.FindAll(root, matcher)
	tweets := make([]tweet.Tweet, len(tweetNodes))
	for i, n := range tweetNodes {
		tweets[i] = tweet.Tweet{
			scrape.Text(n),
		}
	}

	return tweets
}
