package main

import (
	"fmt"
	"math"
)

func main() {
	s := square{0.3}
	c := circle{0.3}

	sp := [3]shaper{&s, &c}
	for _, x := range sp {
		//type assertion
		if t, ok := x.(*square); ok {
			fmt.Printf("The type is :%T\n", t)
		}
		if t, ok := x.(*circle); ok {
			fmt.Printf("The type is :%T\n", t)
			//fmt.Println("The type is *circle.")
		}
		if x == nil {
			fmt.Println("nil")
			continue
		}
		fmt.Println("Area is", x.area())
	}

	for _, x := range sp {
		//type-switch
		switch t := x.(type) {
		case *square:
			fmt.Printf("Type square %T with value %v and area %f\n", t, t, t.area())
		case *circle:
			fmt.Printf("Type circle %T with value %v and area %f\n", t, t, t.area())
		case nil:
			fmt.Println("nil!")
		default:
			fmt.Println("Unexpected type %T\n", t)
		}
	}

	for _, x := range sp {
		if xa, ok := x.(shaper); ok {
			fmt.Printf("%#v implements shaper.area():%f\n", x, xa.area())
		}
	}
}

type shaper interface {
	area() float32
}

type square struct {
	edge float32
}

type circle struct {
	radius float32
}

func (s *square) area() float32 {
	return s.edge * s.edge
}

func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
