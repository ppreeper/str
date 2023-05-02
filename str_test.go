package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Left

var LeftPatternTests = []struct {
	src      string
	pat      string
	count    int
	expected string
}{
	{"D", "", 1, "D"},
	{"D", "a", -1, "D"},
	{"D", "a", 0, "D"},
	{"D", "a", 1, "aD"},
	{"D", "a", 2, "aaD"},
	{"D", "a", 3, "aaaD"},
	{"D", "a", 4, "aaaaD"},
	{"ABCD", "a", 0, "ABCD"},
	{"ABCD", "a", 1, "aABCD"},
	{"ABCD", "a", 2, "aaABCD"},
	{"ABCD", "a", 3, "aaaABCD"},
	{"ABCD", "a", 4, "aaaaABCD"},
	{"ABCD", "a", 5, "aaaaaABCD"},
}

func TestLeft(t *testing.T) {
	for _, tc := range LeftPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := Left(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// TruncLeft
var TruncLeftPatterns = []struct {
	src      string
	count    int
	expected string
}{
	{"A", 1, "A"},
	{"A", -1, ""},
	{"A", 0, ""},
	{"A", 2, "A"},
	{"AB", 1, "A"},
	{"AB", 2, "AB"},
	{"AB", 3, "AB"},
	{"ABC", 1, "A"},
	{"ABC", 2, "AB"},
	{"ABC", 3, "ABC"},
	{"ABC", 4, "ABC"},
}

func TestTruncLeft(t *testing.T) {
	for _, tc := range TruncLeftPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncLeft(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// LeftLen

var LeftLenPatternTests = []struct {
	src      string
	pat      string
	count    int
	expected string
}{
	{"D", "", 1, "D"},
	{"D", "a", -1, ""},
	{"D", "a", 0, ""},
	{"D", "a", 1, "D"},
	{"D", "a", 2, "aD"},
	{"D", "a", 3, "aaD"},
	{"D", "a", 4, "aaaD"},
	{"ABCD", "a", 0, ""},
	{"ABCD", "a", 1, "D"},
	{"ABCD", "a", 2, "CD"},
	{"ABCD", "a", 3, "BCD"},
	{"ABCD", "a", 4, "ABCD"},
	{"ABCD", "a", 5, "aABCD"},
}

func TestLeftLen(t *testing.T) {
	for _, tc := range LeftLenPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LeftLen(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// LJust

var LJustPatternTests = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "D "},
	{"D", -1, "D"},
	{"D", 0, "D"},
	{"D", 1, "D "},
	{"D", 2, "D  "},
	{"D", 3, "D   "},
	{"D", 4, "D    "},
	{"ABCD", 0, "ABCD"},
	{"ABCD", 1, "ABCD "},
	{"ABCD", 2, "ABCD  "},
	{"ABCD", 3, "ABCD   "},
	{"ABCD", 4, "ABCD    "},
	{"ABCD", 5, "ABCD     "},
}

func TestLJust(t *testing.T) {
	for _, tc := range LJustPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJust(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// LJustLen

var LJustLenPatternTests = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "D"},
	{"D", -1, ""},
	{"D", 0, ""},
	{"D", 1, "D"},
	{"D", 2, "D "},
	{"D", 3, "D  "},
	{"D", 4, "D   "},
	{"ABCD", 0, ""},
	{"ABCD", 1, "A"},
	{"ABCD", 2, "AB"},
	{"ABCD", 3, "ABC"},
	{"ABCD", 4, "ABCD"},
	{"ABCD", 5, "ABCD "},
}

func TestLJustLen(t *testing.T) {
	for _, tc := range LJustLenPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJustLen(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// Right

var RightPatternTests = []struct {
	src      string
	pat      string
	count    int
	expected string
}{
	{"D", "", 1, "D"},
	{"D", "a", -1, "D"},
	{"D", "a", 0, "D"},
	{"D", "a", 1, "Da"},
	{"D", "a", 2, "Daa"},
	{"D", "a", 3, "Daaa"},
	{"D", "a", 4, "Daaaa"},
	{"ABCD", "a", 0, "ABCD"},
	{"ABCD", "a", 1, "ABCDa"},
	{"ABCD", "a", 2, "ABCDaa"},
	{"ABCD", "a", 3, "ABCDaaa"},
	{"ABCD", "a", 4, "ABCDaaaa"},
	{"ABCD", "a", 5, "ABCDaaaaa"},
}

func TestRight(t *testing.T) {
	for _, tc := range RightPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := Right(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// TruncRight
var TruncRightPatterns = []struct {
	src      string
	count    int
	expected string
}{
	{"A", 1, "A"},
	{"A", -1, ""},
	{"A", 0, ""},
	{"A", 2, "A"},
	{"AB", 1, "B"},
	{"AB", 2, "AB"},
	{"AB", 3, "AB"},
	{"ABC", 1, "C"},
	{"ABC", 2, "BC"},
	{"ABC", 3, "ABC"},
	{"ABC", 4, "ABC"},
}

func TestTruncRight(t *testing.T) {
	for _, tc := range TruncRightPatterns {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncRight(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// RightLen

var RightLenPatternTests = []struct {
	src      string
	pat      string
	count    int
	expected string
}{
	{"D", "", 1, "D"},
	{"D", "a", -1, ""},
	{"D", "a", 0, ""},
	{"D", "a", 1, "D"},
	{"D", "a", 2, "Da"},
	{"D", "a", 3, "Daa"},
	{"D", "a", 4, "Daaa"},
	{"ABCD", "a", -1, ""},
	{"ABCD", "a", 0, ""},
	{"ABCD", "a", 1, "A"},
	{"ABCD", "a", 2, "AB"},
	{"ABCD", "a", 3, "ABC"},
	{"ABCD", "a", 4, "ABCD"},
	{"ABCD", "a", 5, "ABCDa"},
}

func TestRightLen(t *testing.T) {
	for _, tc := range RightLenPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RightLen(tc.src, tc.pat, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// RJust

var RJustPatternTests = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, " D"},
	{"D", -1, "D"},
	{"D", 0, "D"},
	{"D", 1, " D"},
	{"D", 2, "  D"},
	{"D", 3, "   D"},
	{"D", 4, "    D"},
	{"ABCD", 0, "ABCD"},
	{"ABCD", 1, " ABCD"},
	{"ABCD", 2, "  ABCD"},
	{"ABCD", 3, "   ABCD"},
	{"ABCD", 4, "    ABCD"},
	{"ABCD", 5, "     ABCD"},
}

func TestRJust(t *testing.T) {
	for _, tc := range RJustPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJust(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// RJustLen

var RJustLenPatternTests = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "D"},
	{"D", -1, ""},
	{"D", 0, ""},
	{"D", 1, "D"},
	{"D", 2, " D"},
	{"D", 3, "  D"},
	{"D", 4, "   D"},
	{"ABCD", 0, ""},
	{"ABCD", 1, "D"},
	{"ABCD", 2, "CD"},
	{"ABCD", 3, "BCD"},
	{"ABCD", 4, "ABCD"},
	{"ABCD", 5, " ABCD"},
}

func TestRJustLen(t *testing.T) {
	for _, tc := range RJustLenPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJustLen(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// ZFill

var ZFillPatternTests = []struct {
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
	for _, tc := range ZFillPatternTests {
		tc := tc
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := ZFill(tc.src, tc.count)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

// ZFillLen

var ZFillLenPatternTests = []struct {
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
	for _, tc := range ZFillLenPatternTests {
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
