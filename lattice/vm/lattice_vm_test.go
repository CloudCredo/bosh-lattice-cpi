package vm_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"

	fakeltcclient "github.com/cloudcredo/bosh-lattice-cpi/lattice/client/fakes"
	fakevm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm/fakes"
	fakesutil "github.com/cloudcredo/bosh-lattice-cpi/util/fakes"
)

var _ = Describe("LatticeVM", func() {
	var (
		latticeClient   *fakeltcclient.FakeLatticeClient
		sshClient       *fakesutil.FakeSshClient
		agentEnvService *fakevm.FakeAgentEnvService
		logger          boshlog.Logger
		vm              LatticeVM
	)

	BeforeEach(func() {
		latticeClient = fakeltcclient.NewFakeLatticeClient()

		agentEnvService = &fakevm.FakeAgentEnvService{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		vm = NewLatticeVM(1234, latticeClient, sshClient, agentEnvService, logger)
	})

	Describe("Delete", func() {
		Context("valid VM ID is used", func() {
			BeforeEach(func() {
				vm = NewLatticeVM(1234567, latticeClient, sshClient, agentEnvService, logger)
			})

			It("deletes the VM successfully", func() {
				err := vm.Delete()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("invalid VM ID is used", func() {
			BeforeEach(func() {
				vm = NewLatticeVM(00000, latticeClient, sshClient, agentEnvService, logger)
			})

			It("fails deleting the VM", func() {
				err := vm.Delete()
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Reboot", func() {
		Context("valid VM ID is used", func() {
			BeforeEach(func() {
				vm = NewLatticeVM(1234567, latticeClient, sshClient, agentEnvService, logger)
			})

			It("reboots the VM successfully", func() {
				err := vm.Reboot()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("invalid VM ID is used", func() {
			BeforeEach(func() {
				vm = NewLatticeVM(00000, latticeClient, sshClient, agentEnvService, logger)
			})

			It("fails rebooting the VM", func() {
				err := vm.Reboot()
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("SetMetadata", func() {
		var (
			metadata VMMetadata
		)

		Context("no tags found in metadata", func() {
			BeforeEach(func() {
				metadataBytes := []byte(`{
				  "director": "fake-director-uuid",
				  "name": "fake-director"
				}`)

				metadata = bltcvm.VMMetadata{}
				err := json.Unmarshal(metadataBytes, &metadata)
				Expect(err).ToNot(HaveOccurred())
			})

			It("does not set any tag values on the VM", func() {
				err := vm.SetMetadata(metadata)

				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("found tags in metadata", func() {
			BeforeEach(func() {
				metadataBytes := []byte(`{
				  "director": "fake-director-uuid",
				  "name": "fake-director",
				  "tags": "test, tag, director"
				}`)

				metadata = bltcvm.VMMetadata{}
				err := json.Unmarshal(metadataBytes, &metadata)
				Expect(err).ToNot(HaveOccurred())
			})

			It("the tags value is empty", func() {
				metadata["tags"] = ""
				err := vm.SetMetadata(metadata)

				Expect(err).ToNot(HaveOccurred())
			})

			It("at least one tag found", func() {
				err := vm.SetMetadata(metadata)

				Expect(err).ToNot(HaveOccurred())
			})

			Context("when SLVG.SetTags call fails", func() {
				It("fails with error", func() {
					err := vm.SetMetadata(metadata)

					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("ConfigureNetworks", func() {
		var (
			networks Networks
		)

		BeforeEach(func() {
			networks = Networks{}
			vm = NewLatticeVM(1234567, latticeClient, sshClient, agentEnvService, logger)
		})

		It("returns NotSupportedError", func() {
			err := vm.ConfigureNetworks(networks)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Not supported"))
			Expect(err.(NotSupportedError).Type()).To(Equal("Bosh::Clouds::NotSupported"))
		})
	})
})
