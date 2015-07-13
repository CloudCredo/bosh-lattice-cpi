package action_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	fakecmd "github.com/cloudfoundry/bosh-utils/fileutil/fakes"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"

	fakeltcclient "github.com/cloudcredo/bosh-lattice-cpi/lattice/client/fakes"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

var _ = Describe("concreteFactory", func() {
	var (
		latticeClient   *fakeltcclient.FakeLatticeClient
		fs              *fakesys.FakeFileSystem
		cmdRunner       *fakesys.FakeCmdRunner
		compressor      *fakecmd.FakeCompressor
		logger          boshlog.Logger

		options = ConcreteFactoryOptions{
			StemcellsDir: "/tmp/stemcells",
		}

		factory Factory
	)

	var (
		agentEnvServiceFactory bltcvm.AgentEnvServiceFactory

		vmFinder       bltcvm.Finder
	)

	BeforeEach(func() {
		latticeClient = fakeltcclient.NewFakeLatticeClient()
		fs = fakesys.NewFakeFileSystem()
		cmdRunner = fakesys.NewFakeCmdRunner()
		compressor = fakecmd.NewFakeCompressor()
		logger = boshlog.NewLogger(boshlog.LevelNone)

		factory = NewConcreteFactory(
			latticeClient,
			options,
			logger,
		)
	})

	BeforeEach(func() {
		agentEnvServiceFactory = bltcvm.NewLatticeAgentEnvServiceFactory(latticeClient, logger)

		vmFinder = bltcvm.NewLatticeFinder(
			latticeClient,
			agentEnvServiceFactory,
			logger,
		)
	})

	Context("Stemcell methods", func() {
		It("create_stemcell", func() {
			_, err := factory.Create("create_stemcell")
			Expect(err).ToNot(HaveOccurred())
		})

		It("delete_stemcell", func() {
			_, err := factory.Create("delete_stemcell")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("VM methods", func() {
		It("create_vm", func() {
			vmCreator := bltcvm.NewLatticeCreator(
				latticeClient,
				agentEnvServiceFactory,
				options.Agent,
				logger,
			)

			action, err := factory.Create("create_vm")
			Expect(err).ToNot(HaveOccurred())
			Expect(action).To(Equal(NewCreateVM(vmCreator)))
		})

		It("delete_vm", func() {
			action, err := factory.Create("delete_vm")
			Expect(err).ToNot(HaveOccurred())
			Expect(action).To(Equal(NewDeleteVM(vmFinder)))
		})

		It("has_vm", func() {
			action, err := factory.Create("has_vm")
			Expect(err).ToNot(HaveOccurred())
			Expect(action).To(Equal(NewHasVM(vmFinder)))
		})

		It("reboot_vm", func() {
			action, err := factory.Create("reboot_vm")
			Expect(err).ToNot(HaveOccurred())
			Expect(action).To(Equal(NewRebootVM(vmFinder)))
		})

		It("set_vm_metadata", func() {
			action, err := factory.Create("set_vm_metadata")
			Expect(err).ToNot(HaveOccurred())
			Expect(action).To(Equal(NewSetVMMetadata(vmFinder)))
		})

		It("configure_networks", func() {
			action, err := factory.Create("configure_networks")
			Expect(err).ToNot(HaveOccurred())
			Expect(action).To(Equal(NewConfigureNetworks(vmFinder)))
		})
	})

	Context("Unsupported methods", func() {
		It("creates an iSCSI disk", func() {
			action, err := factory.Create("create_disk")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("deletes the detached iSCSI disk", func() {
			action, err := factory.Create("delete_disk")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("attaches an iSCSI disk to a virtual guest", func() {
			action, err := factory.Create("attach_disk")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("detaches the iSCSI disk from virtual guest", func() {
			action, err := factory.Create("detach_disk")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("returns error because CPI machine is not self-aware if action is current_vm_id", func() {
			action, err := factory.Create("current_vm_id")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("returns error because snapshotting is not implemented if action is snapshot_disk", func() {
			action, err := factory.Create("snapshot_disk")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("returns error because snapshotting is not implemented if action is delete_snapshot", func() {
			action, err := factory.Create("delete_snapshot")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("returns error since CPI should not keep state if action is get_disks", func() {
			action, err := factory.Create("get_disks")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})

		It("returns error because ping is not official CPI method if action is ping", func() {
			action, err := factory.Create("ping")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})
	})

	Context("Misc", func() {
		It("returns error if action cannot be created", func() {
			action, err := factory.Create("fake-unknown-action")
			Expect(err).To(HaveOccurred())
			Expect(action).To(BeNil())
		})
	})
})
