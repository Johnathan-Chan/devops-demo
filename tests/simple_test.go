package tests

import (
	"fmt"
	"testing"
	"time"
)

type Student struct {
	Uuid string
	Name string
	Age int
	Sex bool
	Birthday string
	Height float64
	Weight float64
}

func (s Student) GetOne()  {
	fmt.Printf("%v\n", s)
}

func (s *Student) GetAll()  {
	fmt.Printf("%v\n", *s)
}

func TestSimple(t *testing.T) {
	s := &Student{
		Name: "zhangsan",
		Age: 12,
		Sex: true,
		Birthday: "2021/03/07",
		Height: 170.2,
		Weight: 120.5,
	}
	start := time.Now()
	s.GetOne()
	fmt.Println(time.Since(start))
	fmt.Println("-----------")
	start = time.Now()
	s.GetAll()
	fmt.Println(time.Since(start))

}
