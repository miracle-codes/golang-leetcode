package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	var head Student
	head.Name = "lin"
	head.Age = 24
	head.Score = 99

	//尾部插入法
	var tail = &head
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}

	trans(&head)

}
