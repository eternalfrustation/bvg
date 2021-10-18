package bvg

import (
	"bytes"
	"math"
	"io"
	"fmt"
	"os"
	"testing"
)

func TestBvg(t *testing.T) {
	buffer := &bytes.Buffer{}
	in := New(buffer)
	center := NewPtCol(0, 0, 255, 255, 255, 255)
	corner := NewPtCol(-1, -1, 255, 255, 255, 255)
	bytesEncoded := 0
	for i := -1.0; i < 1; i += 0.1 {
		in.Lines = append(in.Lines, NewLine(center, NewPtCol(i, 0.6, 232, 54, 87, 255)))
		bytesEncoded += 41
	}
	t.Log("Added Lines" + fmt.Sprintf(" bytes Encoded: %d", bytesEncoded))
	for i := -1.0; i < 1; i += 0.1 {
		in.Triangles = append(in.Triangles, NewTriangle(center, NewPtCol(i, 0.6, 232, 54, 87, 255), corner))
		bytesEncoded += 61
	}
	t.Log("Added Triangles" + fmt.Sprintf(" bytes Encoded: %d", bytesEncoded))
	for i := -1.0; i < 0.998; i += 0.1 {
		in.Circles = append(in.Circles, NewCircle(center.RelPt(i, math.Pi*i), 0.5, 0.999))
		bytesEncoded += 41
	}
	t.Log("Added Circles" + fmt.Sprintf(" bytes Encoded: %d", bytesEncoded))
	for i := -1.0; i < 1; i += 0.1 {
		in.Polys = append(in.Polys, NewPoly(center, NewPtCol(i, 0.6, 232, 54, 87, 255), NewPtCol(i*i, 0.6, 232, 54, 87, 255), NewPtCol(i*i*i, 0.6, 232, 54, 87, 255)))
		bytesEncoded += 1 + 4 + 60
	}
	t.Log("Added Polys" + fmt.Sprintf(" bytes Encoded: %d", bytesEncoded))
	for i := -1.0; i < 1; i += 0.1 {
		in.LineStrips = append(in.LineStrips, NewLineStrip(center, NewPtCol(i, 0.6, 232, 54, 87, 255), NewPtCol(i*i, 0.6, 232, 54, 87, 255), NewPtCol(i*i*i, 0.6, 232, 54, 87, 255)))
		bytesEncoded += 1 + 4 + 60
	}
	t.Log("Added LineStrips" + fmt.Sprintf(" bytes Encoded: %d", bytesEncoded))
	t.Log("Encoding to memory")
	in.Encode()
//	fmt.Println(len(buffer.Bytes()))
	out, err := Decode(buffer.Bytes())
	if err != nil {
		t.Error(err)
	}
	testFile, err := os.Create("testBVG.bvg")
	if err != nil {
		panic(err)
	}
	io.Copy(testFile, buffer)
	if len(in.Lines) != len(out.Lines) {
		t.Errorf("Length of input %d \n Length of output %d \n", len(in.Lines), len(out.Lines))
		t.Error("The length of line array are not equal")
	}
	if len(in.Triangles) != len(out.Triangles) {
		t.Error("The length of triangle array are not equal")
	}
	if len(in.Circles) != len(out.Circles) {
		t.Error("The length of circle array are not equal")
	}
	if len(in.Polys) != len(out.Polys) {
		t.Error("The length of poly array are not equal")
	}
	if len(in.LineStrips) != len(out.LineStrips) {
		t.Error("The length of line strips array are not equal")
	}
}
