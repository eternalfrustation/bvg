package bvg

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// This stores the writer and commands to be written
type Bvg struct {
	W io.Writer
}

// A point must lie between (-1,-1) and (1,1)
// anything outside will not be rendered
type Point struct {
	X, Y       float64
	R, G, B, A uint8
}

func main() {
	bvgfile, err := os.Create("test.bvg")
	bvg := New(bvgfile)
	bvg.DrawLine(Point{1, 1, 255, 255, 255, 255}, Point{-1, -1, 255, 255, 255, 255})
	bvg.DrawRect(Point{0.5, 0.5, 255, 255, 255, 255}, Point{-0.5, -0.5, 255, 255, 255, 255})
	bvg.DrawCircle(Point{0, 0, 255, 255, 255, 255}, Point{0.3, 0.3, 255, 255, 255, 255})
	fmt.Println(err)
}

// Returns a new bvg struct with the specified writer
func New(writer io.Writer) *Bvg {
	return &Bvg{
		W: writer,
	}
}

// Draws a line from point p1 to p2 with the
// colors inside the point
func (b *Bvg) DrawLine(p1, p2 Point) {
	binary.Write(b.W, binary.LittleEndian, int8(108))
	binary.Write(b.W, binary.LittleEndian, p1)
	binary.Write(b.W, binary.LittleEndian, p2)
}

// Draw a rectangle from point p1 to p2
func (b *Bvg) DrawRect(p1, p2 Point) {
	binary.Write(b.W, binary.LittleEndian, int8(114))
	binary.Write(b.W, binary.LittleEndian, p1)
	binary.Write(b.W, binary.LittleEndian, p2)
}

// draw a circle with p1 as center and distance between p1
// an p2 as radius with colour gradient from colors of p1 at
// the center to colours of p2 at the circumference
func (b *Bvg) DrawCircle(p1, p2 Point) {
	binary.Write(b.W, binary.LittleEndian, int8(99))
	binary.Write(b.W, binary.LittleEndian, p1)
	binary.Write(b.W, binary.LittleEndian, p2)

}

// draws a polygon from the points given, it triangulates by
// drawing a triagle from the firts three points, then it
// draws a triangle from the 2nd 3rd and 4th point and so on
func (b *Bvg) DrawPoly(pts ...Point) {
	binary.Write(b.W, binary.LittleEndian, int8(112))
	binary.Write(b.W, binary.LittleEndian, int32(len(pts)))
	for _, p := range pts {
		binary.Write(b.W, binary.LittleEndian, p)
	}
}

// Draws a bezier curve with 1st and last point being
// Starting and Ending points and the rest inbetween being
// control points
func (b *Bvg) DrawBez(pts ...Point) {
	binary.Write(b.W, binary.LittleEndian, int8(98))
	binary.Write(b.W, binary.LittleEndian, int32(len(pts)))
	for _, p := range pts {
		binary.Write(b.W, binary.LittleEndian, p)
	}
}
