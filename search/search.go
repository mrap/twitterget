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

func TweetsToUser(u user.User) []tweet.Tweet {
	reqURL := SearchURL
	_url.SetQueryParams(&reqURL, map[string]string{
		"q": "to:" + u.ScreenName,
		"f": "tweets",
	})

	res, err := http.Get(reqURL.String())
	PanicIf(err)
	root, err := html.Parse(res.Body)
	PanicIf(err)

	tweetsMatcher := func(n *html.Node) bool {
		return n.DataAtom == atom.Div && strings.HasPrefix(scrape.Attr(n, "class"), "tweet original-tweet")
	}
	tweetScreenNameMatcher := func(n *html.Node) bool {
		return n.DataAtom == atom.Span && strings.HasPrefix(scrape.Attr(n, "class"), "username")
	}
	tweetTextMatcher := func(n *html.Node) bool {
		return n.DataAtom == atom.P && strings.HasSuffix(scrape.Attr(n, "class"), "tweet-text")
	}

	tweetNodes := scrape.FindAll(root, tweetsMatcher)
	tweets := make([]tweet.Tweet, len(tweetNodes))
	for i, n := range tweetNodes {
		t := tweet.Tweet{}
		if child, ok := scrape.Find(n, tweetScreenNameMatcher); ok {
			t.Author = *user.NewUser(scrape.Text(child))
		}
		if child, ok := scrape.Find(n, tweetTextMatcher); ok {
			t.Text = scrape.Text(child)
		}
		tweets[i] = t
	}

	return tweets
}
