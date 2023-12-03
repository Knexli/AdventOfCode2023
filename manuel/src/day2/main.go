package main

import (
	"common"
	"strconv"
	"strings"
)

type game struct {
	red   []int
	green []int
	blue  []int
	id    int
}

func main() {

	maxGame := game{
		red:   []int{12},
		green: []int{13},
		blue:  []int{14},
		id:    -1,
	}

	lines := common.ReadInputFile()

	games := getGames(lines)

	sum := getSumOfPossibleGameIds(maxGame, games)
	println(sum)
	println("++++++++++++++++ Part Two ++++++++++++++++++++")

	sum = getSumOfPowersOfMinGames(games)
	println(sum)
}

func getSumOfPowersOfMinGames(games []game) int {
	sum := 0
	for _, g := range games {
		min := struct {
			red   int
			blue  int
			green int
		}{
			red:   0,
			blue:  0,
			green: 0,
		}
		for i := 0; i < len(g.green); i++ {
			min.red = max(min.red, g.red[i])
			min.blue = max(min.blue, g.blue[i])
			min.green = max(min.green, g.green[i])
		}
		sum += min.green * min.red * min.blue
	}

	return sum
}

func getSumOfPossibleGameIds(maxGame game, games []game) int {
	sum := 0
	for _, g := range games {
		possible := true
		for i := 0; i < len(g.red); i++ {
			if g.red[i] > maxGame.red[0] || g.blue[i] > maxGame.blue[0] || g.green[i] > maxGame.green[0] {
				//fmt.Printf("Game: %d r: %d b: %d g: %d\n", g.id, g.red[i], g.blue[i], g.green[i])
				possible = false
			}
		}
		if possible {
			sum += g.id
		}

	}
	return sum
}

func getGames(games []string) []game {
	result := make([]game, 0)
	for _, s := range games {
		result = append(result, getGame(s))
	}
	return result
}

func getGame(s string) game {
	chunks := strings.Split(s, ":")
	number, err := strconv.Atoi(strings.Split(chunks[0], " ")[1])

	if err != nil {
		panic(err)
	}

	game := game{
		red:   make([]int, 0),
		green: make([]int, 0),
		blue:  make([]int, 0),
		id:    number,
	}

	turns := strings.Split(chunks[1], ";")
	for _, turn := range turns {
		cubes := strings.Split(turn, ",")
		for _, cube := range cubes {
			parts := strings.Split(strings.Trim(cube, " "), " ")
			color := parts[1]
			count, err := strconv.Atoi(parts[0])

			if err != nil {
				panic(err)
			}

			switch color {
			case "red":
				game.green = append(game.green, 0)
				game.blue = append(game.blue, 0)
				game.red = append(game.red, count)
			case "blue":
				game.red = append(game.red, 0)
				game.green = append(game.green, 0)
				game.blue = append(game.blue, count)
			case "green":
				game.green = append(game.green, count)
				game.red = append(game.red, 0)
				game.blue = append(game.blue, 0)
			default:
				panic("unknown color: " + color)
			}

		}

	}

	return game
}
