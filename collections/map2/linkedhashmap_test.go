package map2

import (
	"fmt"
	"testing"
)

func TestNewLinkedHashMap(t *testing.T) {
	lhm := NewLinkedHashMap()
	lhm.Put("name", "linhuadong")
	lhm.Put("age", 22)
	lhm.Put("address", "nanjing")
	lhm.Put("hobby", "basketball")
	fmt.Println("--------iterate keys---------")
	for _, k := range lhm.Keys() {
		fmt.Println(k)
	}
	fmt.Println("--------iterate keys again---------")
	for _, k := range lhm.Keys() {
		fmt.Println(k)
	}
	fmt.Println("--------iterate values-------")
	for _, v := range lhm.Values() {
		fmt.Println(v)
	}
	fmt.Println("--------iterate values again-------")
	for _, v := range lhm.Values() {
		fmt.Println(v)
	}
	fmt.Println("--------Test Remove--------")
	address := lhm.Remove("address")
	fmt.Println(address)
	fmt.Println("--------iterate keys again---------")
	for _, k := range lhm.Keys() {
		fmt.Printf("%v-->%v\n", k, lhm.Get(k))
	}
}

func TestLinkedhashmap_Get(t *testing.T) {
	lhm := NewLinkedHashMap()
	lhm.Put("name", "linhuadong")
	lhm.Put("age", 22)
	lhm.Put("address", "nanjing")
	lhm.Put("hobby", "basketball")
	fmt.Println(lhm.Get("name"))
	fmt.Println(lhm.Get("sex"))
	lhm = NewLinkedHashMap()
	u1 := struct {
		name string
		age  int
	}{name: "linhuaodng", age: 22}
	u2 := struct {
		name string
		age  int
	}{name: "scathon", age: 23}
	u3 := struct {
		name string
		age  int
	}{name: "linhd", age: 24}
	lhm.Put(u1, "123")
	lhm.Put(u2, "456")
	lhm.Put(u3, "789")
	user := lhm.Get(struct {
		name string
		aget int
	}{name: "linhuadong", aget: 22})
	fmt.Println(user)
	fmt.Println("find u1: ", lhm.Get(u1))
	fmt.Println("Remove u2")
	lhm.Remove(u2)
	fmt.Println("find u2: ", lhm.Get(u2))
	fmt.Println("iterate map")
	for _, k := range lhm.Keys() {
		fmt.Printf("%v-->%v\n", k, lhm.Get(k))
	}
}

func TestLinkedhashmap_Put(t *testing.T) {
	lhm := NewLinkedHashMap()
	lhm.Put("name", "linhuadong")
	lhm.Put("age", 22)
	lhm.Put("address", "nanjing")
	lhm.Put("hobby", "basketball")
	fmt.Println("find name", lhm.Get("name"))
	fmt.Println("Cover name by new value : rayhauton")
	lhm.Put("name", "rayhauton")
	fmt.Println("find name: ", lhm.Get("name"))
	// test complex type.
	lhm = NewLinkedHashMap()
	u1 := struct {
		name string
		age  int
	}{name: "linhuaodng", age: 22}
	u2 := struct {
		name string
		age  int
	}{name: "scathon", age: 23}
	lhm.Put(u1, "123")
	lhm.Put(u2, "456")
	user := lhm.Get(struct {
		name string
		age  int
	}{name: "linhuadong", age: 22})
	fmt.Println(user)
	fmt.Println("find u1: ", lhm.Get(u1))
	fmt.Println("cover u1 with new object")
	lhm.Put(u1, "888")
	fmt.Println("find u1: ", lhm.Get(u1))
	fmt.Println("iterate map")
	for _, k := range lhm.Keys() {
		fmt.Printf("%v-->%v\n", k, lhm.Get(k))
	}
}

func TestLinkedhashmap_Size(t *testing.T) {
	lhm := NewLinkedHashMap()
	lhm.Put("name", "linhuadong")
	lhm.Put("age", 22)
	lhm.Put("address", "nanjing")
	lhm.Put("hobby", "basketball")
	fmt.Printf("map size: %d\n", lhm.Size())
}

func TestTableSizeFor(t *testing.T) {
	fmt.Println(tableSizeFor(-1))
	fmt.Println(tableSizeFor(7))
	fmt.Println(tableSizeFor(16))
	fmt.Println(tableSizeFor((1 << 30) - 2))
}
