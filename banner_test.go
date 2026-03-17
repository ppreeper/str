package str

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ppreeper/assert"
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
	{
		"message", "//", "-", 0, true, true,
		strings.Join([]string{"//-------", "//message", "//-------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "//", "-", 1, true, true,
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "//", "-", 2, true, true,
		strings.Join([]string{"//-----------", "//  message", "//-----------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"message", "##", "-", 3, true, true,
		strings.Join([]string{"##-------------", "##   message", "##-------------"}, "\n"),
		strings.Join([]string{"//---------", "// message", "//---------"}, "\n"),
	},
	{
		"NoHeader", "//", "-", 2, false, true,
		strings.Join([]string{"//  NoHeader", "//------------"}, "\n"),
		strings.Join([]string{"//----------", "// NoHeader", "//----------"}, "\n"),
	},
	{
		"NoFooter", "//", "-", 2, true, false,
		strings.Join([]string{"//------------", "//  NoFooter"}, "\n"),
		strings.Join([]string{"//----------", "// NoFooter", "//----------"}, "\n"),
	},
	{
		"NoHeaderFooter", "//", "-", 2, false, false,
		strings.Join([]string{"//  NoHeaderFooter"}, "\n"),
		strings.Join([]string{"//----------------", "// NoHeaderFooter", "//----------------"}, "\n"),
	},
	{
		"message A", "//", "-", 0, true, true,
		strings.Join([]string{"//---------", "//message A", "//---------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "//-----------"}, "\n"),
	},
	{
		"message A", "//", "-", 1, true, true,
		strings.Join([]string{"//-----------", "// message A", "//-----------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "//-----------"}, "\n"),
	},
	{
		"message A message B", "//", "-", 1, true, true,
		strings.Join([]string{"//---------------------", "// message A message B", "//---------------------"}, "\n"),
		strings.Join([]string{"//---------------------", "// message A message B", "//---------------------"}, "\n"),
	},
	{
		"message A" + "\n" + "message B", "//", "-", 1, true, true,
		strings.Join([]string{"//-----------", "// message A", "// message B", "//-----------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "// message B", "//-----------"}, "\n"),
	},
	{
		"message A" + "\n" + "message B" + "\n" + "message C", "//", "-", 1, true, true,
		strings.Join([]string{"//-----------", "// message A", "// message B", "// message C", "//-----------"}, "\n"),
		strings.Join([]string{"//-----------", "// message A", "// message B", "// message C", "//-----------"}, "\n"),
	},
	{
		"message A" + "\n" + "message B" + "\n" + "\t" + "message C", "//", "-", 1, true, true,
		strings.Join([]string{"//---------------", "// message A", "// message B", "//     message C", "//---------------"}, "\n"),
		strings.Join([]string{"//---------------", "// message A", "// message B", "//     message C", "//---------------"}, "\n"),
	},
}

func TestBannerDefault(t *testing.T) {
	for _, tc := range BannerPatterns {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			b := NewBanner()
			res := b.SPrint(tc.msg)
			assert.Equal(t, tc.resDefault, res)
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
			fmt.Printf("Expected\n%s\n", tc.res)
			res := b.SPrint(tc.msg)
			fmt.Printf("Got:\n%s\n", res)

			assert.Equal(t, tc.res, res)
		})
	}
}
