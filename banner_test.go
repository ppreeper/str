package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var BannerPatterns = []struct {
	msg        string
	spat       string
	pat        string
	padding    int
	header     bool
	footer     bool
	res        string
	resDefault string
}{
	{"message", "//", "-", 0, true, true, "//-------\n//message\n//-------\n", "//---------\n// message\n//---------\n"},
	{"message", "//", "-", 1, true, true, "//---------\n// message\n//---------\n", "//---------\n// message\n//---------\n"},
	{"message", "//", "-", 2, true, true, "//-----------\n//  message\n//-----------\n", "//---------\n// message\n//---------\n"},
	{"message", "##", "-", 3, true, true, "##-------------\n##   message\n##-------------\n", "//---------\n// message\n//---------\n"},
	{"NoHeader", "//", "-", 2, false, true, "//  NoHeader\n//------------\n", "//----------\n// NoHeader\n//----------\n"},
	{"NoFooter", "//", "-", 2, true, false, "//------------\n//  NoFooter\n", "//----------\n// NoFooter\n//----------\n"},
	{"NoHeaderFooter", "//", "-", 2, false, false, "//  NoHeaderFooter\n", "//----------------\n// NoHeaderFooter\n//----------------\n"},
}

func TestBannerDefault(t *testing.T) {
	for _, tc := range BannerPatterns {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			b := NewBanner()
			res := b.SPrintln(tc.msg)
			// fmt.Println(res)
			// fmt.Println(tc.resDefault)
			if !assert.True(t, assert.ObjectsAreEqualValues(tc.resDefault, res)) {
				t.Logf("expected: %s", tc.resDefault)
				t.Logf("actual: %s", res)
			}
		})
	}
}

func TestBannerMsg(t *testing.T) {
	for _, tc := range BannerPatterns {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			b := NewBanner().StartPattern(tc.spat).Pattern(tc.pat).Padding(tc.padding)
			if !tc.header {
				b.NoHeader()
			}
			if !tc.footer {
				b.NoFooter()
			}
			res := b.SPrintln(tc.msg)
			// fmt.Println(res)
			// fmt.Println(tc.res)
			if !assert.True(t, assert.ObjectsAreEqualValues(tc.res, res)) {
				t.Logf("expected: %s", tc.res)
				t.Logf("actual: %s", res)
			}
		})
	}
}

var BannerMultiLinePatterns = []struct {
	msg string
	res string
}{
	{
		"message",
		"//---------" + "\n" +
			"// message" + "\n" +
			"//---------" + "\n",
	},
	{
		"message A message B",
		"//---------------------" + "\n" +
			"// message A message B" + "\n" +
			"//---------------------" + "\n",
	},
	{
		"message A" + "\n" + "message B",
		"//-----------" + "\n" +
			"// message A" + "\n" +
			"// message B" + "\n" +
			"//-----------" + "\n",
	},
	{
		"message A" + "\n" + "message B" + "\n" + "message C",
		"//-----------" + "\n" +
			"// message A" + "\n" +
			"// message B" + "\n" +
			"// message C" + "\n" +
			"//-----------" + "\n",
	},
	{
		"message A" + "\n" + "message B" + "\n" + "\t" + "message C",
		"//------------" + "\n" +
			"// message A" + "\n" +
			"// message B" + "\n" +
			"// " + "\t" + "message C" + "\n" +
			"//------------" + "\n",
	},
}

func TestBannerMultiLine(t *testing.T) {
	b := NewBanner()
	for _, tc := range BannerMultiLinePatterns {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			res := b.SPrintln(tc.msg)
			if !assert.True(t, assert.ObjectsAreEqualValues(tc.res, res)) {
				t.Logf("expected: %s", tc.res)
				t.Logf("actual: %s", res)
			}
		})
	}
}
