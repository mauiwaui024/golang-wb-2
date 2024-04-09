package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(3, 4)
	point2 := NewPoint(6, 8)

	// Вычисляем расстояние между ними
	distance := point1.DistanceTo(point2)

	fmt.Printf("distance between points (%.2f, %.2f) и (%.2f, %.2f) is %.2f\n", point1.x, point1.y, point2.x, point2.y, distance)
}
