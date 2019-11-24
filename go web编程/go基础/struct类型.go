package main

import (
	"fmt"
)

type Skills []string

//定义一个 "人" 类型，有名字和年龄两个属性
type person struct {
	name string
	age  int
}

//学生有人属性，学科属性，技能属性
type Student struct {
	person     // 匿名字段，struct类型
	speciality string
	Skills     // 匿名字段，string slice类型
	int
	age int
}

type intStruct1 struct {
	i1, i2 int
}

func main() {
	Func1()
	Func2()
	funcStudent()
	Func3()
}

func Func1() {
	var p person // 声明一个person类型的变量
	p.name = "xuanli"
	p.age = 21
	fmt.Println("名字是", p.name)

	var p1 = person{"xiaohe", 21}
	fmt.Println("名字是", p1.name)

	p2 := person{name: "hbh", age: 21}
	fmt.Println("名字是", p2.name)

	Older(p, p1)
}

func Func2() {
	var int1 intStruct1
	int1.i1 = 10
	int1.i2 = 20
	fmt.Println(int1.i1, int1.i2)
}

func funcStudent() {
	var s Student
	s.age = 18
	s.name = "张三"
	s.speciality = "计算机"
	fmt.Println(s.name, s.age, s.speciality)
	s.person = person{"王二", 21}
	s1 := Student{person{"李四", 17}, "美术", []string{"python", "java", "golang"}, 10, 10}
	fmt.Println(s1.name, s1.age, s1.speciality, s1.Skills, s1.int)
}

func Older(p person, p1 person) {
	// 比较年龄
	if p.age > p1.age {
		fmt.Println(p.name, "年龄大")
	} else if p.age < p1.age {
		fmt.Println(p1.name, "年龄大")
	} else {
		fmt.Println("same")
	}
}

func Func3() {
	// 当struct字段和匿名类型重复的时候
	var p3 = person{"tom", 10}
	s2 := Student{speciality: "幼教", age: 20, person: p3}
	fmt.Println(s2.age)        // 会先访问外层的字段
	fmt.Println(s2.person.age) // 内层可以通过匿名字段访问
}
