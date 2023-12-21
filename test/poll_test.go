package test

import (
	"fmt"
	"sync"
	"testing"
)

type person struct {
	age  int
	name string
}

func TestPool(t *testing.T) {
	pool := sync.Pool{ //对象池
		New: func() interface{} {
			return new(person)
		},
	}
	p1 := pool.Get().(*person)
	p1.name = "张三"
	p1.age = 12
	fmt.Println("地址：", &p1, "，值：", *p1)
	pool.Put(p1)

	p2 := pool.Get().(*person)
	fmt.Println(p2)
}
