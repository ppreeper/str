package str

import (
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

func (b *Banner) SPrint(msg ...string) string {
	// Join the message lines into a single string
	msgs := []string{}
	msgsLen := []int{}
	for _, v := range strings.Split(strings.Join(msg, " "), "\n") {
		m := b.SPat + RJust("", b.Pad) + strings.ReplaceAll(v, "\t", "    ")
		msgsLen = append(msgsLen, len(m))
		msgs = append(msgs, m)
	}
	// Calculate the total length of the banner)
	totalLen := slices.Max(msgsLen) + b.Pad

	banner := []string{}
	// Append the header pattern to the banner if the header is enabled
	if b.Header {
		banner = append(banner, RightLen(b.SPat, b.Pat, totalLen))
	}

	// Append the banner message to the banner
	banner = append(banner, strings.Join(msgs, "\n"))

	// Append the footer pattern to the banner if the footer is enabled
	if b.Footer {
		banner = append(banner, RightLen(b.SPat, b.Pat, totalLen))
	}

	return strings.Join(banner, "\n")
}
