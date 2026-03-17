![Build Status](https://github.com/ppreeper/str/actions/workflows/go.yml/badge.svg)
![Coverage Status](badges/coverage.svg)

# str

String Library for padding and other useful functions. This has been tuned for high speed operation.

## Quick Ref

This library provides several common functions:

Truncate String

- TruncLeft
- TruncRight

Common Padding

- Left
- LeftLen
- Right
- RightLen

Zero Fill Pattern

- ZFill
- ZFillLen

Justified text

- LJust
- LJustLen
- RJust
- RJustLen.

The functions Left, Right, LJust, RJust, ZFill add a pattern to the string and return the result, which is of size input + pattern \* count.

The functions LeftLen, RightLen, LJustLen, RJustLen, ZFillLen add a pattern to the string and return a string of the specified size.

```go
Left(src string, pat string, count int)
Right(src string, pat string, count int)
ZFill(src string, count int)
LJust(src string, count int)
RJust(src string, count int)
```

```go
LeftLen(src string, pat string, length int)
RightLen(src string, pat string, length int)
ZFillLen(src string, length int)
LJustLen(src string, length int)
RJustLen(src string, length int)
```

## Example Usage

Please look at the example code.

### main.go

```go
package main

import (
	"fmt"

	"github.com/ppreeper/str"
)

func main() {
	fmt.Println(str.Right("Hello", "!", 20))
	fmt.Println(str.Left("Exit now", "→", 20))
	fmt.Println("-----")
	fmt.Println(str.Left("D", "a", 4))
	fmt.Println(str.Right("D", "a", 4))
	fmt.Println(str.ZFill("D", 4))
	fmt.Println(str.LJust("D", 4))
	fmt.Println(str.RJust("D", 4))
	fmt.Println("-----")
	fmt.Println(str.Left("ABCD", "a", 5))
	fmt.Println(str.Right("ABCD", "a", 5))
	fmt.Println(str.ZFill("ABCD", 5))
	fmt.Println(str.LJust("ABCD", 5))
	fmt.Println(str.RJust("ABCD", 5))
	fmt.Println("-----")
	fmt.Println(str.LeftLen("D", "a", 4))
	fmt.Println(str.RightLen("D", "a", 4))
	fmt.Println(str.ZFillLen("D", 4))
	fmt.Println(str.LJustLen("D", 4))
	fmt.Println(str.RJustLen("D", 4))
	fmt.Println("-----")
	fmt.Println(str.LeftLen("ABCD", "a", 5))
	fmt.Println(str.RightLen("ABCD", "a", 5))
	fmt.Println(str.ZFillLen("ABCD", 5))
	fmt.Println(str.LJustLen("ABCD", 5))
	fmt.Println(str.RJustLen("ABCD", 5))
}
```

### output

```
Hello!!!!!!!!!!!!!!!!!!!!
→→→→→→→→→→→→→→→→→→→→Exit now
-----
aaaaD
Daaaa
0000D
D
    D
-----
aaaaaABCD
ABCDaaaaa
00000ABCD
ABCD
     ABCD
-----
aaaD
Daaa
000D
D
   D
-----
aABCD
ABCDa
0ABCD
ABCD
 ABCD
```

## Benchmarks

Current benchmarks

```
> go test -run ^$ -cpu 1 -benchmem -benchtime 5s -bench .
goos: linux
goarch: amd64
pkg: github.com/ppreeper/str
cpu: AMD Ryzen 7 7730U with Radeon Graphics
BenchmarkLeft_D_a_4        	157036516	        39.14 ns/op	       8 B/op	       1 allocs/op
BenchmarkLeft_D_aa_2       	162978108	        36.92 ns/op	       8 B/op	       1 allocs/op
BenchmarkLeft_D_utf8_4     	123987159	        48.55 ns/op	      16 B/op	       1 allocs/op
BenchmarkLeft_D_utf8_2_2   	120486356	        47.02 ns/op	      16 B/op	       1 allocs/op
BenchmarkLeft_D_a_200_1    	50996754	       118.6  ns/op	     416 B/op	       2 allocs/op
BenchmarkLeft_D_a_100_2    	49326704	       116.1  ns/op	     416 B/op	       2 allocs/op
BenchmarkRight_D_a_4       	147770694	        39.45 ns/op	       8 B/op	       1 allocs/op
BenchmarkRight_D_aa_2      	155074754	        42.53 ns/op	       8 B/op	       1 allocs/op
BenchmarkRight_D_utf8_4    	100000000	        55.46 ns/op	      16 B/op	       1 allocs/op
BenchmarkRight_D_utf8_2_2  	121760617	        50.69 ns/op	      16 B/op	       1 allocs/op
BenchmarkRight_D_a_200_1   	49634521	       122.4  ns/op	     416 B/op	       2 allocs/op
BenchmarkRight_D_a_100_2   	51117037	       117.2  ns/op	     416 B/op	       2 allocs/op
BenchmarkTruncLeft_ABCD_2  	1000000000	         0.4239 ns/op	   0 B/op	       0 allocs/op
BenchmarkTruncRight_ABCD_2 	1000000000	         0.4377 ns/op	   0 B/op	       0 allocs/op
BenchmarkLeftLen_D_utf8_4  	85349582	        66.18 ns/op	      32 B/op	       2 allocs/op
BenchmarkLeftLen_D_a_4     	97565781	        57.92 ns/op	      16 B/op	       2 allocs/op
BenchmarkLeftLen_D_a_200   	49306604	       133.9  ns/op	     416 B/op	       2 allocs/op
BenchmarkRightLen_D_utf8_4 	69864686	        74.59 ns/op	      32 B/op	       2 allocs/op
BenchmarkRightLen_D_a_4    	92109033	        63.33 ns/op	      16 B/op	       2 allocs/op
BenchmarkRightLen_D_a_200  	44261211	       126.5  ns/op	     416 B/op	       2 allocs/op
BenchmarkLJust_utf8_4      	353983874	        16.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkLJust_D_4         	353393116	        16.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkLJust_D_a_200     	50105112	       120.8  ns/op	     416 B/op	       2 allocs/op
BenchmarkRJust_utf8_4      	353591252	        16.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkRJust_D_4         	355942543	        16.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkRJust_D_a_200     	49439799	       122.4  ns/op	     416 B/op	       2 allocs/op
BenchmarkLJustLen_utf8_4   	205331203	        29.84 ns/op	       8 B/op	       1 allocs/op
BenchmarkLJustLen_D_4      	208440348	        27.88 ns/op	       5 B/op	       1 allocs/op
BenchmarkLJustLen_D_a_200  	50121644	       120.6  ns/op	     416 B/op	       2 allocs/op
BenchmarkRJustLen_utf8_4   	177719958	        34.20 ns/op	       8 B/op	       1 allocs/op
BenchmarkRJustLen_D_4      	206099131	        28.32 ns/op	       5 B/op	       1 allocs/op
BenchmarkRJustLen_D_a_200  	46419052	       120.5  ns/op	     416 B/op	       2 allocs/op
BenchmarkZFill_D_4         	266757704	        22.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkZFill_ABCD_8      	270810826	        22.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkZFill_D_a_200     	45792751	       117.3  ns/op	     416 B/op	       2 allocs/op
BenchmarkZFillLen_D_4      	152266942	        39.01 ns/op	       5 B/op	       1 allocs/op
BenchmarkZFillLen_ABCD_8   	142578507	        43.14 ns/op	      16 B/op	       1 allocs/op
BenchmarkZFillLen_D_a_200  	48201610	       120.3  ns/op	     416 B/op	       2 allocs/op
PASS
ok  	github.com/ppreeper/str	282.959s
```
