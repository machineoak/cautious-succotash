package main

import (
	"fmt"
)

type Tile struct {
	value int
}

func create3darr(x, y, z int) [][][]*Tile {
	result := make([][][]*Tile, x)
	for i := 0; i < x; i++ {
		result[i] = make([][]*Tile, y)
		for j := 0; j < y; j++ {
			result[i][j] = make([]*Tile, z)
			for k := 0; k < z; k++ {
				result[i][j][k] = new(Tile)
				result[i][j][k].value = i + j + k
			}
		}
	}
	return result
}

func main() {
	X := 3
	Y := 4
	Z := 5

	mat := create3darr(X, Y, Z)
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			for k := 0; k < Z; k++ {
				fmt.Printf("%d ", mat[i][j][k].value)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
