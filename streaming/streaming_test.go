package streaming_test

import (
	"github.com/ChimeraCoder/anaconda"
	. "github.com/mrap/twitterget/api"
	. "github.com/mrap/twitterget/streaming"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Streaming", func() {
	var (
		api    *anaconda.TwitterApi
		stream anaconda.Stream
	)

	BeforeEach(func() {
		api = NewApi(*LoadAuthConfig("secrets"))
		stream = StartUserStream(api)
	})

	It("should return a stream with an open channel", func() {
		Eventually(stream.C).Should(Receive())
	})
})
