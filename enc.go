package bvg

import (
	"encoding/binary"
	//	"fmt"
	//	"fmt"
	"io"
	"math"
	//	"os"
)

// This stores the writer and commands to be written
type Bvg struct {
	Writer     io.Writer
	Reader     io.Reader
	Points     []*Point
	Lines      []*Line
	Circles    []*Circle
	Triangles  []*Triangle
	Polys      []*Poly
	Bezs       []*Bez
	LineStrips []*LineStrip
}

// A point must lie between (-1,-1) and (1,1)
// anything outside will not be rendered
type Point struct {
	X, Y       float64
	R, G, B, A uint8
}

// Returns a new Point with the given Position
func NewPt(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Dist(p1 *Point) float64 {
	return math.Sqrt((p.X-p1.X)*(p.X-p1.X) + (p.Y-p1.Y)*(p.Y-p1.Y))
}

func (p *Point) Write(w io.Writer) {
	binary.Write(w, binary.LittleEndian, p.X)
	binary.Write(w, binary.LittleEndian, p.Y)
	binary.Write(w, binary.LittleEndian, p.R)
	binary.Write(w, binary.LittleEndian, p.G)
	binary.Write(w, binary.LittleEndian, p.B)
	binary.Write(w, binary.LittleEndian, p.A)
}

// Returns a new Point with the given position and color
func NewPtCol(x, y float64, r, g, b, a uint8) *Point {
	return &Point{
		X: x,
		Y: y,
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (p *Point) RelPt(l, theta float64) *Point {
	return &Point{
		X: p.X + l*math.Cos(theta),
		Y: p.Y + l*math.Sin(theta),
		R: p.R,
		G: p.G,
		B: p.B,
		A: p.A,
	}
}

type Line struct {
	P1 *Point
	P2 *Point
}

func NewLine(p1, p2 *Point) *Line {
	return &Line{
		P1: p1,
		P2: p2,
	}
}

type Circle struct {
	P *Point
	// r is the complete radius of the circle 
	// the alpha at r is 0
	// t is threshold upto which the color of the circle 
	// does not fade
	R, T float64
}

func NewCircle(p *Point, r, t float64) *Circle {
	return &Circle{
		P:  p,
		R: r,
		T: t,
	}
}

type Triangle struct {
	P1, P2, P3 *Point
}

func NewTriangle(p1, p2, p3 *Point) *Triangle {
	return &Triangle{
		P1: p1,
		P2: p2,
		P3: p3,
	}
}

type Poly struct {
	Pts []*Point
}

func NewPoly(p ...*Point) *Poly {
	return &Poly{
		Pts: p,
	}
}

type Bez struct {
	Pts []*Point
}

func NewBez(Points ...*Point) *Bez {
	return &Bez{
		Pts: Points,
	}
}

type LineStrip struct {
	Pts []*Point
}

func NewLineStrip(pts ...*Point) *LineStrip {
	return &LineStrip{
		Pts: pts,
	}
}

// Returns a new bvg struct with the specified writer
func New(writer io.Writer) *Bvg {
	return &Bvg{
		Writer: writer,
	}
}

func (b *Bvg) Encode() {
	for _, val := range b.Lines {
		b.DrawLine(*val)
	}
	for _, val := range b.Triangles {
		b.DrawTriangle(*val)
	}
	for _, val := range b.Circles {
		b.DrawCircle(*val)
	}
	for _, val := range b.Polys {
		b.DrawPoly(*val)
	}
	for _, val := range b.LineStrips {
		b.DrawLineStrip(*val)
	}
}

// Draws a line from point p1 to p2 with the
// colors inside the point
func (b *Bvg) DrawLine(l Line) {
	binary.Write(b.Writer, binary.LittleEndian, int8('l'))
	l.P1.Write(b.Writer)
	l.P2.Write(b.Writer)
}

// draw a triangle from p1, p2, p3
func (b *Bvg) DrawTriangle(t Triangle) {
	binary.Write(b.Writer, binary.LittleEndian, int8('t'))
	t.P1.Write(b.Writer)
	t.P2.Write(b.Writer)
	t.P3.Write(b.Writer)
}

// draw a circle with p1 as center and distance between p1
// an p2 as radius with colour gradient from colors of p1 at
// the center to colours of p2 at the circumference
func (b *Bvg) DrawCircle(c Circle) {
	binary.Write(b.Writer, binary.LittleEndian, int8('c'))
	c.P.Write(b.Writer)
	binary.Write(b.Writer, binary.LittleEndian, c.R)
	binary.Write(b.Writer, binary.LittleEndian, c.T)
}

// draws a polygon from the points given, it triangulates by
// using a triangle strip from the given points
func (b *Bvg) DrawPoly(p Poly) {
	binary.Write(b.Writer, binary.LittleEndian, int8('p'))
	binary.Write(b.Writer, binary.LittleEndian, uint32(len(p.Pts)))
	//	fmt.Println(len(p.Pts))
	for _, p := range p.Pts {
		p.Write(b.Writer)
	}
}

// Draws a bezier curve with 1st and last point being
// Starting and Ending points and the rest inbetween being
// control points
func (b *Bvg) DrawBez(bez Bez) {
	binary.Write(b.Writer, binary.LittleEndian, int8('b'))
	binary.Write(b.Writer, binary.LittleEndian, uint32(len(bez.Pts)))
	for _, p := range bez.Pts {
		p.Write(b.Writer)
	}
}

// Draw a line strip from the given points
func (b *Bvg) DrawLineStrip(l LineStrip) {
	binary.Write(b.Writer, binary.LittleEndian, int8(60))
	binary.Write(b.Writer, binary.LittleEndian, uint32(len(l.Pts)))
	for _, p := range l.Pts {
		p.Write(b.Writer)

	}
}
