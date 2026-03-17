package str

import (
	"strings"
	"testing"

	"github.com/ppreeper/assert"
)

var BannerPatterns = []struct {
	msg        string
	spat       string
	pat        string
	epat       string
	padding    int
	header     bool
	footer     bool
	res        string
	resDefault string
}{
	{
		"message", "//", "-", "", 0, true, true,
		strings.Join([]string{"//-------", "//message", "//-------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "//", "-", "//", 0, true, true,
		strings.Join([]string{"//-------//", "//message", "//-------//"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "//", "-", "", 1, true, true,
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "//", "-", "", 2, true, true,
		strings.Join([]string{"//-----------", "//  message", "//-----------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "##", "-", "", 3, true, true,
		strings.Join([]string{"##-------------", "##   message", "##-------------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"NoHeader", "//", "-", "", 2, false, true,
		strings.Join([]string{"//  NoHeader", "//------------"}, "\n"),
		strings.Join([]string{"//----------", "// NoHeader", "//----------"}, "\n"),
	},
	{
		"NoFooter", "//", "-", "", 2, true, false,
		strings.Join([]string{"//------------", "//  NoFooter"}, "\n"),
		strings.Join([]string{"//----------", "// NoFooter", "//----------"}, "\n"),
	},
	{
		"NoHeaderFooter", "//", "-", "", 2, false, false,
		strings.Join([]string{"//  NoHeaderFooter"}, "\n"),
		strings.Join([]string{"//----------------", "// NoHeaderFooter", "//----------------"}, "\n"),
	},
	{
		"message A", "//", "-", "", 0, true, true,
		strings.Join([]string{"//---------", "//message A", "//---------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "//-----------"}, "\n"),
	},
	{
		"message A", "//", "-", "", 1, true, true,
		strings.Join([]string{"//-----------", "// message A", "//-----------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "//-----------"}, "\n"),
	},
	{
		"message A message B", "//", "-", "", 1, true, true,
		strings.Join([]string{"//---------------------", "// message A message B", "//---------------------"}, "\n"),
		strings.Join([]string{"//---------------------", "// message A message B", "//---------------------"}, "\n"),
	},
	{
		"message A" + "\n" + "message B", "//", "-", "", 1, true, true,
		strings.Join([]string{"//-----------", "// message A", "// message B", "//-----------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "// message B", "//-----------"}, "\n"),
	},
	{
		"message A" + "\n" + "message B" + "\n" + "message C", "//", "-", "", 1, true, true,
		strings.Join([]string{"//-----------", "// message A", "// message B", "// message C", "//-----------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "// message B", "// message C", "//-----------"}, "\n"),
	},
	{
		"message A" + "\n" + "message B" + "\n" + "\t" + "message C", "//", "-", "", 1, true, true,
		strings.Join([]string{"//---------------", "// message A", "// message B", "//     message C", "//---------------"}, "\n"),
		strings.Join([]string{"//---------------", "// message A", "// message B", "//     message C", "//---------------"}, "\n"),
	},
}

func TestBannerEmpty(t *testing.T) {
	b := NewBanner()
	assert.Equal(t, b, &Banner{
		SPat:   "//",
		Pat:    "-",
		EPat:   "",
		Pad:    1,
		Header: true,
		Footer: true,
	})
}

func TestBannerWithConfig(t *testing.T) {
	b := NewBannerWithConfig(Banner{
		SPat:   "//",
		Pat:    "-",
		EPat:   "",
		Pad:    1,
		Header: true,
		Footer: true,
	})
	assert.Equal(t, b, &Banner{
		SPat:   "//",
		Pat:    "-",
		EPat:   "",
		Pad:    1,
		Header: true,
		Footer: true,
	})
}

func TestBannerDefault(t *testing.T) {
	for _, tc := range BannerPatterns {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			b := NewBanner()
			res := b.SPrint(tc.msg)
			// fmt.Printf("expected:\n%s\n", tc.resDefault)
			// fmt.Printf("got:\n%s\n", res)
			assert.Equal(t, tc.resDefault, res)
		})
	}
}

func TestBannerMsg(t *testing.T) {
	for _, tc := range BannerPatterns {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			b := NewBanner().StartPattern(tc.spat).Pattern(tc.pat).Padding(tc.padding).EndPattern(tc.epat)
			if !tc.header {
				b.NoHeader()
			}
			if !tc.footer {
				b.NoFooter()
			}
			res := b.SPrint(tc.msg)
			// fmt.Printf("expected:\n%s\n", tc.res)
			// fmt.Printf("got:\n%s\n", res)
			assert.Equal(t, tc.res, res)
		})
	}
}

// BenchmarkBannerDefault benchmarks the default banner generation.
func BenchmarkBannerDefault(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := NewBanner()
			b.SPrint("message")
		}
	})
}
