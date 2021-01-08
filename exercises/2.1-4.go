package exercises

//二进制相加
func BinarySystemAdd(a []uint8, b []uint8) []uint8 {
	l := len(a)
	c := make([]uint8, l+1)
	for i := l - 1; i >= 0; i-- {
		switch a[i] + b[i] + c[i+1] {
		case 0:
			c[i+1] = 0
		case 1:
			c[i+1] = 1
		case 2:
			c[i+1] = 0
			c[i] = 1
		case 3:
			c[i+1] = 1
			c[i] = 1
		}
	}
	return c
}
