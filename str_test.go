package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// LeftRightPatterns
var LeftRightPatterns = []struct {
	src   string
	pat   string
	count int
	left  string
	right string
}{
	{"D", "", 1, "D", "D"},
	{"D", "a", -1, "D", "D"},
	{"D", "a", 0, "D", "D"},
	{"D", "a", 1, "aD", "Da"},
	{"D", "a", 2, "aaD", "Daa"},
	{"D", "a", 3, "aaaD", "Daaa"},
	{"D", "aa", 2, "aaaaD", "Daaaa"},
	{"ABCD", "a", 0, "ABCD", "ABCD"},
	{"ABCD", "a", 1, "aABCD", "ABCDa"},
	{"ABCD", "a", 2, "aaABCD", "ABCDaa"},
	{"ABCD", "a", 3, "aaaABCD", "ABCDaaa"},
	{"ABCD", "aa", 2, "aaaaABCD", "ABCDaaaa"},
}

// Left
func TestLeft(t *testing.T) {
	for _, tc := range LeftRightPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := Left(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.left, res))
		})
	}
}

// Right
func TestRight(t *testing.T) {
	for _, tc := range LeftRightPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := Right(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.right, res))
		})
	}
}

// TruncLeftRightPatterns
var TruncLeftRightPatterns = []struct {
	src   string
	count int
	left  string
	right string
}{
	{"A", 1, "A", "A"},
	{"A", -1, "", ""},
	{"A", 0, "", ""},
	{"A", 2, "A", "A"},
	{"AB", 1, "A", "B"},
	{"AB", 2, "AB", "AB"},
	{"AB", 3, "AB", "AB"},
	{"ABC", 1, "A", "C"},
	{"ABC", 2, "AB", "BC"},
	{"ABC", 3, "ABC", "ABC"},
	{"ABC", 4, "ABC", "ABC"},
}

// TruncLeft
func TestTruncLeft(t *testing.T) {
	for _, tc := range TruncLeftRightPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncLeft(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.left, res))
		})
	}
}

// TruncRight
func TestTruncRight(t *testing.T) {
	for _, tc := range TruncLeftRightPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncRight(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.right, res))
		})
	}
}

// LeftRightLenPatterns
var LeftRightLenPatterns = []struct {
	src   string
	pat   string
	count int
	left  string
	right string
}{
	{"D", "", 1, "D", "D"},
	{"D", "a", -1, "", ""},
	{"D", "a", 0, "", ""},
	{"D", "a", 1, "D", "D"},
	{"D", "a", 2, "aD", "Da"},
	{"D", "a", 3, "aaD", "Daa"},
	{"D", "a", 4, "aaaD", "Daaa"},
	{"D", "aa", 3, "aaD", "Daa"},
	{"ABCD", "a", 0, "", ""},
	{"ABCD", "a", 1, "D", "A"},
	{"ABCD", "a", 2, "CD", "AB"},
	{"ABCD", "a", 3, "BCD", "ABC"},
	{"ABCD", "a", 4, "ABCD", "ABCD"},
	{"ABCD", "a", 5, "aABCD", "ABCDa"},
	{"ABCD", "aa", 7, "aaaABCD", "ABCDaaa"},
	{"ABCD", "aa", 8, "aaaaABCD", "ABCDaaaa"},
}

// LeftLen
func TestLeftLen(t *testing.T) {
	for _, tc := range LeftRightLenPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LeftLen(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.left, res))
		})
	}
}

// RightLen
func TestRightLen(t *testing.T) {
	for _, tc := range LeftRightLenPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RightLen(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.right, res))
		})
	}
}

// LJustRJustPatterns
var LJustRJustPatterns = []struct {
	src   string
	count int
	left  string
	right string
}{
	{"D", 1, "D ", " D"},
	{"D", -1, "D", "D"},
	{"D", 0, "D", "D"},
	{"D", 1, "D ", " D"},
	{"D", 2, "D  ", "  D"},
	{"D", 3, "D   ", "   D"},
	{"D", 4, "D    ", "    D"},
	{"ABCD", 0, "ABCD", "ABCD"},
	{"ABCD", 1, "ABCD ", " ABCD"},
	{"ABCD", 2, "ABCD  ", "  ABCD"},
	{"ABCD", 3, "ABCD   ", "   ABCD"},
	{"ABCD", 4, "ABCD    ", "    ABCD"},
	{"ABCD", 5, "ABCD     ", "     ABCD"},
}

// LJust
func TestLJust(t *testing.T) {
	for _, tc := range LJustRJustPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJust(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.left, res))
		})
	}
}

// RJust
func TestRJust(t *testing.T) {
	for _, tc := range LJustRJustPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJust(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.right, res))
		})
	}
}

