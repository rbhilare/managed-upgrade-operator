package ocmprovider

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

const baseUrl = "http://test.ocp"

var _ = Describe("Config", func() {
	var (
		mockCtrl    *gomock.Controller
		ocmProvider OcmProviderConfig
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		// Initialing temp value for config's baseurl
		ocmProvider = OcmProviderConfig{
			ConfigManager: ConfigManager{
				OcmBaseUrl: baseUrl,
			},
		}

	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("GetOCMBaseURL", func() {
		// Validating IsValid func
		It("Testing the IsValid function", func() {
			err := ocmProvider.IsValid()
			Expect(err).To(BeNil())
		})
		// Validating GetOCMBaseURL func
		It("Get base url", func() {
			url := ocmProvider.GetOCMBaseURL()
			Expect(url).To(Not(BeNil()))
		})

	})

})
