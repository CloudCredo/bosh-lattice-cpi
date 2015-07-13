package action_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/action"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

var _ = Describe("Networks", func() {
	var (
		networks Networks
	)

	BeforeEach(func() {
		networks = Networks{
			"fake-net1-name": Network{
				Type: "fake-net1-type",

				IP:      "fake-net1-ip",
				Netmask: "fake-net1-netmask",
				Gateway: "fake-net1-gateway",

				DNS:     []string{"fake-net1-dns"},
				Default: []string{"fake-net1-default"},

				CloudProperties: map[string]interface{}{
					"fake-net1-cp-key": "fake-net1-cp-value",
				},
			},
			"fake-net2-name": Network{
				Type: "fake-net2-type",
				IP:   "fake-net2-ip",
			},
		}
	})

	Describe("AsVMNetworks", func() {
		It("returns networks for VM", func() {
			expectedVMNetworks := bltcvm.Networks{
				"fake-net1-name": bltcvm.Network{
					Type: "fake-net1-type",

					IP:      "fake-net1-ip",
					Netmask: "fake-net1-netmask",
					Gateway: "fake-net1-gateway",

					DNS:     []string{"fake-net1-dns"},
					Default: []string{"fake-net1-default"},

					CloudProperties: map[string]interface{}{
						"fake-net1-cp-key": "fake-net1-cp-value",
					},
				},
				"fake-net2-name": bltcvm.Network{
					Type: "fake-net2-type",
					IP:   "fake-net2-ip",
				},
			}

			Expect(networks.AsVMNetworks()).To(Equal(expectedVMNetworks))
		})
	})
})
