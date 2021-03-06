package stemcell_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudcredo/bosh-lattice-cpi/softlayer/stemcell"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	testhelpers "github.com/cloudcredo/bosh-lattice-cpi/test_helpers"
	fakesslclient "github.com/maximilien/softlayer-go/client/fakes"
)

var _ = Describe("SoftLayerStemcell", func() {
	var (
		softLayerClient *fakesslclient.FakeSoftLayerClient
		stemcell        SoftLayerStemcell
		logger          boshlog.Logger
	)

	BeforeEach(func() {
		softLayerClient = fakesslclient.NewFakeSoftLayerClient("fake-username", "fake-api-key")

		logger = boshlog.NewLogger(boshlog.LevelNone)

		stemcell = NewSoftLayerStemcell(1234, "fake-stemcell-uuid", DefaultKind, softLayerClient, logger)
	})

	Describe("#Delete", func() {
		BeforeEach(func() {
			fixturesFileNames := []string{"SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_Delete.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransaction.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
				"SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service_GetObject_None.json"}

			testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fixturesFileNames)
		})

		Context("when stemcell exists", func() {
			It("deletes the stemcell in collection directory that contains unpacked stemcell", func() {
				err := stemcell.Delete()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when stemcell does not exist", func() {
			BeforeEach(func() {
				softLayerClient.DoRawHttpRequestResponse = []byte("false")
			})

			It("returns error if deleting stemcell does not exist", func() {
				err := stemcell.Delete()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
