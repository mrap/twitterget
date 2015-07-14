package search_test

import (
	. "github.com/mrap/twitterget/search"
	"github.com/mrap/twitterget/tweet"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting tweets to a user", func() {
	var (
		username string
		tweets   []tweet.Tweet
	)

	JustBeforeEach(func() {
		username = "the_mrap"
		tweets = TweetsToUser(username)
	})

	It("should return up to 20 tweets", func() {
		Expect(tweets).To(HaveLen(20))
	})
})
