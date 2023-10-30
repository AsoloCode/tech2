package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для расчета расстояния между текущей точкой и другой точкой
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	// Рассчитываем расстояние между ними
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.2f\n", distance)
}
