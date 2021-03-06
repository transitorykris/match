package match

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrimAddress(t *testing.T) {
	Convey("When checking if a rune is an IPv6 or IPv4 character", t, func() {
		Convey("And it is", func() {
			const chars = "abcdefABCDEF0123456789:"
			for _, c := range chars {
				So(TrimAddress(c), ShouldBeFalse)
			}
		})
		Convey("And it isn't", func() {
			// We'll just check a few...
			const chars = "!@#$/QWRTYzxv,"
			for _, c := range chars {
				So(TrimAddress(c), ShouldBeTrue)
			}
		})
	})
}

func TestIP(t *testing.T) {
	Convey("When looking for IPv6 addresses in a string", t, func() {
		Convey("And it does not contain one", func() {
			matches := IP("This does not contain an IPv6 address")
			So(len(matches), ShouldEqual, 0)
		})
		Convey("And it is an IPv6 address only", func() {
			matches := IP("1:2:3::4")
			So(len(matches), ShouldEqual, 1)
			So(matches[0], ShouldEqual, "1:2:3::4")
		})
		Convey("And it contains two in a sentence", func() {
			matches := IP("5:6:7::8, and 2001::/32 are cool")
			So(len(matches), ShouldEqual, 2)
			So(matches[0], ShouldEqual, "5:6:7::8")
			So(matches[1], ShouldEqual, "2001::/32")
		})
		Convey("And it contains an obfuscated address", func() {
			matches := IP("9[:]10{:}11(::)12")
			So(len(matches), ShouldEqual, 1)
			So(matches[0], ShouldEqual, "9:10:11::12")
		})
	})
	Convey("When looking for IPv4 addresses in a string", t, func() {
		Convey("And it is one IPv4 address only", func() {
			matches := IP("1.2.3.4")
			So(len(matches), ShouldEqual, 1)
			So(matches[0], ShouldEqual, "1.2.3.4")
		})
		Convey("And it contains two in a sentence", func() {
			matches := IP("5.6.7.8, and 5.6.7.0/24 are cool")
			So(len(matches), ShouldEqual, 2)
			So(matches[0], ShouldEqual, "5.6.7.8")
			So(matches[1], ShouldEqual, "5.6.7.0/24")
		})
		Convey("And it contains an obfuscated address", func() {
			matches := IP("9[.]10{.}11(.)12")
			So(len(matches), ShouldEqual, 1)
			So(matches[0], ShouldEqual, "9.10.11.12")
		})
	})
}
