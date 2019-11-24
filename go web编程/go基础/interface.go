package main

import (
	"fmt"
)

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human
	school string
	loan   float32
}

type employee struct {
	human
	company string
	money   float32
}

func (h *human) sayHi() {
	fmt.Print("名字:%s，手机号%s\n", h.name, h.phone)
}

func (h *human) sing(lyrics string) {
	fmt.Println("歌词", lyrics)
}

func (h *human) Guzzle(beerStein string) {
	fmt.Println("酒杯", beerStein)
}

func (e employee) sayHi() {
	fmt.Print("名字:%S，手机号%s，公司%s", e.name, e.phone, e.company)
}

func (s *student) borrowMoney(amount float32) {
	s.loan += amount
}

func (e *employee) spendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	sayHi()
	sing(lyrics string)
	guzzle(beerStein string)
}

type youngChap interface {
	sayHi()
	sing(lyrics string)
	borrowMoney(amount float32)
}

type ElderlyGent interface {
	sayHi()
	sing(lyrics string)
	spendSalary(amount float32)
}

func main() {

}
