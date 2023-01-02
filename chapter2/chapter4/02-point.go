package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := Point{0.0, 0.0}
	p2 := Point{4.0, 3.0}
	fmt.Println("get distance", getDis(p1, p2))
	fmt.Println("get distance", p2.getDistance(p1))
}

type Point struct {
	x, y float64
}

// 函数：计算距离
func getDis(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

// 方法：站在面相对象的角度
func (p2 *Point) getDistance(p1 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}
