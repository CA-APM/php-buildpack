package unit_test

import (
	"path/filepath"
	"time"

	"github.com/cloudfoundry/libbuildpack"
	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Options.JSON", func() {
	var versions []string
	var defaults map[string]interface{}
	BeforeEach(func() {
		bpDir, err := cutlass.FindRoot()
		Expect(err).NotTo(HaveOccurred())

		Expect(libbuildpack.NewJSON().Load(filepath.Join(bpDir, "defaults", "options.json"), &defaults)).To(Succeed())

		manifest, err := libbuildpack.NewManifest(bpDir, nil, time.Now())
		Expect(err).NotTo(HaveOccurred())
		versions = manifest.AllDependencyVersions("php")
	})

	It("PHP_72_LATEST will have the latest 7.2 version", func() {
		latest, err := libbuildpack.FindMatchingVersion("7.2.x", versions)
		Expect(err).NotTo(HaveOccurred())

		Expect(defaults["PHP_72_LATEST"]).To(Equal(latest))
	})

	It("PHP_73_LATEST will have the latest 7.3 version", func() {
		latest, err := libbuildpack.FindMatchingVersion("7.3.x", versions)
		Expect(err).NotTo(HaveOccurred())

		Expect(defaults["PHP_73_LATEST"]).To(Equal(latest))
	})

	It("PHP_74_LATEST will have the latest 7.3 version", func() {
		latest, err := libbuildpack.FindMatchingVersion("7.4.x", versions)
		Expect(err).NotTo(HaveOccurred())

		Expect(defaults["PHP_74_LATEST"]).To(Equal(latest))
	})
})
