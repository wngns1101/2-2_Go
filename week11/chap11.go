package main

import (
	"fmt"
	"math"
)

// 직육면체
type Cuboid struct {
	width  float64
	height float64
	length float64
}

// 부피, 겉넓이
func (c Cuboid) Volume() float64 {
	return c.width * c.height * c.length
}

func (c Cuboid) Surface() float64 {
	area := c.length * c.width * 2
	area = area + (c.length * c.height * 2)
	area = area + (c.width * c.height * 2)
	return area
}

// 구
type Sphere struct {
	radius float64
}

// 부피, 겉넓이
func (s Sphere) Volume() float64 {
	// return 4.0 / 3.0 * math.Pi * s.radius * s.radius * s.radius
	return 4.0 / 3.0 * math.Pi * math.Pow(s.radius, 3.0)
}

func (s Sphere) Surface() float64 {
	return 4.0 * math.Pi * math.Pow(s.radius, 2.0)
}

func PrintInfo(shape3d Shape3D) {
	fmt.Println(shape3d)
	fmt.Printf("부피: %0.2f\n", shape3d.Volume())
	fmt.Printf("겉넓이: %0.2f\n", shape3d.Surface())
}

type Shape3D interface {
	Volume() float64
	Surface() float64
}

func main() {
	c1 := Cuboid{width: 5.0, height: 10.5, length: 2.0}
	PrintInfo(c1)
	c2 := Sphere{radius: 3.0}
	PrintInfo(c2)
}
