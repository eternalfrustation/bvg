package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	//	"testing"
)

func test() {
	buffer := &bytes.Buffer{}
	in := New(buffer)
	center := NewPtCol(0, 0, 255, 255, 255, 255)
	corner := NewPtCol(-1, -1, 255, 255, 255, 255)
	for i := -1.0; i < 1; i += 0.1 {
		in.Lines = append(in.Lines, NewLine(center, NewPtCol(i, 0.6, 232, 54, 87, 255)))
	}
	fmt.Println("Added Lines")
	for i := -1.0; i < 1; i += 0.1 {
		in.Triangles = append(in.Triangles, NewTriangle(center, NewPtCol(i, 0.6, 232, 54, 87, 255), corner))
	}
	fmt.Println("Added Triangles")
	for i := -1.0; i < 0.998; i += 0.1 {
		in.Circles = append(in.Circles, NewCircleGrad(center, NewPtCol(i, 0.6, 232, 54, 87, 255)))
	}
	fmt.Println("Added Circles")
	for i := -1.0; i < 1; i += 0.1 {
		in.Polys = append(in.Polys, NewPoly(center, NewPtCol(i, 0.6, 232, 54, 87, 255), NewPtCol(i*i, 0.6, 232, 54, 87, 255), NewPtCol(i*i*i, 0.6, 232, 54, 87, 255)))
	}
	fmt.Println("Added Polys")
	for i := -1.0; i < 1; i += 0.1 {
		in.LineStrips = append(in.LineStrips, NewLineStrip(center, NewPtCol(i, 0.6, 232, 54, 87, 255), NewPtCol(i*i, 0.6, 232, 54, 87, 255), NewPtCol(i*i*i, 0.6, 232, 54, 87, 255)))
	}
	fmt.Println("Added LineStrips")
	fmt.Println("Encoding to memory")
	in.Encode()
//	fmt.Println(len(buffer.Bytes()))
	out, err := Decode(buffer.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	testFile, err := os.Create("testBVG.bvg")
	if err != nil {
		panic(err)
	}
	io.Copy(testFile, buffer)
	if len(in.Lines) != len(out.Lines) {
		fmt.Printf("Length of input %d \n Length of output %d \n", len(in.Lines), len(out.Lines))
		fmt.Println("The length of line array are not equal")
	}
	if len(in.Triangles) != len(out.Triangles) {
		fmt.Println("The length of triangle array are not equal")
	}
	if len(in.Circles) != len(out.Circles) {
		fmt.Println("The length of circle array are not equal")
	}
	if len(in.Polys) != len(out.Polys) {
		fmt.Println("The length of poly array are not equal")
	}
	if len(in.LineStrips) != len(out.LineStrips) {
		fmt.Println("The length of line strips array are not equal")
	}
}
