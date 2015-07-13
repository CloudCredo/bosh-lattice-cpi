package action_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"
)

var _ = Describe("DeleteStemcell", func() {
	var (
		action         DeleteStemcell
	)

	BeforeEach(func() {
		action = NewDeleteStemcell()
	})

	Describe("Run", func() {
		It("Does not raise an error", func() {
			_, err := action.Run(1234)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
