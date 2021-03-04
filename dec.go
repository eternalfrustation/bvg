package bvg

import (
	"encoding/binary"
)

// Decodes the given bytes and returns a pointer to 
// an Bvg struct, You need to set the writer for this new
// struct
func Decode(b []byte) *Bvg {
	bvg := new(Bvg)
	switch int8(b[0]) {
	case int8(108):

	}
	return bvg
}
