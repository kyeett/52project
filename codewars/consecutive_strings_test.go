package main

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func LongestConsec(strarr []string, k int) string {
	var s, largest string

	for i := 0; i < len(strarr)-k+1; i++ {
		s = strings.Join(strarr[i:i+k], "")
		if len(s) > len(largest) {
			largest = s
		}
	}
	return largest
}

func dotest(strarr []string, k int, exp string) {
	var ans = LongestConsec(strarr, k)
	Expect(ans).To(Equal(exp))
}

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta")
		dotest([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
			"oocccffuucccjjjkkkjyyyeehh")
		dotest([]string{}, 3, "")
		dotest([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	})
})
