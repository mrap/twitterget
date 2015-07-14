package search_test

import (
	. "github.com/mrap/twitterget/search"
	"github.com/mrap/twitterget/tweet"
	"github.com/mrap/twitterget/user"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting tweets addressed to a user", func() {
	var (
		usr    *user.User
		tweets []tweet.Tweet
	)

	JustBeforeEach(func() {
		usr = user.NewUser("the_mrap")
		tweets = TweetsToUser(*usr)
	})

	It("should return up to 20 tweets", func() {
		Expect(tweets).To(HaveLen(20))
	})

	It("should populate tweet authors", func() {
		for _, t := range tweets {
			Expect(t.Author.ScreenName).NotTo(BeEmpty())
		}
	})
})
