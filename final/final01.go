// 학번: 201944096
// 성명: 이주훈

package main

import (
	"fmt"
	"math"
	"time"
)

type won struct {
	radius float64
	height float64
}

func (c won) Volume() float64 {
	return math.Pi * c.radius * c.radius * c.height
}

func (c won) Surface() float64 {
	return (2 * math.Pi * c.radius * c.radius) + (2 * 3.14 * c.radius * c.height)
}

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
	fmt.Print(shape3d, " 부피는")
	fmt.Printf("%0.2f 입니다\n", shape3d.Volume())
	fmt.Print("겉넓이는 ")
	fmt.Printf("%0.2f 입니다\n", shape3d.Surface())
	// fmt.Printf("%s 부피: %0.2f\n", shape3d, shape3d.Volume())
	// fmt.Printf("겉넓이: %0.2f\n", shape3d.Surface())
}

type Shape3D interface {
	Volume() float64
	Surface() float64
}

func main() {
	c1 := []Shape3D{}
	c1 = append(c1, won{10, 10}, Cuboid{10, 2, 5}, Sphere{10}, won{2, 6})
	time.Sleep(time.Second * 5)
	go PrintInfo(c1[0])
	go PrintInfo(c1[1])
	go PrintInfo(c1[2])
	go PrintInfo(c1[3])
	time.Sleep(time.Second)
}
