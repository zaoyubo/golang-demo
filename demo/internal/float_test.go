package internal

import (
	"fmt"
	"math"
	"testing"
)

func TestFloat(t *testing.T) {
	var a float64 = 1.5
	var b float64 = 1.3
	var result float64 = a - b

	fmt.Printf("%.20f\n ", result)
	fmt.Println(result == 0.2)

	res := IsEqual(result, 0.2)
	fmt.Println(res)
}

// MIN 为用户自定义的比较精度
const MIN = 0.000001
const MIN_NEG = -0.000001

func IsEqual(f1, f2 float64) bool {
	sub := f1 - f2
	if sub < 0 {
		return sub > MIN_NEG
	} else {
		return sub < MIN

	}
}

func TestFloatMaxInt(t *testing.T) {
	var f1 float32 = 3421619968
	fmt.Printf("%f\n", f1+1)
	fmt.Printf("%f\n", f1)
	fmt.Println(f1 == 3421619970)
	//math.Float32bits 可以为我们打印出32位数据的二进制表示。(注：math.Float64bits可以打印64位数据的二进制)
	PrintBits32(f1)

	var f float32 = 1 << 24 // 1 << 24
	fmt.Printf("%f\n", f)
	fmt.Printf("%f\n", f+1)
	fmt.Println(f == f+1) // "true"!

	//3421619968.000000
	//3421619968.000000
	//true
	//01001111010010111111000111000111
	//Bit Pattern: 0 | 1001 1110 | 100 1011 1111 0001 1100 0111
	//16777216.000000
	//16777216.000000
	//true
}

func PrintBits32(f float32) {
	binary := fmt.Sprintf("%.32b", math.Float32bits(f))
	fmt.Println(binary)
	fmt.Printf("Bit Pattern: %s | %s %s | %s %s %s %s %s %s\n", binary[0:1],
		binary[1:5], binary[5:9], binary[9:12], binary[12:16], binary[16:20], binary[20:24], binary[24:28], binary[28:32])
}
