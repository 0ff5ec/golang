package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++{
			s[y][x] = uint8((x*x+y*y)/2)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
