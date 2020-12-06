package hash

import (
	"fmt"
	"strconv"
	"unsafe"
)

func GetAddressOfObject(obj interface{}) uint64 {
	return *(*uint64)(unsafe.Pointer(&obj))
}

func HashCode(obj interface{}) int32 {
	hashCodeStr := fmt.Sprintf("%d", unsafe.Pointer(&obj))
	hashcode, err := strconv.Atoi(hashCodeStr)
	if err != nil {
		panic("failed to get hashcode")
	}
	return int32(hashcode)
}
