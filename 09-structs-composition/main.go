package main

import "fmt"

type Point2D struct {
	x, y int
}

type Point3D struct {
	Point2D
	z int
}

func main() {
	// Implement a Point2D struct with an x and y coordinate
	fmt.Println(Point2D{1,2})
	// Then implement a Point3D struct with uses the Point2D and adds a z coordinate
	vertex := Point3D{Point2D{x:1,y:2},3}
	// the fields can be retrieved directly
	fmt.Println("(", vertex.x, vertex.y, vertex.z, ")")
	// in case of ambiguity, the base component can be given as well
	fmt.Println("(", vertex.Point2D.x, vertex.Point2D.y, vertex.z, ")")
}



