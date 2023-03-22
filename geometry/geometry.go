package geometry

import (
	"fmt"
	"math"
)

const PRECISION = math.SmallestNonzeroFloat64

func Sqrt(x float64) (z float64, i int) {
	z = 1.0
	for i = 0; ; i++ {
		diff := (z*z - x)
		if math.Abs(diff) < PRECISION {
			return
		}
		z -= diff / (2 * z)
	}
}

type Point struct {
	X float64
	Y float64
}

func (p Point) translate(x, y float64) Point {
	p.X += x
	p.Y += y
	return p
}

type Figure struct {
	points []Point
}

func (figure *Figure) add_point(point Point) {
	figure.points = append(figure.points, point)
}

func (figure Figure) circumference() (distance float64) {
	distance = 0.0
	points := append(figure.points, figure.points[0])
	for i := 0; i < len(points)-1; i++ {
		line := LineSegment{points[i], points[i+1]}
		distance += line.distance()
	}
	return
}

type LineSegment struct {
	p1, p2 Point
}

func (l LineSegment) distance() (distance float64) {
	fmt.Println(l)
	p1 := l.p1
	p2 := l.p2
	x := math.Abs(p1.X - p2.X)
	y := math.Abs(p1.Y - p2.Y)
	distance = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	fmt.Println("distance", distance)
	return
}

func main() {
	p1 := Point{0, 0}
	p2 := Point{0, 5}
	p3 := Point{5, 5}
	p4 := Point{5, 0}

	points := []Point{p1, p2, p3, p4, p1}
	figure := Figure{}
	for _, point := range points {
		newp := point.translate(10, 0)
		figure.add_point(newp)
	}
	r := figure.circumference()
	fmt.Println(r, points)
	fmt.Println(len(points), cap(points))
}
