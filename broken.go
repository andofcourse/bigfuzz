package broken

func Broken(str string, bytes []byte, i int, i8 int8, i16 int16, r rune, i64 int64, ui uint, by byte, ui16 uint16, ui32 uint32, ui64 uint64, f32 float32, f64 float64, b bool) bool {
	if len(str) > 3 &&	str[0] == 'F' && str[1] == 'U' && str[2] == 'Z' && str[3] == 'Z' {
		return true
	}

	if len(bytes) > 1 && bytes[0] == 42 && bytes[1] == 33 {
		return true
	}

	if i > 500 && i % 5 == 0 {
		return true
	}

	if i8 < 42 {
		return true
	}

	if i16 > 42 {
		return true
	}

	if r == '🦍' {
		return true
	}

	if i64 > 100000 {
		return true
	}

	if ui < 99 {
		return true
	}

	if by == 1 {
		return true
	}

	if ui16 != 33 {
		return true
	}

	if ui32 >= 999 {
		return true
	}

	if ui64 == 1234 {
		return true
	}

	if f32 == 3.14159 {
		return true
	}

	if f64 <= 139.3912 {
		return true
	}

	if b == false {
		return true
	}
	
	return false
}