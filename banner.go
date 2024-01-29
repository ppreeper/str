package str

import (
	"fmt"
	"slices"
	"strings"
)

type Banner struct {
	SPat   string
	Pat    string
	EPat   string
	Pad    int
	Header bool
	Footer bool
}

func NewBanner() *Banner {
	return &Banner{
		SPat:   "//",
		Pat:    "-",
		EPat:   "",
		Pad:    1,
		Header: true,
		Footer: true,
	}
}

func NewBannerWithConfig(b Banner) *Banner {
	return &b
}

func (b *Banner) StartPattern(pat string) *Banner {
	b.SPat = pat
	return b
}

func (b *Banner) Pattern(pat string) *Banner {
	b.Pat = pat
	return b
}

func (b *Banner) EndPattern(pat string) *Banner {
	b.EPat = pat
	return b
}

func (b *Banner) Padding(pad int) *Banner {
	b.Pad = pad
	return b
}

func (b *Banner) NoHeader() *Banner {
	b.Header = false
	return b
}

func (b *Banner) NoFooter() *Banner {
	b.Footer = false
	return b
}

func (b *Banner) Println(msg ...string) {
	fmt.Printf("%s", b.SPrintln(msg...))
}

func (b *Banner) SPrintln(msg ...string) string {
	msgs := strings.Join(msg, " ")
	msgsLen := []int{}
	for _, m := range strings.Split(msgs, "\n") {
		msgsLen = append(msgsLen, len(m))
	}
	msgs = strings.ReplaceAll(msgs, "\n", "\n"+b.SPat+RJust("", b.Pad))
	totalLen := len(b.SPat) + slices.Max(msgsLen) + b.Pad*2 + len(b.EPat)
	bannerMsg := b.SPat + RJust("", b.Pad) + msgs
	bannerHeader := ""
	if b.Header {
		bannerHeader = RightLen(b.SPat, b.Pat, totalLen) + "\n"
	}
	bannerFooter := ""
	if b.Footer {
		bannerFooter = "\n" + RightLen(b.SPat, b.Pat, totalLen)
	}

	return bannerHeader + bannerMsg + bannerFooter + "\n"
}
