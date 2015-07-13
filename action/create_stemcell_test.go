package action_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"
)

var _ = Describe("CreateStemcell", func() {
	var (
		action         CreateStemcell
	)

	BeforeEach(func() {
		action = NewCreateStemcell()
	})

	Describe("Run", func() {
		It("returns id for created stemcell from image path", func() {
			_, err := action.Run("fake-path", CreateStemcellCloudProps{Uuid: "fake-stemcell-id"})
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
