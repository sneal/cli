package translatableerror_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTranslatableerror(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Translatable Error Suite")
}
