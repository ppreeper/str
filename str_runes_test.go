package str

import (
	"runtime"
	"testing"

	"github.com/ppreeper/assert"
)

var TruncLeftRightRunesPatterns = []struct {
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
	{"→", 1, "→", "→"},
	{"→", 0, "", ""},
	{"→→", 1, "→", "→"},
	{"→→→", 2, "→→", "→→"},
	{"A→B", 1, "A", "B"},
	{"A→B", 2, "A→", "→B"},
	{"A→B", 3, "A→B", "A→B"},
	{"A→B", 4, "A→B", "A→B"},
	{"日本語", 1, "日", "語"},
	{"日本語", 2, "日本", "本語"},
	{"日本語", 3, "日本語", "日本語"},
}

func TestTruncLeftRunes(t *testing.T) {
	for _, tc := range TruncLeftRightRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncLeftRunes(tc.src, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestTruncRightRunes(t *testing.T) {
	for _, tc := range TruncLeftRightRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := TruncRightRunes(tc.src, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

var LeftRightLenRunesPatterns = []struct {
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
	{"D", "→", 1, "D", "D"},
	{"D", "→", 2, "→D", "D→"},
	{"D", "→", 3, "→→D", "D→→"},
	{"D", "→←", 3, "→←D", "D→←"},
	{"D", "→←", 4, "←→←D", "D→←→"},
	{"→→", "a", 4, "aa→→", "→→aa"},
	{"→→→", "a", 2, "→→", "→→"},
	{"日本語", "a", 2, "本語", "日本"},
	{"D", "日本", 5, "日本日本D", "D日本日本"},
	{"D", "日本", 4, "本日本D", "D日本日"},
	{"ABCD", "→", 2, "CD", "AB"},
	{"ABCD", "→", 4, "ABCD", "ABCD"},
	{"ABCD", "→", 5, "→ABCD", "ABCD→"},
	{"ABCD", "→←", 7, "←→←ABCD", "ABCD→←→"},
}

func TestLeftLenRunes(t *testing.T) {
	for _, tc := range LeftRightLenRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LeftLenRunes(tc.src, tc.pat, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestRightLenRunes(t *testing.T) {
	for _, tc := range LeftRightLenRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RightLenRunes(tc.src, tc.pat, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

var LJustRJustLenRunesPatterns = []struct {
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
	{"→", 1, "→", "→"},
	{"→", 2, "→ ", " →"},
	{"→", 3, "→  ", "  →"},
	{"→→", 1, "→", "→"},
	{"→→→", 2, "→→", "→→"},
	{"日本語", 2, "日本", "本語"},
	{"日本語", 5, "日本語  ", "  日本語"},
	{"ABCD", 4, "ABCD", "ABCD"},
	{"ABCD", 6, "ABCD  ", "  ABCD"},
}

func TestLJustLenRunes(t *testing.T) {
	for _, tc := range LJustRJustLenRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := LJustLenRunes(tc.src, tc.count)
			assert.Equal(t, tc.left, res)
		})
	}
}

func TestRJustLenRunes(t *testing.T) {
	for _, tc := range LJustRJustLenRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := RJustLenRunes(tc.src, tc.count)
			assert.Equal(t, tc.right, res)
		})
	}
}

var ZFillLenRunesPatterns = []struct {
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
	{"→", 1, "→"},
	{"→", 2, "0→"},
	{"→", 3, "00→"},
	{"→→→", 2, "→→"},
	{"日本語", 5, "00日本語"},
	{"日本語", 2, "本語"},
	{"ABCD", 4, "ABCD"},
	{"ABCD", 6, "00ABCD"},
}

func TestZFillLenRunes(t *testing.T) {
	for _, tc := range ZFillLenRunesPatterns {
		t.Run(tc.src, func(t *testing.T) {
			t.Parallel()
			res := ZFillLenRunes(tc.src, tc.count)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func BenchmarkTruncLeftRunes_CJK_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = TruncLeftRunes("日本語テスト", 3)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkTruncRightRunes_CJK_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = TruncRightRunes("日本語テスト", 3)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeftLenRunes_D_arrow_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LeftLenRunes("D", "→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRightLenRunes_D_arrow_4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RightLenRunes("D", "→", 4)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLeftLenRunes_D_CJK_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LeftLenRunes("D", "日本", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRightLenRunes_D_CJK_200(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RightLenRunes("D", "日本", 200)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkLJustLenRunes_CJK_10(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = LJustLenRunes("日本語", 10)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkRJustLenRunes_CJK_10(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = RJustLenRunes("日本語", 10)
		}
		runtime.KeepAlive(r)
	})
}

func BenchmarkZFillLenRunes_CJK_10(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r string
		for pb.Next() {
			r = ZFillLenRunes("日本語", 10)
		}
		runtime.KeepAlive(r)
	})
}
