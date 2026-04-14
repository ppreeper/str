package str

import (
	"unicode/utf8"
	"unsafe"
)

// TruncLeftRunes returns the first length runes of src.
// If src contains fewer runes than length, src is returned unchanged.
func TruncLeftRunes(src string, length int) string {
	if length <= 0 {
		return ""
	}
	count := 0
	for i := range src {
		count++
		if count > length {
			return src[:i]
		}
	}
	return src
}

// TruncRightRunes returns the last length runes of src.
// If src contains fewer runes than length, src is returned unchanged.
func TruncRightRunes(src string, length int) string {
	if length <= 0 {
		return ""
	}
	off := len(src)
	for i := 0; i < length && off > 0; i++ {
		_, size := utf8.DecodeLastRuneInString(src[:off])
		off -= size
	}
	return src[off:]
}

// LeftLenRunes pads src on the left with pat, returning a string of exactly length runes.
// The padding pattern is right-aligned: if the needed padding is not a multiple of
// the rune count of pat, the leftmost runes contain a partial tail of pat.
func LeftLenRunes(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	patRunes := utf8.RuneCountInString(pat)
	if patRunes < 1 {
		return src
	}
	srcRunes := utf8.RuneCountInString(src)
	needed := length - srcRunes
	if needed <= 0 {
		return TruncRightRunes(src, length)
	}
	fullReps := needed / patRunes
	remainder := needed % patRunes

	var partialPat string
	partialBytes := 0
	if remainder > 0 {
		partialPat = TruncRightRunes(pat, remainder)
		partialBytes = len(partialPat)
	}
	fullBytes := fullReps * len(pat)
	total := partialBytes + fullBytes + len(src)

	buf := make([]byte, total)
	off := 0
	if remainder > 0 {
		off += copy(buf, partialPat)
	}
	if fullReps > 0 {
		repeatFill(buf[off:off+fullBytes], pat, fullBytes)
		off += fullBytes
	}
	copy(buf[off:], src)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// RightLenRunes pads src on the right with pat, returning a string of exactly length runes.
// The padding pattern is left-aligned: if the needed padding is not a multiple of
// the rune count of pat, the rightmost runes contain a partial head of pat.
func RightLenRunes(src string, pat string, length int) string {
	if length <= 0 {
		return ""
	}
	patRunes := utf8.RuneCountInString(pat)
	if patRunes < 1 {
		return src
	}
	srcRunes := utf8.RuneCountInString(src)
	needed := length - srcRunes
	if needed <= 0 {
		return TruncLeftRunes(src, length)
	}
	fullReps := needed / patRunes
	remainder := needed % patRunes

	var partialPat string
	partialBytes := 0
	if remainder > 0 {
		partialPat = TruncLeftRunes(pat, remainder)
		partialBytes = len(partialPat)
	}
	fullBytes := fullReps * len(pat)
	total := len(src) + fullBytes + partialBytes

	buf := make([]byte, total)
	off := copy(buf, src)
	if fullReps > 0 {
		repeatFill(buf[off:off+fullBytes], pat, fullBytes)
		off += fullBytes
	}
	if remainder > 0 {
		copy(buf[off:], partialPat)
	}
	return unsafe.String(unsafe.SliceData(buf), total)
}

// LJustLenRunes returns src left-justified with trailing spaces, exactly length runes.
func LJustLenRunes(src string, length int) string {
	if length <= 0 {
		return ""
	}
	srcRunes := utf8.RuneCountInString(src)
	needed := length - srcRunes
	if needed <= 0 {
		return TruncLeftRunes(src, length)
	}
	total := len(src) + needed
	buf := make([]byte, total)
	copy(buf, src)
	byteFill(buf[len(src):], ' ', needed)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// RJustLenRunes returns src right-justified with leading spaces, exactly length runes.
func RJustLenRunes(src string, length int) string {
	if length <= 0 {
		return ""
	}
	srcRunes := utf8.RuneCountInString(src)
	needed := length - srcRunes
	if needed <= 0 {
		return TruncRightRunes(src, length)
	}
	total := needed + len(src)
	buf := make([]byte, total)
	byteFill(buf[:needed], ' ', needed)
	copy(buf[needed:], src)
	return unsafe.String(unsafe.SliceData(buf), total)
}

// ZFillLenRunes prepends zeros to src, returning a string of exactly length runes.
func ZFillLenRunes(src string, length int) string {
	if length <= 0 {
		return ""
	}
	srcRunes := utf8.RuneCountInString(src)
	needed := length - srcRunes
	if needed <= 0 {
		return TruncRightRunes(src, length)
	}
	total := needed + len(src)
	buf := make([]byte, total)
	byteFill(buf[:needed], '0', needed)
	copy(buf[needed:], src)
	return unsafe.String(unsafe.SliceData(buf), total)
}
