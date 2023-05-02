![Build Status](https://github.com/ppreeper/str/actions/workflows/go.yml/badge.svg)

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
cpu: AMD Ryzen 7 7735HS with Radeon Graphics
BenchmarkLeft_D_a_4             177426610               33.98 ns/op            4 B/op          1 allocs/op
BenchmarkLeft_D_aa_2            199887630               29.80 ns/op            4 B/op          1 allocs/op
BenchmarkLeft_D_utf8_4          147985876               40.27 ns/op           16 B/op          1 allocs/op
BenchmarkLeft_D_utf8_2_2        164434153               36.02 ns/op           16 B/op          1 allocs/op
BenchmarkLeft_D_a_200_1         56919225               103.7 ns/op           416 B/op          2 allocs/op
BenchmarkLeft_D_a_100_2         60211620               100.8 ns/op           416 B/op          2 allocs/op
BenchmarkRight_D_a_4            173296929               34.60 ns/op            4 B/op          1 allocs/op
BenchmarkRight_D_aa_2           196698732               30.74 ns/op            4 B/op          1 allocs/op
BenchmarkRight_D_utf8_4         151394588               40.14 ns/op           16 B/op          1 allocs/op
BenchmarkRight_D_utf8_2_2       168813403               35.31 ns/op           16 B/op          1 allocs/op
BenchmarkRight_D_a_200_1        57623094               104.8 ns/op           416 B/op          2 allocs/op
BenchmarkRight_D_a_100_2        60113691               102.1 ns/op           416 B/op          2 allocs/op
BenchmarkTruncLeft_ABCD_2       324979262               18.32 ns/op            0 B/op          0 allocs/op
BenchmarkTruncRight_ABCD_2      329157414               18.20 ns/op            0 B/op          0 allocs/op
BenchmarkLeftLen_D_utf8_4       55832942               107.5 ns/op            32 B/op          2 allocs/op
BenchmarkLeftLen_D_a_4          84549571                70.22 ns/op           16 B/op          2 allocs/op
BenchmarkLeftLen_D_a_200         4705328              1282 ns/op            1520 B/op          4 allocs/op
BenchmarkRightLen_D_utf8_4      55781134               108.3 ns/op            32 B/op          2 allocs/op
BenchmarkRightLen_D_a_4         85283271                69.94 ns/op           16 B/op          2 allocs/op
BenchmarkRightLen_D_a_200        4587787              1279 ns/op            1520 B/op          4 allocs/op
BenchmarkLJust_utf8_4           175381755               34.65 ns/op            4 B/op          1 allocs/op
BenchmarkLJust_D_4              176793922               33.80 ns/op            4 B/op          1 allocs/op
BenchmarkLJust_D_a_200          58400977               103.0 ns/op           416 B/op          2 allocs/op
BenchmarkRJust_utf8_4           174684237               33.89 ns/op            4 B/op          1 allocs/op
BenchmarkRJust_D_4              173668522               34.42 ns/op            4 B/op          1 allocs/op
BenchmarkRJust_D_a_200          58833978               103.2 ns/op           416 B/op          2 allocs/op
BenchmarkLJustLen_utf8_4        75187804                77.23 ns/op           16 B/op          2 allocs/op
BenchmarkLJustLen_D_4           84242854                69.98 ns/op           16 B/op          2 allocs/op
BenchmarkLJustLen_D_a_200        4625905              1282 ns/op            1520 B/op          4 allocs/op
BenchmarkRJustLen_utf8_4        84421971                73.52 ns/op           16 B/op          2 allocs/op
BenchmarkRJustLen_D_4           84456016                70.90 ns/op           16 B/op          2 allocs/op
BenchmarkRJustLen_D_a_200        4756957              1275 ns/op            1520 B/op          4 allocs/op
BenchmarkZFill_D_4              174243123               34.84 ns/op            4 B/op          1 allocs/op
BenchmarkZFill_ABCD_8           164768880               36.71 ns/op            8 B/op          1 allocs/op
BenchmarkZFill_D_a_200          56639906               102.5 ns/op           416 B/op          2 allocs/op
BenchmarkZFillLen_D_4           82058068                70.89 ns/op           16 B/op          2 allocs/op
BenchmarkZFillLen_ABCD_8        57930297               102.0 ns/op            24 B/op          2 allocs/op
BenchmarkZFillLen_D_a_200        4752658              1272 ns/op            1520 B/op          4 allocs/op
PASS
ok      github.com/ppreeper/str 288.716s
```
