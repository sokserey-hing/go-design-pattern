package main

// Adapter: A construct which adapts an existing interface X to conform to the required interface Y.

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}} // this the interface you are given
}

// the interface we have

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	maxX += 1 // because it's zero-based
	maxY += 1 // because it's zero-based

	// preallocate the array
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := 0; j < maxX; j++ {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*' // set pixel
	}

	b := strings.Builder{}

	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	// fmt.Print(b.String())

	return b.String()

}

// adapter

type vectorToRasterAdapter struct {
	points []Point
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRasterAdapter(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		// adapter.addLine(line)
		adapter.addLineCached(line)
	}
	return &adapter // as RasterImage
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("we have generated", len(a.points), "points")
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

// using cache
var pointCache = map[[16]byte][]Point{} // 16-byte hash

func (a *vectorToRasterAdapter) addLineCached(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			a.points = append(a.points, pt)
		}
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	// be sure to add these to the cache
	pointCache[h] = a.points
	fmt.Println("generated", len(a.points), "points")
}

func main() {
	fmt.Println("----------------Adapter Pattern----------------")
	rc := NewRectangle(6, 4)
	// adapter
	a := VectorToRasterAdapter(rc)
	// fmt.Print(DrawPoints(a))
	_ = VectorToRasterAdapter(rc) // 2nd adapter
	fmt.Print(DrawPoints(a))

}
