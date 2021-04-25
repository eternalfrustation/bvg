package bvg

import (
	"encoding/binary"
	"bytes"
	"fmt"
	//"os"
	"math"
)

// Decodes the given bytes and returns a pointer to
// an Bvg struct, You need to set the writer for this new
// struct
func Decode(b []byte) (bvg *Bvg , err error) {
	buffer := &bytes.Buffer{}
	bvg = New(buffer)
	bvg.Reader = bytes.NewBuffer(b)
	for i := 0; i < len(b)-19; {
//		fmt.Println(string(b[i]) + " at " + fmt.Sprintf("%dth byte", i))
		switch b[i] {
		case byte('l'):
			Pts, n := PointsFromBytes(b[i+1:i+41], 2)
			bvg.Lines = append(bvg.Lines, NewLine(Pts[0], Pts[1]))
			i += n + 1
		case byte('t'):
			Pts, n := PointsFromBytes(b[i+1:i+61], 3)
			bvg.Triangles = append(bvg.Triangles, NewTriangle(
				Pts[0],
				Pts[1],
				Pts[2],
			))
			i += n + 1
		case byte('c'):
			Pts, n := PointsFromBytes(b[i+1:i+42], 2)
			bvg.Circles = append(bvg.Circles, NewCircleGrad(Pts[0], Pts[1]))
			i += n + 1
		case byte('p'):
			n := binary.LittleEndian.Uint32(b[i+1 : i+5])
		//	fmt.Println(n)
			Pts, x := PointsFromBytes(b[i+5:i+5+20*int(n+1)], int(n))
			bvg.Polys = append(bvg.Polys, NewPoly(Pts...))
			i += x + 5
		case byte('b'):
			n := binary.LittleEndian.Uint32(b[i+1 : i+5])
			Pts, x := PointsFromBytes(b[i+5:i+5+20*int(n)], int(n))
			bvg.Bezs = append(bvg.Bezs, NewBez(Pts...))
			i += x + 5 
		case byte(60):
			n := binary.LittleEndian.Uint32(b[i+1 : i+5])
			Pts, x := PointsFromBytes(b[i+5:i+5+20*int(n)], int(n))
			bvg.LineStrips = append(bvg.LineStrips, NewLineStrip(Pts...))
			i += x + 5
		default:
			panic(fmt.Sprintf("Cannot identify the type of shape at %dth byte", i))
		}
	}

	return bvg, err
}

func Float64FromBytes(a []byte) float64 {
	bits := binary.LittleEndian.Uint64(a)
//	os.Stdout.Write([]byte(fmt.Sprint(b)))
//	fmt.Println(a)
	return math.Float64frombits(bits)
}

// Takes a slice of bytes and the number of points to get
// and returns an arry of points and the number of bytes
// used to construct them
func PointsFromBytes(arr []byte, l int) ([]*Point, int) {
	PArr := make([]*Point, l)
	for i := 0; i < l; i++ {
		x := Float64FromBytes(arr[20*i : 20*i+8])
		y := Float64FromBytes(arr[20*i+8 : 20*i+16])
		r, g, b, a := uint8(arr[20*i+16]), uint8(arr[20*i+17]), uint8(arr[20*i+18]), uint8(arr[20*i+19])
		PArr[i] = NewPtCol(x, y, r, g, b, a)
	//	fmt.Println(*PArr[i])
	}
	//fmt.Println(len(arr))
	return PArr, l * 20
}
