package str

import "strings"

// Banner generates text banners suitable for section comment headers.
// It supports configurable start/end patterns, fill patterns, padding,
// and optional header/footer separator lines.
type Banner struct {
	SPat   string
	Pat    string
	EPat   string
	Pad    int
	Header bool
	Footer bool
}

// NewBanner returns a Banner with default configuration:
// start pattern "//", fill pattern "-", no end pattern, padding 1,
// and both header and footer enabled.
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

// NewBannerWithConfig returns a Banner initialized from the provided config.
func NewBannerWithConfig(b Banner) *Banner {
	return &b
}

// StartPattern sets the line-start prefix pattern.
func (b *Banner) StartPattern(pat string) *Banner {
	b.SPat = pat
	return b
}

// Pattern sets the fill/separator pattern used in header and footer lines.
func (b *Banner) Pattern(pat string) *Banner {
	b.Pat = pat
	return b
}

// EndPattern sets the line-end suffix pattern for header and footer lines.
func (b *Banner) EndPattern(pat string) *Banner {
	b.EPat = pat
	return b
}

// Padding sets the number of spaces between the start pattern and message text.
func (b *Banner) Padding(pad int) *Banner {
	b.Pad = pad
	return b
}

// NoHeader disables the header separator line.
func (b *Banner) NoHeader() *Banner {
	b.Header = false
	return b
}

// NoFooter disables the footer separator line.
func (b *Banner) NoFooter() *Banner {
	b.Footer = false
	return b
}

// SPrint generates the complete banner string. It accepts variadic strings
// (joined with a space), splits on newlines, converts tabs to 4 spaces,
// and auto-sizes the separator width to the longest line.
func (b *Banner) SPrint(msg ...string) string {
	lines := strings.Split(strings.Join(msg, " "), "\n")

	// Build prefix once: start pattern + padding spaces.
	prefix := Right(b.SPat, " ", b.Pad)

	// Process lines in place, track max length.
	maxLen := 0
	for i, v := range lines {
		lines[i] = prefix + strings.ReplaceAll(v, "\t", "    ")
		if len(lines[i]) > maxLen {
			maxLen = len(lines[i])
		}
	}
	totalLen := maxLen + b.Pad

	// Compute separator line once for header and footer.
	var sep string
	if b.Header || b.Footer {
		sep = RightLen(b.SPat, b.Pat, totalLen)
	}

	// Pre-size builder to avoid internal growth.
	estCap := 0
	if b.Header {
		estCap += len(sep) + len(b.EPat) + 1
	}
	for _, line := range lines {
		estCap += len(line) + 1
	}
	if b.Footer {
		estCap += len(sep) + len(b.EPat) + 1
	}

	var sb strings.Builder
	sb.Grow(estCap)

	if b.Header {
		sb.WriteString(sep)
		sb.WriteString(b.EPat)
		sb.WriteByte('\n')
	}
	for i, line := range lines {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(line)
	}
	if b.Footer {
		sb.WriteByte('\n')
		sb.WriteString(sep)
		sb.WriteString(b.EPat)
	}
	return sb.String()
}
