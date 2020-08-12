package math

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func TestMaxInt(t *testing.T) {
	a, b := 1, 2
	println(MaxInt(a, b))
}

func TestMinInt(t *testing.T) {
	a, b := 1, 2
	println(MinInt(a, b))
}

func TestInt2BinaryString(t *testing.T) {
	fmt.Println(Int2BinaryString(28))
	fmt.Println(Int2BinaryString(math.MaxInt64))
	fmt.Println(Int8ToBinaryString(41))
	fmt.Println(Int8ToBinaryString(math.MaxInt8))
}

func TestSizeofInt(t *testing.T) {
	var i1 int = 1
	var i2 int8 = 1
	var i3 int16 = 1
	var i4 int32 = 1
	var i5 int64 = 1
	fmt.Printf("sizeof:%v, max: %d\n", unsafe.Sizeof(i1), math.MaxInt8)
	fmt.Printf("sizeof:%v, max: %d\n", unsafe.Sizeof(i2), math.MaxInt8)
	fmt.Printf("sizeof:%v, max: %d\n", unsafe.Sizeof(i3), math.MaxInt16)
	fmt.Printf("sizeof:%v, max: %d\n", unsafe.Sizeof(i4), math.MaxInt32)
	fmt.Printf("sizeof:%v, max: %d\n", unsafe.Sizeof(i5), math.MaxInt64)
	fmt.Println("======================")
	var i6 uint = 1
	var i7 uint8 = 1
	var i8 uint16 = 1
	var i9 uint32 = 1
	var i10 uint64 = 1
	fmt.Println(unsafe.Sizeof(i6))
	fmt.Println(unsafe.Sizeof(i7))
	fmt.Println(unsafe.Sizeof(i8))
	fmt.Println(unsafe.Sizeof(i9))
	fmt.Println(unsafe.Sizeof(i10))
}

func TestInt32ToBytes(t *testing.T) {
	bytes := Int32ToBytes(13)
	fmt.Printf("%v\n", bytes)

	num := 1239878979
	binStr := Int32ToBinaryString(num)
	fmt.Println(binStr)
	bytes = Int32ToBytes(num)
	var res string
	for _, bt := range bytes {
		res += Int8ToBinaryString(int(bt))
	}
	fmt.Println(res)
}
