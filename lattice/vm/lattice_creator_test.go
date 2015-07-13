package vm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"

	fakevm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm/fakes"
	fakeltclient "github.com/cloudcredo/bosh-lattice-cpi/lattice/client/fakes"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

var _ = Describe("LatticeCreator", func() {
	var (
		latticeClient          *fakeltclient.FakeLatticeClient
		agentEnvServiceFactory *fakevm.FakeAgentEnvServiceFactory
		agentOptions           AgentOptions
		logger                 boshlog.Logger
		creator                LatticeCreator
	)

	BeforeEach(func() {
		latticeClient = fakeltclient.NewFakeLatticeClient()

		agentEnvServiceFactory = &fakevm.FakeAgentEnvServiceFactory{}
		agentOptions = AgentOptions{Mbus: "fake-mbus"}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		creator = NewLatticeCreator(
			latticeClient,
			agentEnvServiceFactory,
			agentOptions,
			logger,
		)
	})

	Describe("Create", func() {
		var (
			agentID    string
			cloudProps VMCloudProperties
			networks   Networks
			env        Environment
		)

		Context("valid arguments", func() {
			BeforeEach(func() {
				agentID = "fake-agent-id"
				cloudProps = VMCloudProperties{}
				networks = Networks{}
				env = Environment{}
			})

			It("returns a new LatticeVM with correct virtual guest ID and LatticeClient", func() {
				vm, err := creator.Create(agentID, cloudProps, networks, env)
				Expect(err).ToNot(HaveOccurred())
				Expect(vm.ID()).To(Equal(1234567))
			})
		})
	})
})