// LJustRJustLenPatterns
var LJustRJustLenPatterns = []struct {
	src   string
	count int
	left  string
	right string
}{
	{"D", 1, "D", "D"},
	{"D", -1, "", ""},
	{"D", 0, "", ""},
	{"D", 1, "D", "D"},
	{"D", 2, "D ", " D"},
	{"D", 3, "D  ", "  D"},
	{"D", 4, "D   ", "   D"},
	{"ABCD", 0, "", ""},
	{"ABCD", 1, "A", "D"},
	{"ABCD", 2, "AB", "CD"},
	{"ABCD", 3, "ABC", "BCD"},
	{"ABCD", 4, "ABCD", "ABCD"},
	{"ABCD", 5, "ABCD ", " ABCD"},
}

// LJustLen
func TestLJustLen(t *testing.T) {
	for _, tc := range LJustRJustLenPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJustLen(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.left, res))
		})
	}
}

// RJustLen
func TestRJustLen(t *testing.T) {
	for _, tc := range LJustRJustLenPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJustLen(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.right, res))
		})
	}
}

// ZFillPatterns
var ZFillPatterns = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "0D"},
	{"D", -1, "D"},
	{"D", 0, "D"},
	{"D", 1, "0D"},
	{"D", 2, "00D"},
	{"D", 3, "000D"},
	{"D", 4, "0000D"},
	{"ABCD", 0, "ABCD"},
	{"ABCD", 1, "0ABCD"},
	{"ABCD", 2, "00ABCD"},
	{"ABCD", 3, "000ABCD"},
	{"ABCD", 4, "0000ABCD"},
	{"ABCD", 5, "00000ABCD"},
}

func TestZFill(t *testing.T) {
	for _, tc := range ZFillPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := ZFill(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// ZFillLenPatterns
var ZFillLenPatterns = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "D"},
	{"D", -1, ""},
	{"D", 0, ""},
	{"D", 1, "D"},
	{"D", 2, "0D"},
	{"D", 3, "00D"},
	{"D", 4, "000D"},
	{"ABCD", -1, ""},
	{"ABCD", 0, ""},
	{"ABCD", 1, "D"},
	{"ABCD", 2, "CD"},
	{"ABCD", 3, "BCD"},
	{"ABCD", 4, "ABCD"},
	{"ABCD", 5, "0ABCD"},
}

func TestZFillLen(t *testing.T) {
	for _, tc := range ZFillLenPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := ZFillLen(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// Benchmarks

// Left

func BenchmarkLeft_D_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Left("D", "→", 4)
	}
}

func BenchmarkLeft_D_a_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Left("D", "a", 4)
	}
}

func BenchmarkLeft_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Left("D", "a", 200)
	}
}

// TruncLeft
func BenchmarkTruncLeft_DDDD_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TruncLeft("DDDD", 2)
	}
}

// LeftLen
func BenchmarkLeftLen_D_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeftLen("D", "→", 4)
	}
}

func BenchmarkLeftLen_D_a_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeftLen("D", "a", 4)
	}
}

func BenchmarkLeftLen_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeftLen("D", "a", 200)
	}
}

// LJust

func BenchmarkLJust_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LJust("→", 4)
	}
}

func BenchmarkLJust_D_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LJust("D", 4)
	}
}

func BenchmarkLJust_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LJust("D", 200)
	}
}

// LJustLen

func BenchmarkLJustLen_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LJustLen("→", 4)
	}
}

func BenchmarkLJustLen_D_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LJustLen("D", 4)
	}
}

func BenchmarkLJustLen_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LJustLen("D", 200)
	}
}

// Right

func BenchmarkRight_D_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Right("D", "→", 4)
	}
}

func BenchmarkRight_D_a_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Right("D", "a", 4)
	}
}

func BenchmarkRight_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Right("D", "a", 200)
	}
}

// TruncRight
func BenchmarkTruncRight_DDDD_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TruncRight("DDDD", 2)
	}
}

// RightLen

func BenchmarkRightLen_D_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RightLen("D", "→", 4)
	}
}

func BenchmarkRightLen_D_a_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RightLen("D", "a", 4)
	}
}

func BenchmarkRightLen_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RightLen("D", "a", 200)
	}
}

// RJust

func BenchmarkRJust_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RJust("→", 4)
	}
}

func BenchmarkRJust_D_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RJust("D", 4)
	}
}

func BenchmarkRJust_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RJust("D", 200)
	}
}

// RJustLen

func BenchmarkRJustLen_utf8_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RJustLen("→", 4)
	}
}

func BenchmarkRJustLen_D_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RJustLen("D", 4)
	}
}

func BenchmarkRJustLen_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RJustLen("D", 200)
	}
}

// ZFill

func BenchmarkZFill_D_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZFill("D", 4)
	}
}

func BenchmarkZFill_DDDD_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZFill("DDDD", 4)
	}
}

func BenchmarkZFill_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZFill("D", 200)
	}
}

// ZFillLen

func BenchmarkZFillLen_D_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZFillLen("D", 4)
	}
}

func BenchmarkZFillLen_DDDD_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZFillLen("DDDD", 4)
	}
}

func BenchmarkZFillLen_D_a_200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZFillLen("D", 200)
	}
}
