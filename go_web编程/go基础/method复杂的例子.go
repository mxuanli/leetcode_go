package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxSlice []Box

func main() {
	BoxSlice1 := BoxSlice{
		Box{10, 10, 10, RED},
		Box{15, 15, 15, WHITE},
		Box{12, 12, 12, BLUE},
		Box{13, 13, 13, YELLOW},
	}
	fmt.Printf("一共有%d个盒子\n", len(BoxSlice1))
	fmt.Println("盒子1体积", BoxSlice1[0].Volume())
	fmt.Println("盒子2颜色", BoxSlice1[0].color.colorString())
	fmt.Println("最大盒子的颜色", BoxSlice1.BiggestColor().colorString())
	BoxSlice1[0].setColor(YELLOW)
	fmt.Println("修改后盒子一的颜色", BoxSlice1[0].color.colorString())
}

func (bl BoxSlice) BiggestColor() Color { // 最大盒子的颜色
	c := Color(WHITE)
	v := 0.00
	for _, box := range bl {
		if bv := box.Volume(); bv > v {
			v = bv
			c = box.color
		}
	}
	return c
}

func (bl BoxSlice) setBlack() { // 全部设置成黑色
	for _, box := range bl {
		box.setColor(BLACK)
	}
}

func (b Box) Volume() float64 { // 计算体积
	volume := b.height * b.width * b.depth
	return volume
}

func (b *Box) setColor(c Color) { // 设置颜色
	b.color = c
}

func (c Color) colorString() string { // 颜色字节转文字
	colorStringsArray := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return colorStringsArray[c]
}
