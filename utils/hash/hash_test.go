package hash

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestHashCode(t *testing.T) {
	obj := struct {
		name string
	}{name: "scathon"}
	hashcode := HashCode(obj)
	fmt.Println(hashcode)
}

func TestUnsafePointer(t *testing.T) {
	u := struct {
		name string
		age  int
	}{name: "linhuadong", age: 22}
	p := unsafe.Pointer(&u)
	fmt.Println(p)
	fmt.Println(&u)
}
func TestPointer(t *testing.T) {
	u := struct {
		name string
		age  int
	}{name: "linhuadong", age: 22}
	fmt.Printf("pointer is : %p\n", &u)
	pointer := fmt.Sprintf("%p", &u)
	fmt.Printf("pointer is %s\n", pointer)
	p := unsafe.Pointer(&u)
	fmt.Println(p)
}

func TestClosure(t *testing.T) {
	square := func() func() int {
		var x int
		return func() int {
			x++
			return x * x
		}
	}()
	fmt.Printf("%d\n", square())
	fmt.Printf("%d\n", square())
	fmt.Printf("%d\n", square())
	fmt.Printf("%d\n", square())
}
