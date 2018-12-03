package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 3
// 2d grid
// #123 @ 3,2: 5x4
// id   @ X,Y: WxH

type Claim struct {
	ID string
	X  int
	Y  int
	W  int
	H  int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	claims := make([]Claim, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		s1 := strings.Split(s, "@")
		s2 := strings.Split(s1[1], ":")
		coords := strings.Split(s2[0], ",")

		x, y := splitCoords(coords)
		dims := strings.Split(s2[1], "x")

		w, h := splitDims(dims, err)
		claims = append(claims, Claim{ID: strings.TrimSpace(s1[0]), X: x, Y: y, W: w, H: h})
	}

	maxX, maxY := 0, 0
	for _, c := range claims {
		x, y := c.X+c.W, c.Y+c.H
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	grid := make([]int, maxX*maxY)

	for _, c := range claims {
		for j := c.Y; j < (c.Y + c.H); j++ {
			for i := c.X; i < (c.X + c.W); i++ {
				//fmt.Println(i + j * maxX)
				grid[i+j*maxX] += 1
			}
		}
	}

	// Check how many square inches of claims overlap
	overlap := totalOverlap(grid)
	fmt.Println(overlap)

	// Check if a claim doesn't overlap
	for _, c := range claims {
		if soleOwnership(c, grid, maxX) {
			fmt.Println(c.ID)
		}
	}
}

func totalOverlap(grid []int) int {
	overlap := 0
	for _, v := range grid {
		if v >= 2 {
			overlap += 1
		}
	}
	return overlap
}

func soleOwnership(c Claim, grid []int, maxX int) bool {
	for j := c.Y; j < (c.Y + c.H); j++ {
		for i := c.X; i < (c.X + c.W); i++ {
			//fmt.Println(i + j * maxX)
			if grid[i+j*maxX] != 1 {
				return false
			}
		}
	}
	return true
}

func splitDims(dims []string, err error) (int, int) {
	w, err := strconv.Atoi(strings.TrimSpace(dims[0]))
	if err != nil {
		fmt.Println(err)
	}
	h, err := strconv.Atoi(strings.TrimSpace(dims[1]))
	if err != nil {
		fmt.Println(err)
	}
	return w, h
}

func splitCoords(coords []string) (int, int) {
	x, err := strconv.Atoi(strings.TrimSpace(coords[0]))
	if err != nil {
		fmt.Println(err)
	}
	y, err := strconv.Atoi(strings.TrimSpace(coords[1]))
	if err != nil {
		fmt.Println(err)
	}
	return x, y
}
