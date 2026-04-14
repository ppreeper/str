package str

import "unsafe"

// repeatFill fills buf[0:n] with s repeated, using copy-doubling.
// n must equal len(s) * k for some integer k >= 1.
func repeatFill(buf []byte, s string, n int) {
	bp := copy(buf, s)
	for bp < n {
		bp += copy(buf[bp:n], buf[:bp])
	}
}

// byteFill fills buf[0:n] with byte b, using copy-doubling.
func byteFill(buf []byte, b byte, n int) {
	buf[0] = b
	bp := 1
	for bp < n {
		bp += copy(buf[bp:n], buf[:bp])
	}
}

// Left prepends pat repeated count times to the left of src.
func Left(src string, pat string, count int) string {
	if count <= 0 || len(pat) < 1 {
		return src
	}
	padLen := len(pat) * count
	total := padLen + len(src)
	buf := make([]byte, total)
	repeatFill(buf[:padLen], pat, padLen)
	copy(buf[padLen:], src)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// Right appends pat repeated count times to the right of src.
func Right(src string, pat string, count int) string {
	if count <= 0 || len(pat) < 1 {
		return src
	}
	padLen := len(pat) * count
	total := len(src) + padLen
	buf := make([]byte, total)
	copy(buf, src)
	repeatFill(buf[len(src):], pat, padLen)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// TruncLeft returns the first length bytes of src.
// Length is measured in bytes, not runes; callers working with multi-byte
// characters should ensure length falls on a rune boundary.
func TruncLeft(src string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(src) <= length {
		return src
	}
	return src[:length]
}

// TruncRight returns the last length bytes of src.
// Length is measured in bytes, not runes; callers working with multi-byte
// characters should ensure length falls on a rune boundary.
func TruncRight(src string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(src) <= length {
		return src
	}
	return src[len(src)-length:]
}

// LeftLen pads src on the left with pat, returning a string of exactly length bytes.
// The padding pattern is right-aligned: if the needed padding is not a multiple of
// len(pat), the leftmost bytes contain a partial tail of pat.
// Length is measured in bytes, not runes; callers working with multi-byte
// characters should ensure length falls on a rune boundary.
func LeftLen(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	needed := length - len(src)
	if needed <= 0 {
		return src[len(src)-length:]
	}
	buf := make([]byte, length)
	fullReps := needed / len(pat)
	remainder := needed % len(pat)
	off := 0
	if remainder > 0 {
		off += copy(buf, pat[len(pat)-remainder:])
	}
	if fullReps > 0 {
		repeatFill(buf[off:needed], pat, fullReps*len(pat))
	}
	copy(buf[needed:], src)
	return unsafe.String(unsafe.SliceData(buf), length)
}

// RightLen pads src on the right with pat, returning a string of exactly length bytes.
// The padding pattern is left-aligned: if the needed padding is not a multiple of
// len(pat), the rightmost bytes contain a partial head of pat.
// Length is measured in bytes, not runes; callers working with multi-byte
// characters should ensure length falls on a rune boundary.
func RightLen(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	if len(pat) < 1 {
		return src
	}
	needed := length - len(src)
	if needed <= 0 {
		return src[:length]
	}
	buf := make([]byte, length)
	copy(buf, src)
	fullReps := needed / len(pat)
	remainder := needed % len(pat)
	fullLen := fullReps * len(pat)
	if fullReps > 0 {
		repeatFill(buf[len(src):len(src)+fullLen], pat, fullLen)
	}
	if remainder > 0 {
		copy(buf[len(src)+fullLen:], pat[:remainder])
	}
	return unsafe.String(unsafe.SliceData(buf), length)
}

// LJust returns src left-justified by appending count spaces to the right.
func LJust(src string, count int) string {
	if count <= 0 {
		return src
	}
	total := len(src) + count
	buf := make([]byte, total)
	copy(buf, src)
	byteFill(buf[len(src):], ' ', count)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// LJustLen returns src left-justified with trailing spaces, exactly length bytes.
// Length is measured in bytes, not runes; callers working with multi-byte
// characters should ensure length falls on a rune boundary.
func LJustLen(src string, length int) string {
	if length <= 0 {
		return ""
	}
	needed := length - len(src)
	if needed <= 0 {
		return src[:length]
	}
	buf := make([]byte, length)
	copy(buf, src)
	byteFill(buf[len(src):], ' ', needed)
	return unsafe.String(unsafe.SliceData(buf), length)
}

// RJust returns src right-justified by prepending count spaces to the left.
func RJust(src string, count int) string {
	if count <= 0 {
		return src
	}
	total := count + len(src)
	buf := make([]byte, total)
	byteFill(buf[:count], ' ', count)
	copy(buf[count:], src)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// RJustLen returns src right-justified with leading spaces, exactly length bytes.
// Length is measured in bytes, not runes; callers working with multi-byte
// characters should ensure length falls on a rune boundary.
func RJustLen(src string, length int) string {
	if length <= 0 {
		return ""
	}
	needed := length - len(src)
	if needed <= 0 {
		return src[len(src)-length:]
	}
	buf := make([]byte, length)
	byteFill(buf[:needed], ' ', needed)
	copy(buf[needed:], src)
	return unsafe.String(unsafe.SliceData(buf), length)
}

// ZFill prepends count zeros to the left of src.
func ZFill(src string, count int) string {
	if count <= 0 {
		return src
	}
	total := count + len(src)
	buf := make([]byte, total)
	byteFill(buf[:count], '0', count)
	copy(buf[count:], src)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// ZFillLen prepends zeros to src on the left, returning a string of exactly length bytes.
// Length is measured in bytes, not runes.
func ZFillLen(src string, length int) string {
	if length <= 0 {
		return ""
	}
	needed := length - len(src)
	if needed <= 0 {
		return src[len(src)-length:]
	}
	buf := make([]byte, length)
	byteFill(buf[:needed], '0', needed)
	copy(buf[needed:], src)
	return unsafe.String(unsafe.SliceData(buf), length)
}
