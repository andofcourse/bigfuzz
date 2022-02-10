package broken

import (
	"testing"
)

func FuzzBroken(f *testing.F) {

    f.Fuzz(func(t *testing.T, str string, bytes []byte, i int, i8 int8, i16 int16, r rune, i64 int64, ui uint, by byte, ui16 uint16, ui32 uint32, ui64 uint64, f32 float32, f64 float64, b bool) {
        Broken(str, bytes, i , i8 , i16 , r , i64 , ui , by , ui16 , ui32 , ui64 , f32 , f64 , b )
    })
}