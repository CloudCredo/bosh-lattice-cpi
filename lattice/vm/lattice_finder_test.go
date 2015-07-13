package vm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	fakevm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm/fakes"
	fakeltcclient "github.com/cloudcredo/bosh-lattice-cpi/lattice/client/fakes"
)

var _ = Describe("LatticeFinder", func() {
	var (
		latticeClient          *fakeltcclient.FakeLatticeClient
		agentEnvServiceFactory *fakevm.FakeAgentEnvServiceFactory
		logger                 boshlog.Logger
		finder                 LatticeFinder
	)

	BeforeEach(func() {
		latticeClient = fakeltcclient.NewFakeLatticeClient()
		agentEnvServiceFactory = &fakevm.FakeAgentEnvServiceFactory{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		finder = NewLatticeFinder(
			latticeClient,
			agentEnvServiceFactory,
			logger,
		)
	})

	Describe("Find", func() {
		var (
			vmID int
		)

		Context("when the VM ID is valid and existing", func() {
			BeforeEach(func() {
				vmID = 1234
			})

			It("finds and returns a new LatticeVM object with correct ID", func() {
				vm, found, err := finder.Find(vmID)
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeTrue(), "could not find VM")
				Expect(vm.ID()).To(Equal(vmID), "found VM but ID does not match")
			})
		})

		Context("when the VM ID does not exist", func() {
			It("fails finding the VM", func() {
				_, found, _ := finder.Find(000000)
				Expect(found).To(BeFalse())
			})
		})
	})
})
