package str

import (
	"strings"
)

// Left pads string on left side with pattern pat, count count times
func Left(src string, pat string, count int) string {
	if count <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	return strings.Repeat(pat, count) + src
}

// Right pads string on right side with pattern pat, count count times
func Right(src string, pat string, count int) string {
	if count <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	return src + strings.Repeat(pat, count)
}

// TruncLeft returns the first character or characters of a string.
func TruncLeft(src string, length int) string {
	if length <= 0 {
		return ""
	}
	if length > len(src) {
		return src
	}
	ret := []rune(src)
	return string(ret[:length])
}

// TruncRight returns the last character or characters of a string.
func TruncRight(src string, length int) string {
	if length <= 0 {
		return ""
	}
	if length > len(src) {
		return src
	}
	ret := []rune(src)
	return string(ret[len(src)-length:])
}

// LeftLen pads string on left side with pattern p, returns string of length l
func LeftLen(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	return string([]rune(src + strings.Repeat(pat, length))[:length])
}

// RightLen pads string on left side with p, returns string of length l
func RightLen(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	return string([]rune(src + strings.Repeat(pat, length))[:length])
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
	if length > len(ret) {
		return src
	}
	return string([]rune(ret)[len(ret)-length:])
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
	if length > len(ret) {
		return src
	}
	return string([]rune(ret)[:length])
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
	if length > len(ret) {
		return src
	}
	return string([]rune(ret)[len(ret)-length:])
}
