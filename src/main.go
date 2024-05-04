package main

import (
	"fmt"
	"math"
	"time"
)

func memset[T any](arr []T, val T) {
	if len(arr) == 0 {
		return
	}

	arr[0] = val
	for ptr := 1; ptr < len(arr); ptr *= 2 {
		copy(arr[ptr:], arr[:ptr])
	}
}

func main() {
	rot8_x, rot8_z := 0.0, 0.0
	z := make([]float64, 1760)
	b := make([]byte, 1760)
	lum_str := ".,-~:;=!*#$@"

	fmt.Print("\x1b[2J")

	for {
		memset(b, 32)
		memset(z, 0)
		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(rot8_x)
				f := math.Sin(j)
				g := math.Cos(rot8_x)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(rot8_z)
				n := math.Sin(rot8_z)
				t := c*h*g - f*e
				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				o := int(x + 80*y)
				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))
				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = lum_str[N]
					} else {
						b[o] = lum_str[0]
					}
				}
			}
		}
		fmt.Print("\x1b[H")
		for k := 0; k < 1761; k++ {
			if k%80 != 0 {
				fmt.Printf("%c", b[k])
			} else {
				fmt.Print("\n")
			}
			rot8_x += 0.00004
			rot8_z += 0.00002
		}
		time.Sleep(35 * time.Millisecond)
	}
}
