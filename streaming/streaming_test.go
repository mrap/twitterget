package streaming_test

import (
	"github.com/ChimeraCoder/anaconda"
	. "github.com/mrap/twitterget/streaming"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Streaming", func() {
	var stream anaconda.Stream

	BeforeEach(func() {
		stream = StartUserStream()
	})

	It("should return a stream with an open channel", func() {
		Eventually(stream.C).Should(Receive())
	})
})
