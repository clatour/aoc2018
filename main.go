package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Day 6

type Point struct {
	X int
	Y int
}

type BoundingSpace struct {
	Min Point
	Max Point
	Grid []byte
}

func NewBoundingSpace(min, max Point) BoundingSpace {
	x, y := (max.X - min.X) + 1, (max.Y - min.Y) + 1
	grid := make([]byte, x * y)
	for i := range grid {
		grid[i] = 0xff
	}
	return BoundingSpace{min, max, grid}
}

func (b BoundingSpace) String() string {
	var s strings.Builder

	x, y := (b.Max.X - b.Min.X) + 1, (b.Max.Y - b.Min.Y) + 1

	for j := 0; j < x * y; j += x {
		fmt.Fprintf(&s, "%X\n", b.Grid[j:j+x])
	}

	return s.String()
}

func (p Point) translate(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}

func (b *BoundingSpace) draw(p Point, v byte) {
	translated := p.translate(-b.Min.X, -b.Min.Y)
	width, height := (b.Max.X - b.Min.X) + 1, (b.Max.Y - b.Min.Y) + 1
	index := translated.Y * width + translated.X

	if 0 <= translated.Y && translated.Y < height && 0 <= translated.X && translated.X < width {
		b.Grid[index] = v
	}
}

func (b *BoundingSpace) getRuneAtPoint(p Point) (byte) {
	translated := p.translate(-b.Min.X, -b.Min.Y)
	width, height := (b.Max.X - b.Min.X) + 1, (b.Max.Y - b.Min.Y) + 1
	index := translated.Y * width + translated.X

	if 0 <= translated.Y && translated.Y < height && 0 <= translated.X && translated.X < width {
		return b.Grid[index]
	}

	return '.'
}

func ManhattanDistance(a, b Point) int {
	x := a.X - b.X
	y := a.Y - b.Y

	if x < 0 { x = -x }
	if y < 0 { y = -y }
	return x + y
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	points := parse(f)
	min, max := boundaries(points)

	boundingSpace := NewBoundingSpace(min, max)

	//drawChar := make(map[Point]byte)
	//
	//for i, v := range points {
	//	drawChar[v] = byte(i)
	//	boundingSpace.draw(v, byte(i))
	//}
	//
	//// Part 1
	//for i := min.X; i < max.X + 1; i++ {
	//	for j := min.Y; j < max.Y + 1; j++ {
	//		m := make(map[Point]int)
	//		p := Point{i, j}
	//		for _, point := range points {
	//			v := ManhattanDistance(p, point)
	//			m[point] = v
	//		}
	//
	//		min := math.MaxInt32
	//		for _, v := range m {
	//			if v < min {
	//				min = v
	//			}
	//		}
	//
	//		closest := make([]Point, 0)
	//		for k, v := range m {
	//			if v == min {
	//				closest = append(closest, k)
	//			}
	//		}
	//
	//		if len(closest) == 1 {
	//			boundingSpace.draw(p, drawChar[closest[0]])
	//		}
	//	}
	//}
	//
	//counts := make(map[byte]int)
	//width, height := (boundingSpace.Max.X - boundingSpace.Min.X) + 1, (boundingSpace.Max.Y - boundingSpace.Min.Y) + 1
	//
	//ignore := make(map[byte]bool, 0)
	//ignore[0xff] = true
	//for i, v := range boundingSpace.Grid {
	//	x := i % width
	//	y := i / height
	//	if x != 0 && y != 0 && x != width - 1 && y != height -1 {
	//		counts[v] += 1
	//	} else {
	//		ignore[v] = true
	//	}
	//}
	//
	//fmt.Println(boundingSpace)
	//
	//largest := 0
	//for k, v := range counts {
	//	fmt.Printf("%X, %d\n", k, v)
	//
	//	if v > largest && !ignore[k] { largest = v }
	//}
	//
	//fmt.Println(largest)


	// Part 2
	for i := min.X; i < max.X + 1; i++ {
		for j := min.Y; j < max.Y + 1; j++ {
			distances := make([]int,0)
			p := Point{i, j}
			for _, point := range points {
				distances = append(distances, ManhattanDistance(p, point))
			}

			sum := 0
			for _, v := range distances {
				sum += v
			}

			if sum < 10000 {
				// Point is within region
				boundingSpace.draw(p, 0xbb)
			}
		}
	}

	area := 0
	for _, v := range boundingSpace.Grid {
		if v == 0xbb {
			area += 1
		}
	}
	fmt.Println(boundingSpace)
	fmt.Println(area)
}

func parse(r io.Reader) []Point {
	points := make([]Point, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		s1 := strings.Split(t, ",")
		one, two := s1[0], s1[1]
		x, err := strconv.Atoi(strings.TrimSpace(one))
		if err != nil {
			fmt.Println(err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(two))
		if err != nil {
			fmt.Println(err)
		}
		points = append(points, Point{x, y})
	}

	return points
}

func boundaries(points []Point) (Point, Point) {
	min, max := Point{math.MaxInt32, math.MaxInt32}, Point{0,0}

	for _, p := range points {
		if p.X > max.X { max.X = p.X }
		if p.Y > max.Y { max.Y = p.Y }

		if p.X < min.X { min.X = p.X }
		if p.Y < min.Y { min.Y = p.Y }
	}

	return min, max
}
