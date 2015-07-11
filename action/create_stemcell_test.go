package action_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"

	fakestem "github.com/cloudcredo/bosh-lattice-cpi/softlayer/stemcell/fakes"
)

var _ = Describe("CreateStemcell", func() {
	var (
		stemcellFinder *fakestem.FakeFinder
		action         CreateStemcell
	)

	BeforeEach(func() {
		stemcellFinder = &fakestem.FakeFinder{}
		action = NewCreateStemcell(stemcellFinder)
	})

	Describe("Run", func() {
		It("returns id for created stemcell from image path", func() {
			stemcellFinder.FindFound, stemcellFinder.FindErr = true, nil
			stemcellFinder.FindStemcell = fakestem.NewFakeStemcell(1234, "fake-stemcell-id", fakestem.FakeStemcellKind)

			_, err := action.Run("fake-path", CreateStemcellCloudProps{Uuid: "fake-stemcell-id"})
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
