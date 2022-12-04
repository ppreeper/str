package str

// Left pads string on left side with pattern p, count c times
func Left(s string, p string, c int) string {
	if c <= 0 || len(p) < 1 {
		return s
	}

	retString := make([]byte, len(p)*c+len(s))
	pad := retString[:len(p)*c]
	bp := copy(pad, p)
	for bp < len(pad) {
		copy(pad[bp:], pad[:bp])
		bp *= 2
	}
	copy(retString[len(pad):], s)
	return string(retString)
}

// Right pads string on right side with pattern p, count c times
func Right(s string, p string, c int) string {
	if c <= 0 || len(p) < 1 {
		return s
	}

	retString := make([]byte, len(p)*c+len(s))
	copy(retString, s)
	b := retString[len(s):]
	bp := copy(b, p)
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
	return string(retString)
}

// TruncLeft returns the first character or characters of a string.
func TruncLeft(s string, l int) string {
	if l <= 0 {
		return ""
	}
	if l > len(s) {
		return s
	}
	ret := []byte(s)
	return string(ret[:l])
}

// TruncRight returns the last character or characters of a string.
func TruncRight(s string, l int) string {
	if l <= 0 {
		return ""
	}
	if l > len(s) {
		return s
	}
	ret := []byte(s)
	return string(ret[len(s)-l:])
}

// LeftLen pads string on left side with pattern p, returns string of length l
func LeftLen(s string, p string, l int) string {
	if l <= 0 {
		return ""
	}

	retString := make([]byte, len(p)*l+len(s))
	pad := retString[:len(p)*l]
	bp := copy(pad, p)
	for bp < len(pad) {
		copy(pad[bp:], pad[:bp])
		bp *= 2
	}
	copy(retString[len(pad):], s)
	return string(retString[len(retString)-l:])
}

// RightLen pads string on left side with p, returns string of length l
func RightLen(s string, p string, l int) string {
	if l <= 0 {
		return ""
	}

	retString := make([]byte, len(p)*l+len(s))
	copy(retString, s)
	pad := retString[len(s):]
	bp := copy(pad, p)
	for bp < len(pad) {
		copy(pad[bp:], pad[:bp])
		bp *= 2
	}
	return string(retString[:l])
}

// ZFill fills string with 0 on left c time
func ZFill(s string, c int) string {
	return Left(s, "0", c)
}

// ZFillLen fills string with 0 on left side returns string length of l
func ZFillLen(s string, l int) string {
	return LeftLen(s, "0", l)
}

// LJust returns left justified string with space(s) filler
func LJust(s string, c int) string {
	return Right(s, " ", c)
}

// LJustLen returns left justified string with space(s) filler of length
func LJustLen(s string, l int) string {
	return RightLen(s, " ", l)
}

// RJust returns right justified string with space(s) filler
func RJust(s string, c int) string {
	return Left(s, " ", c)
}

// RJustLen returns right justified string with space(s) filler of length
func RJustLen(s string, l int) string {
	return LeftLen(s, " ", l)
}
