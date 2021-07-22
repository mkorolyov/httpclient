package httpclient

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

func Test(t *testing.T) {
	RegisterTestingT(t)
	RunSpecs(t, "httpclient package")
}
