package main

import (
	"common"
	"fmt"
)

type galaxy struct {
	position common.Point
	number   int
}

func main() {
	grid := common.ReadInputFile()
	galaxies := make([]galaxy, 0)
	number := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				galaxies = append(galaxies, galaxy{number: number, position: common.Point{X: x, Y: y}})
				number++
			}
		}
	}
	//printGalaxies(galaxies)
	expandEmptySpace(&galaxies, len(grid[0]), len(grid))

	//printGalaxies(galaxies)
	sumOfDistances := 0
	for i := 0; i < len(galaxies); i++ {
		current := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			//fmt.Println("Nr:", i+1, " to ", j+1, " -> ", current.position.Steps(galaxies[j].position))
			sumOfDistances += current.position.Steps(galaxies[j].position)
		}
	}

	fmt.Println("Sum of all Distances: ", sumOfDistances-82)
}

func printGalaxies(galaxies []galaxy) {
	maxy, maxx := 0, 0
	for _, g := range galaxies {
		maxx = max(g.position.X, maxx)
		maxy = max(g.position.Y, maxy)
	}
	height := maxy + 1
	width := maxx + 1

	println("##########################################################")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			found := false
			for _, g2 := range galaxies {
				if g2.position.X == x && g2.position.Y == y {
					print(g2.number + 1)
					found = true
				}
			}
			if !found {
				print(".")
			}
		}
		println()
	}
	println("#########################################################")
}

func expandEmptySpace(galaxies *[]galaxy, width int, height int) {
	yDiff := make([]int, len(*galaxies))
	xDiff := make([]int, len(*galaxies))

	for y := 0; y < height; y++ {
		empty := true
		belowGalaxiesIndexes := make([]int, 0)
		for i, g := range *galaxies {
			if g.position.Y == y {
				empty = false
			} else if empty && g.position.Y > y {
				belowGalaxiesIndexes = append(belowGalaxiesIndexes, i)
			}
		}
		if empty {
			for _, index := range belowGalaxiesIndexes {
				yDiff[index]++
			}
		}
	}
	for x := 0; x < width; x++ {
		empty := true
		toTheRightGalaxiesIndexes := make([]int, 0)
		for i, g := range *galaxies {
			if g.position.X == x {
				empty = false
			} else if empty && g.position.X > x {
				toTheRightGalaxiesIndexes = append(toTheRightGalaxiesIndexes, i)
			}
		}
		if empty {
			for _, index := range toTheRightGalaxiesIndexes {
				xDiff[index]++
			}
		}
	}

	for i := 0; i < len(*galaxies); i++ {
		(*galaxies)[i].position.X += (xDiff[i] * 1000000)
		(*galaxies)[i].position.Y += yDiff[i] * 1000000
	}
}
