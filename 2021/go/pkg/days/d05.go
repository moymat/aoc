package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moymat/aoc2021/pkg/helpers"
)

type Point struct {
	x int
	y int
}

type DiagramPoint struct {
	Point
	overlap int
}

type Direction = int

const (
	TopLeft = iota
	TopRight
	BottomLeft
	BottomRight
)

func convertToPoint(row string) []Point {
	values := strings.Split(row, " -> ")

	var points []Point

	for _, point := range values {
		val := strings.Split(point, ",")

		x, err := strconv.Atoi(val[0])
		helpers.CheckError(err)
		y, err := strconv.Atoi(val[1])
		helpers.CheckError(err)

		points = append(points, Point{x: x, y: y})
	}
	return points
}

func isHorizontal(points *[]Point) bool {
	return (*points)[0].y == (*points)[1].y
}

func isVertical(points *[]Point) bool {
	return (*points)[0].x == (*points)[1].x
}

func getDirection(points *[]Point) Direction {
	if (*points)[0].x < (*points)[1].x {
		if (*points)[0].y < (*points)[1].y {
			return BottomRight
		}
		return TopRight
	}
	if (*points)[0].y < (*points)[1].y {
		return BottomLeft
	}
	return BottomRight
}

func getPointFromDriagramIdx(diagram *[]DiagramPoint, point *DiagramPoint) int {
	for i, diagramPoint := range *diagram {
		if point.x == diagramPoint.x && point.y == diagramPoint.y {
			return i
		}
	}

	return -1
}

func updateDiagram(diagram *[]DiagramPoint, newPoint *DiagramPoint) {
	idx := getPointFromDriagramIdx(diagram, newPoint)
	if idx > -1 {
		(*diagram)[idx].overlap++
	} else {
		*diagram = append(*diagram, *newPoint)
	}
}

func createDiagram(pointsList *[][]Point) []DiagramPoint {
	var diagram []DiagramPoint

	for _, points := range *pointsList {
		if isHorizontal(&points) {
			y := points[0].y
			min, max := helpers.MinMax(points[0].x, points[1].x)
			for i := min; i <= max; i++ {
				diagramPoint := DiagramPoint{Point: Point{x: i, y: y}, overlap: 1}
				updateDiagram(&diagram, &diagramPoint)
			}
		} else if isVertical(&points) {
			x := points[0].x
			min, max := helpers.MinMax(points[0].y, points[1].y)
			for i := min; i <= max; i++ {
				diagramPoint := DiagramPoint{Point: Point{x: x, y: i}, overlap: 1}
				updateDiagram(&diagram, &diagramPoint)
			}
		} else {
			direction := getDirection(&points)
			if direction == BottomRight || direction == TopRight {
				minX, maxX := points[0].x, points[1].x
				y := points[0].y
				moves := 0
				for i := minX; i <= maxX; i++ {
					var diagramPoint DiagramPoint
					if direction == BottomRight {
						diagramPoint = DiagramPoint{Point: Point{x: i, y: y + moves}, overlap: 1}
					} else {
						diagramPoint = DiagramPoint{Point: Point{x: i, y: y - moves}, overlap: 1}
					}
					updateDiagram(&diagram, &diagramPoint)
					moves++
				}
			} else {
				minX, maxX := points[1].x, points[0].x
				y := points[1].y
				moves := 0
				for i := minX; i <= maxX; i++ {
					var diagramPoint DiagramPoint
					if direction == BottomLeft {
						diagramPoint = DiagramPoint{Point: Point{x: i, y: y - moves}, overlap: 1}
					} else {
						diagramPoint = DiagramPoint{Point: Point{x: i, y: y + moves}, overlap: 1}
					}
					updateDiagram(&diagram, &diagramPoint)
					moves++
				}
			}

		}
	}

	return diagram
}

func RunD05(file string) {
	inputs := helpers.GetInput("d05", file)
	rows := strings.Split(inputs, "\n")

	var pointsList [][]Point
	for _, row := range rows {
		pointsList = append(pointsList, convertToPoint(row))
	}

	diagram := createDiagram(&pointsList)

	total2 := 0
	for _, points := range diagram {
		if points.overlap >= 2 {
			total2++
		}
	}

	fmt.Println(total2)
}
