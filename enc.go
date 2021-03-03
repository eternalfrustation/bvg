package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Bvg struct {
	W io.Writer
}

// A point must lie between (-1,-1) and (1,1)
// anything outside will not be rendered
type Point struct {
	X, Y    float64
	R, G, B uint8
}

func main() {
	bvgfile, err := os.Create("test.bvg")
	bvg := New(bvgfile)
	bvg.DrawLine(Point{1, 1, 255, 255, 255}, Point{-1, -1, 255, 255, 255})
	fmt.Println(err)
}

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

func (b *Bvg) DrawRect(p1, p2 Point) {
	binary.Write(b.W, binary.LittleEndian, int8(114))
	binary.Write(b.W, binary.LittleEndian, p1)
	binary.Write(b.W, binary.LittleEndian, p2)
}

func (b *Bvg) DrawCircle(p1, p2 Point) {
	binary.Write(b.W, binary.LittleEndian, int8(99))
	binary.Write(b.W, binary.LittleEndian, p1)
	binary.Write(b.W, binary.LittleEndian, p2)

}

func (b *Bvg) DrawPoly(pts ...Point) {
	binary.Write(b.W, binary.LittleEndian, int8(112))
	binary.Write(b.W, binary.LittleEndian, int16(len(pts)))
	for _, p := range pts {
		binary.Write(b.W, binary.LittleEndian, p)
	}
}

func (b *Bvg) DrawBez(pts ...Point) {
	binary.Write(b.W, binary.LittleEndian, int8(98))
	binary.Write(b.W, binary.LittleEndian, int16(len(pts)))
	for _, p := range pts {
		binary.Write(b.W, binary.LittleEndian, p)
	}
}
