package str

import (
	"strings"
)

// Left pads string on left side with pattern pat, count count times
func Left(src string, pat string, count int) string {
	if count <= 0 || len(pat) < 1 {
		return src
	}
	return strings.Repeat(pat, count) + src
}

// Right pads string on right side with pattern pat, count count times
func Right(src string, pat string, count int) string {
	if count <= 0 || len(pat) < 1 {
		return src
	}
	return src + strings.Repeat(pat, count)
}

// TruncLeft returns the first character or characters of a string.
func TruncLeft(src string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(src) <= length {
		return src
	}
	return src[:length]
}

// TruncRight returns the last character or characters of a string.
func TruncRight(src string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(src) <= length {
		return src
	}
	return src[len(src)-length:]
}

// LeftLen pads string on left side with pattern, returns right side string of length
func LeftLen(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	ret := strings.Repeat(pat, length) + src
	return ret[len(ret)-length:]
}

// RightLen pads string on left side with p, returns string of length l
func RightLen(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	ret := src + strings.Repeat(pat, length)
	return ret[:length]
}

// LJust returns left justified string with space(s) filler
func LJust(src string, count int) string {
	if count <= 0 {
		return src
	}
	return src + strings.Repeat(" ", count)
}

// LJustLen returns left justified string with space(s) filler of length
func LJustLen(src string, length int) string {
	if length <= 0 {
		return ""
	}
	ret := src + strings.Repeat(" ", length)
	if len(ret) <= length {
		return src
	}
	return ret[:length]
}

// RJust returns right justified string with space(s) filler
func RJust(src string, count int) string {
	if count <= 0 {
		return src
	}
	return strings.Repeat(" ", count) + src
}

// RJustLen returns right justified string with space(s) filler of length
func RJustLen(src string, length int) string {
	if length <= 0 {
		return ""
	}
	ret := strings.Repeat(" ", length) + src
	if len(ret) <= length {
		return src
	}
	return ret[len(ret)-length:]
}

// ZFill fills string with 0 on left c time
func ZFill(src string, count int) string {
	if count <= 0 {
		return src
	}
	return strings.Repeat("0", count) + src
}

// ZFillLen fills string with 0 on left side returns string length of l
func ZFillLen(src string, length int) string {
	if length <= 0 {
		return ""
	}
	ret := strings.Repeat("0", length) + src
	if len(ret) <= length {
		return src
	}
	return ret[len(ret)-length:]
}
