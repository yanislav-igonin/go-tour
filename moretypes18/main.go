package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(i * j)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
