package stream

import (
	"fmt"
	"strings"
	"testing"
)

func TestStream_Filter(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5, 6}
	stream := NewStream(nums)
	result := stream.Collect()
	fmt.Println(result)
	result = stream.Filter(func(num interface{}) bool {
		n := num.(int)
		return n%2 == 0
	}).Collect()
	fmt.Println(result)
}

func TestStream_Foreach(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5, 6}
	stream := NewStream(nums)
	var consumerFunc Consumer = func(e interface{}) {
		fmt.Printf("%v\n", e)
	}
	stream.Foreach(consumerFunc)
	fmt.Println("=======================")
	stream.Filter(func(num interface{}) bool {
		n := num.(int)
		return n%2 == 0
	}).Foreach(consumerFunc)
	fmt.Println("End")
}

func TestStream_Map(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5, 6}
	stream := NewStream(nums)
	stream.Map(func(e interface{}) interface{} {
		num := e.(int)
		return num + 2
	}).Filter(func(e interface{}) bool {
		return e.(int)%2 == 0
	}).Foreach(func(e interface{}) {
		fmt.Printf("%-3v", e)
	})
	fmt.Println()
}
func TestStream_Map2(t *testing.T) {
	nums := []interface{}{"linhuadong", "scathon", "scathonlin", "ScathonLin", "RayHauton"}
	stream := NewStream(nums)
	stream.Map(func(e interface{}) interface{} {
		name := e.(string)
		return strings.ToUpper(name)
	}).Filter(func(e interface{}) bool {
		return strings.Contains(e.(string), "SCATHON")
	}).Foreach(func(e interface{}) {
		fmt.Println(e)
	})
	fmt.Println()
}

func TestStream_Map3(t *testing.T) {
	type user struct {
		age int
	}
	nums := []interface{}{1, 2, 3, 4, 5, 6}
	stream := NewStream(nums)
	stream.Map(func(e interface{}) interface{} {
		return user{e.(int)}
	}).Filter(func(e interface{}) bool {
		return e.(user).age%2 == 0
	}).Foreach(func(e interface{}) {
		fmt.Println(e)
	})
	fmt.Println()
}

func TestStream_Count(t *testing.T) {
	type user struct {
		age int
	}
	nums := []interface{}{1, 2, 3, 4, 5, 6}
	stream := NewStream(nums)
	cnt := stream.Map(func(e interface{}) interface{} {
		return user{e.(int)}
	}).Filter(func(e interface{}) bool {
		return e.(user).age%2 == 0
	}).Count()
	fmt.Println(cnt)
}
func TestStream_Count2(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5, 6}
	stream := NewStream(nums)
	cnt := stream.Count()
	fmt.Println(cnt)
}

func TestStream_Distinct(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	ages := []interface{}{1, 2, 6, 8, 2, 5, 9, 5}
	names := []string{"A", "B", "D", "C", "B", "I", "H", "G"}
	i := -1
	stream := NewStream(ages)
	res := stream.Map(func(e interface{}) interface{} {
		i++
		return user{names[i], e.(int)}
	}).Distinct(func(e1, e2 interface{}) int {
		u1 := e1.(user)
		u2 := e2.(user)
		nameCompare := strings.Compare(u1.name, u2.name)
		if nameCompare != 0 {
			return nameCompare
		}
		return u1.age - u2.age
	})
	for _, u := range res {
		fmt.Println(u)
	}
}

func TestStream_CollectToMap(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	ages := []interface{}{1, 2, 6, 8, 5, 9, 5}
	names := []string{"A", "B", "D", "C", "I", "H", "G"}
	i := -1
	stream := NewStream(ages)
	res := stream.Map(func(e interface{}) interface{} {
		i++
		return user{names[i], e.(int)}
	}).CollectToMap(func(e interface{}) interface{} {
		u := e.(user)
		return u.name
	}, func(e interface{}) interface{} {
		u := e.(user)
		return u
	})
	fmt.Println(res)
}
