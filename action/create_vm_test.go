package action_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"

	fakevm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm/fakes"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

var _ = Describe("CreateVM", func() {
	var (
		vmCreator      *fakevm.FakeCreator
		action         CreateVM
	)

	BeforeEach(func() {
		vmCreator = &fakevm.FakeCreator{}
		action = NewCreateVM(vmCreator)
	})

	Describe("Run", func() {
		var (
			vmCloudProp  bltcvm.VMCloudProperties
			networks     Networks
			env          Environment
		)

		BeforeEach(func() {
			vmCloudProp = bltcvm.VMCloudProperties{
			}
			networks = Networks{"fake-net-name": Network{IP: "fake-ip"}}
			env = Environment{"fake-env-key": "fake-env-value"}
		})

		It("tries to find stemcell with given stemcell cid", func() {
			vmCreator.CreateVM = fakevm.NewFakeVM(1234)

			_, err := action.Run("fake-agent-id", vmCloudProp, networks, env)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when stemcell is found with given stemcell cid", func() {
			It("returns id for created VM", func() {
				vmCreator.CreateVM = fakevm.NewFakeVM(1234)

				id, err := action.Run("fake-agent-id", vmCloudProp, networks, env)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(VMCID(1234)))
			})

			It("creates VM with requested agent ID, stemcell, cloud properties, and networks", func() {
				vmCreator.CreateVM = fakevm.NewFakeVM(1234)

				_, err := action.Run("fake-agent-id", vmCloudProp, networks, env)
				Expect(err).ToNot(HaveOccurred())

				Expect(vmCreator.CreateAgentID).To(Equal("fake-agent-id"))
				Expect(vmCreator.CreateVMCloudProperties).To(Equal(vmCloudProp))
				Expect(vmCreator.CreateNetworks).To(Equal(networks.AsVMNetworks()))
				Expect(vmCreator.CreateEnvironment).To(Equal(
					bltcvm.Environment{"fake-env-key": "fake-env-value"},
				))
			})

			It("returns error if creating VM fails", func() {
				vmCreator.CreateErr = errors.New("fake-create-err")

				id, err := action.Run("fake-agent-id", vmCloudProp, networks, env)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-create-err"))
				Expect(id).To(Equal(VMCID(0)))
			})
		})

		Context("when stemcell is not found with given cid", func() {
			It("returns error because VM cannot be created without a stemcell", func() {
				id, err := action.Run("fake-agent-id", vmCloudProp, networks, env)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Expected to find stemcell"))
				Expect(id).To(Equal(VMCID(0)))
			})
		})

		Context("when stemcell finding fails", func() {
			It("returns error because VM cannot be created without a stemcell", func() {
				id, err := action.Run("fake-agent-id", vmCloudProp, networks, env)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-find-err"))
				Expect(id).To(Equal(VMCID(0)))
			})
		})
	})
})
