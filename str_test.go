package str

import (
	"runtime"
	"testing"

	"github.com/ppreeper/assert"
)

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

func TestLeft(t *testing.T) {
	for _, tc := range LeftRightPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := Left(tc.src, tc.pat, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestRight(t *testing.T) {
	for _, tc := range LeftRightPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := Right(tc.src, tc.pat, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

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

func TestTruncLeft(t *testing.T) {
	for _, tc := range TruncLeftRightPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncLeft(tc.src, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestTruncRight(t *testing.T) {
	for _, tc := range TruncLeftRightPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncRight(tc.src, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

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

func TestLeftLen(t *testing.T) {
	for _, tc := range LeftRightLenPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LeftLen(tc.src, tc.pat, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestRightLen(t *testing.T) {
	for _, tc := range LeftRightLenPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RightLen(tc.src, tc.pat, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

var LJustRJustPatterns = []struct {
	src   string
	count int
	left  string
	right string
}{
	{"D", 1, "D ", " D"},
	{"D", -1, "D", "D"},
	{"D", 0, "D", "D"},
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

func TestLJust(t *testing.T) {
	for _, tc := range LJustRJustPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJust(tc.src, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestRJust(t *testing.T) {
	for _, tc := range LJustRJustPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJust(tc.src, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

var LJustRJustLenPatterns = []struct {
	src   string
	count int
	left  string
	right string
}{
	{"D", 1, "D", "D"},
	{"D", -1, "", ""},
	{"D", 0, "", ""},
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

func TestLJustLen(t *testing.T) {
	for _, tc := range LJustRJustLenPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJustLen(tc.src, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestRJustLen(t *testing.T) {
	for _, tc := range LJustRJustLenPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJustLen(tc.src, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

var ZFillPatterns = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "0D"},
	{"D", -1, "D"},
	{"D", 0, "D"},
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
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := ZFill(tc.src, tc.count)
			assert.Equal(t, tc.expected, res)
		})
	}
}

var ZFillLenPatterns = []struct {
	src      string
	count    int
	expected string
}{
	{"D", 1, "D"},
	{"D", -1, ""},
	{"D", 0, ""},
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
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := ZFillLen(tc.src, tc.count)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func BenchmarkLeft_D_a_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Left("D", "a", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeft_D_aa_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Left("D", "aa", 2)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeft_D_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Left("D", "→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeft_D_utf8_2_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Left("D", "→→", 2)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeft_D_a_200_1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Left("D", "a", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeft_D_a_100_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Left("D", "aa", 100)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRight_D_a_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Right("D", "a", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRight_D_aa_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Right("D", "aa", 2)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRight_D_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Right("D", "→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRight_D_utf8_2_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Right("D", "→→", 2)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRight_D_a_200_1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Right("D", "a", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRight_D_a_100_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = Right("D", "aa", 100)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkTruncLeft_ABCD_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = TruncLeft("ABCD", 2)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkTruncRight_ABCD_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = TruncRight("ABCD", 2)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeftLen_D_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LeftLen("D", "→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeftLen_D_a_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LeftLen("D", "a", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeftLen_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LeftLen("D", "a", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRightLen_D_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RightLen("D", "→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRightLen_D_a_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RightLen("D", "a", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRightLen_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RightLen("D", "a", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJust_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJust("→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJust_D_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJust("D", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJust_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJust("D", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJust_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJust("→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJust_D_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJust("D", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJust_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJust("D", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJustLen_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJustLen("→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJustLen_D_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJustLen("D", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJustLen_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJustLen("D", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJustLen_utf8_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJustLen("→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJustLen_D_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJustLen("D", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJustLen_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJustLen("D", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFill_D_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFill("D", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFill_ABCD_8(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFill("DDDD", 8)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFill_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFill("D", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFillLen_D_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFillLen("D", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFillLen_ABCD_8(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFillLen("DDDD", 8)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFillLen_D_a_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFillLen("D", 200)
		}
		runtime.KeepAlive(r)
	})
}
