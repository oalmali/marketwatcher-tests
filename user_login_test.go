package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

	})

	var loginPage = map[string]string{
		"email":    "#email",
		"password": "#password",
		"submit":   "Submit",
	}

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	// Test Feature

	It("should direct to login page", func() {
		By("redirecting the user to the login form from the home page", func() {
			Expect(page.Navigate("http://marketwatcher.tech")).To(Succeed())
			Eventually(page).Should(HaveURL("http://marketwatcher.tech/landing?redirect=%2Fdashboard"))
		})

	})

	It("should manage user Authentication", func() {
		By("loging in user with correct username and password", func() {
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())

		})

	})

	It("should manage user Authentication", func() {
		By("loging in user with correct username and password", func() {
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("wrongPassword")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})

	})
})
