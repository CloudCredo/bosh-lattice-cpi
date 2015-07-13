package vm_test

import (
	. "github.com/onsi/ginkgo"

	. "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakeltcclient "github.com/cloudcredo/bosh-lattice-cpi/lattice/client/fakes"
)

var _ = Describe("LatticeAgentEnvService", func() {
	var (
		vmId            int
		latticeClient   *fakeltcclient.FakeLatticeClient
		agentEnvService LatticeAgentEnvService
		logger          boshlog.Logger
	)

	BeforeEach(func() {
		vmId = 1234567
		latticeClient = fakeltcclient.NewFakeLatticeClient()
		logger = boshlog.NewLogger(boshlog.LevelNone)
		agentEnvService = NewLatticeAgentEnvService(vmId, latticeClient, logger)
	})

	Context("#Fetch", func() {
		It("Returns an AgentEnv object built with current metadata when fetched", func() {
			//Implement me!
		})
	})

	Context("#Update", func() {
		It("Sets the VM's metadata using the AgentEnv object passed", func() {
			//Implement me!
		})
	})
})
