package math

import (
	"strconv"
	"strings"
)

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const (
	sizeofInt   = 8
	sizeofInt8  = 1
	sizeofInt16 = 2
	sizeofInt32 = 4
	sizeofInt64 = 8
)

func Int2BinaryString(num int) string {
	return int2BinaryString(num, sizeofInt)
}
func Int8ToBinaryString(num int) string {
	return int2BinaryString(num, sizeofInt8)
}
func Int16ToBinaryString(num int) string {
	return int2BinaryString(num, sizeofInt16)
}
func Int32ToBinaryString(num int) string {
	return int2BinaryString(num, sizeofInt32)
}
func Int64ToBinaryString(num int) string {
	return int2BinaryString(num, sizeofInt64)
}
func int2BinaryString(num int, sizeof uint) string {
	var builder strings.Builder
	var bitCursor = sizeof << 3
	for i := bitCursor; i > 0; i-- {
		builder.WriteString(strconv.Itoa((num >> (i - 1)) & 1))
	}
	return builder.String()
}

func int2bytes(num int, sizeof int) []byte {
	bytes := make([]byte, sizeof)
	for i := 0; i < sizeof; i++ {
		base := (i+1)<<8 - 1
		bytes[sizeof-i-1] = byte(num & base)
		num >>= 8
	}
	return bytes
}

func Int32ToBytes(num int) []byte {
	return int2bytes(num, sizeofInt32)
}
