package main

import (
	"fmt"
	"math"
)

// Point представляет точку на плоскости с координатами x и y.
type Point struct {
	x float64
	y float64
}

// NewPoint создает новую точку с заданными координатами x и y.
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// DistanceTo вычисляет расстояние между текущей точкой и переданной точкой p.
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создание двух точек
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисление расстояния между точками
	distance := point1.DistanceTo(point2)

	// Форматирование строки для вывода числа с плавающей точкой с 2 знаками после запятой.
	fmt.Printf("Distance between point1 and point2: %.2f\n", distance)
}
