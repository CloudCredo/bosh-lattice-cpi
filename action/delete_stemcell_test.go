package action_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"

	fakestem "github.com/cloudcredo/bosh-lattice-cpi/softlayer/stemcell/fakes"
)

var _ = Describe("DeleteStemcell", func() {
	var (
		stemcellFinder *fakestem.FakeFinder
		action         DeleteStemcell
	)

	BeforeEach(func() {
		stemcellFinder = &fakestem.FakeFinder{}
		action = NewDeleteStemcell(stemcellFinder)
	})

	Describe("Run", func() {
		It("Does not raise an error", func() {
			_, err := action.Run(1234)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
