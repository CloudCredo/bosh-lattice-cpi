package action_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

var _ = Describe("ConcreteFactoryOptions", func() {
	var (
		options ConcreteFactoryOptions

		validOptions = ConcreteFactoryOptions{
			StemcellsDir: "/tmp/stemcells",

			Agent: bltcvm.AgentOptions{
				Mbus: "fake-mbus",
				NTP:  []string{},

				Blobstore: bltcvm.BlobstoreOptions{
					Type: "fake-blobstore-type",
				},
			},
		}
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			options = validOptions
		})

		It("returns error if StemcellsDir is empty", func() {
			options.StemcellsDir = ""

			err := options.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide non-empty StemcellsDir"))
		})

		It("returns error if agent section is not valid", func() {
			options.Agent.Mbus = ""

			err := options.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validating Agent configuration"))
		})
	})
})
