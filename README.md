![Build Status](https://github.com/ppreeper/str/actions/workflows/go.yml/badge.svg)
![Coverage Status](.github/badges/coverage.svg)

# str

String library for padding, truncation, justification, zero-filling, and banner generation. Tuned for high speed operation with 1 allocation per padding call using `[]byte` buffers and copy-doubling.

## Install

```sh
go get github.com/ppreeper/str
```

## Byte-Based Functions

These functions measure length in bytes. They are the fastest option when working with ASCII or when byte-level control is needed.

### Padding (count-based)

Prepend or append `pat` repeated `count` times:

```go
Left(src, pat string, count int) string   // "a",3 + "D" -> "aaaD"
Right(src, pat string, count int) string  // "D" + "a",3 -> "Daaa"
LJust(src string, count int) string       // "D" + spaces,3 -> "D   "
RJust(src string, count int) string       // spaces,3 + "D" -> "   D"
ZFill(src string, count int) string       // zeros,3 + "D" -> "000D"
```

### Padding (length-based)

Pad or truncate to exactly `length` bytes:

```go
LeftLen(src, pat string, length int) string   // pad left to length
RightLen(src, pat string, length int) string  // pad right to length
LJustLen(src string, length int) string       // left-justify to length
RJustLen(src string, length int) string       // right-justify to length
ZFillLen(src string, length int) string       // zero-fill to length
```

### Truncation

```go
TruncLeft(src string, length int) string   // first length bytes
TruncRight(src string, length int) string  // last length bytes
```

## Rune-Aware Functions

These functions measure length in runes (Unicode code points). Use them when working with multi-byte characters like CJK, arrows, or emoji.

### Truncation

```go
TruncLeftRunes(src string, length int) string   // first length runes
TruncRightRunes(src string, length int) string  // last length runes
```

### Padding (length-based)

Pad or truncate to exactly `length` runes:

```go
LeftLenRunes(src, pat string, length int) string   // pad left to length runes
RightLenRunes(src, pat string, length int) string  // pad right to length runes
LJustLenRunes(src string, length int) string       // left-justify to length runes
RJustLenRunes(src string, length int) string       // right-justify to length runes
ZFillLenRunes(src string, length int) string       // zero-fill to length runes
```

## Banner

`Banner` generates text banners suitable for section comment headers with configurable start/end patterns, fill patterns, padding, and optional header/footer separator lines.

```go
b := str.NewBanner()
fmt.Println(b.SPrint("section title"))
```

Output:

```
//---------------
// section title
//---------------
```

Builder-style configuration:

```go
b := str.NewBanner().
    StartPattern("##").
    Pattern("=").
    EndPattern("##").
    Padding(2).
    NoFooter()
fmt.Println(b.SPrint("heading"))
```

Or initialize from a struct:

```go
b := str.NewBannerWithConfig(str.Banner{
    SPat:   "/*",
    Pat:    "*",
    EPat:   "*/",
    Pad:    1,
    Header: true,
    Footer: true,
})
```

## Example Usage

```go
package main

import (
	"fmt"

	"github.com/ppreeper/str"
)

func main() {
	// Byte-based padding
	fmt.Println(str.Left("D", "a", 4))        // "aaaaD"
	fmt.Println(str.Right("D", "a", 4))       // "Daaaa"
	fmt.Println(str.LeftLen("D", "a", 4))      // "aaaD"
	fmt.Println(str.RightLen("D", "a", 4))     // "Daaa"
	fmt.Println(str.ZFill("D", 4))             // "0000D"
	fmt.Println(str.ZFillLen("D", 4))          // "000D"
	fmt.Println(str.LJustLen("ABCD", 8))       // "ABCD    "
	fmt.Println(str.RJustLen("ABCD", 8))       // "    ABCD"

	// Rune-aware padding
	fmt.Println(str.TruncLeftRunes("日本語", 2))          // "日本"
	fmt.Println(str.TruncRightRunes("日本語", 2))         // "本語"
	fmt.Println(str.LeftLenRunes("D", "→", 4))           // "→→→D"
	fmt.Println(str.RightLenRunes("D", "→", 4))          // "D→→→"
	fmt.Println(str.LJustLenRunes("日本語", 6))           // "日本語   "
	fmt.Println(str.RJustLenRunes("日本語", 6))           // "   日本語"

	// Banner
	b := str.NewBanner()
	fmt.Println(b.SPrint("section title"))
}
```

## Benchmarks

All padding functions produce exactly 1 heap allocation. Truncation functions are zero-allocation.

```
> go test -run ^$ -cpu 1 -benchmem -bench .
goos: linux
goarch: amd64
pkg: github.com/ppreeper/str
cpu: AMD Ryzen 7 7730U with Radeon Graphics
BenchmarkLeft_D_a_4              	75506827	        16.28 ns/op	       5 B/op	       1 allocs/op
BenchmarkLeft_D_aa_2             	77745696	        14.12 ns/op	       5 B/op	       1 allocs/op
BenchmarkLeft_D_utf8_4           	50030342	        21.54 ns/op	      16 B/op	       1 allocs/op
BenchmarkLeft_D_utf8_2_2         	57504530	        19.55 ns/op	      16 B/op	       1 allocs/op
BenchmarkLeft_D_a_200_1          	19404681	        61.16 ns/op	     208 B/op	       1 allocs/op
BenchmarkLeft_D_a_100_2          	18929103	        61.63 ns/op	     208 B/op	       1 allocs/op
BenchmarkRight_D_a_4             	53383916	        20.83 ns/op	       5 B/op	       1 allocs/op
BenchmarkRight_D_aa_2            	57542660	        18.15 ns/op	       5 B/op	       1 allocs/op
BenchmarkRight_D_utf8_4          	45921417	        28.40 ns/op	      16 B/op	       1 allocs/op
BenchmarkRight_D_utf8_2_2        	51044054	        23.38 ns/op	      16 B/op	       1 allocs/op
BenchmarkRight_D_a_200_1         	16810609	        70.48 ns/op	     208 B/op	       1 allocs/op
BenchmarkRight_D_a_100_2         	15634129	        69.06 ns/op	     208 B/op	       1 allocs/op
BenchmarkTruncLeft_ABCD_2        	1000000000	         0.4864 ns/op	   0 B/op	       0 allocs/op
BenchmarkTruncRight_ABCD_2       	1000000000	         0.4802 ns/op	   0 B/op	       0 allocs/op
BenchmarkLeftLen_D_utf8_4        	66808675	        17.04 ns/op	       4 B/op	       1 allocs/op
BenchmarkLeftLen_D_a_4           	53315006	        23.20 ns/op	       4 B/op	       1 allocs/op
BenchmarkLeftLen_D_a_200         	15900476	        73.93 ns/op	     208 B/op	       1 allocs/op
BenchmarkRightLen_D_utf8_4       	57091423	        19.01 ns/op	       4 B/op	       1 allocs/op
BenchmarkRightLen_D_a_4          	47540800	        24.05 ns/op	       4 B/op	       1 allocs/op
BenchmarkRightLen_D_a_200        	14624670	        78.93 ns/op	     208 B/op	       1 allocs/op
BenchmarkLJust_utf8_4            	53084726	        21.24 ns/op	       8 B/op	       1 allocs/op
BenchmarkLJust_D_4               	56924216	        20.18 ns/op	       5 B/op	       1 allocs/op
BenchmarkLJust_D_a_200           	15875647	        70.88 ns/op	     208 B/op	       1 allocs/op
BenchmarkRJust_utf8_4            	65033007	        17.60 ns/op	       8 B/op	       1 allocs/op
BenchmarkRJust_D_4               	63506403	        17.42 ns/op	       5 B/op	       1 allocs/op
BenchmarkRJust_D_a_200           	18586486	        62.93 ns/op	     208 B/op	       1 allocs/op
BenchmarkLJustLen_utf8_4         	79476198	        14.54 ns/op	       4 B/op	       1 allocs/op
BenchmarkLJustLen_D_4            	59799530	        19.01 ns/op	       4 B/op	       1 allocs/op
BenchmarkLJustLen_D_a_200        	15276103	        71.83 ns/op	     208 B/op	       1 allocs/op
BenchmarkRJustLen_utf8_4         	100000000	        10.74 ns/op	       4 B/op	       1 allocs/op
BenchmarkRJustLen_D_4            	75640084	        15.96 ns/op	       4 B/op	       1 allocs/op
BenchmarkRJustLen_D_a_200        	18963429	        61.38 ns/op	     208 B/op	       1 allocs/op
BenchmarkZFill_D_4               	67223497	        15.89 ns/op	       5 B/op	       1 allocs/op
BenchmarkZFill_ABCD_8            	43466224	        26.65 ns/op	      16 B/op	       1 allocs/op
BenchmarkZFill_D_a_200           	17102382	        62.12 ns/op	     208 B/op	       1 allocs/op
BenchmarkZFillLen_D_4            	67280620	        14.87 ns/op	       4 B/op	       1 allocs/op
BenchmarkZFillLen_ABCD_8         	56542996	        17.97 ns/op	       8 B/op	       1 allocs/op
BenchmarkZFillLen_D_a_200        	18416648	        61.83 ns/op	     208 B/op	       1 allocs/op
BenchmarkTruncLeftRunes_CJK_2    	89924634	        13.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkTruncRightRunes_CJK_2   	56321506	        21.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkLeftLenRunes_D_arrow_4  	34818801	        33.19 ns/op	      16 B/op	       1 allocs/op
BenchmarkRightLenRunes_D_arrow_4 	35144624	        33.13 ns/op	      16 B/op	       1 allocs/op
BenchmarkLeftLenRunes_D_CJK_200  	 8129402	       142.0 ns/op	     640 B/op	       1 allocs/op
BenchmarkRightLenRunes_D_CJK_200 	 7820985	       147.6 ns/op	     640 B/op	       1 allocs/op
BenchmarkLJustLenRunes_CJK_10    	27054973	        39.16 ns/op	      16 B/op	       1 allocs/op
BenchmarkRJustLenRunes_CJK_10    	30725977	        37.43 ns/op	      16 B/op	       1 allocs/op
BenchmarkZFillLenRunes_CJK_10    	24935230	        45.84 ns/op	      16 B/op	       1 allocs/op
BenchmarkBannerDefault           	 6933273	       186.2 ns/op	      96 B/op	       4 allocs/op
PASS
ok  	github.com/ppreeper/str	59.539s
```
