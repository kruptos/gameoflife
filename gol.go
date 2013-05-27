package main
 
import (
	"fmt"
)
 
func checkNeighbors(x int, y int, space[][] int) int {
	neighbors := 0
	for i := x-1; i < x + 2; i++ {
		for j := y-1; j < y + 2; j++ {
			if i < 0 && j < 0 {
			} else if i < 0 && j >= 0 {
			} else if i > 9 && j >= 0 {
			} else if j < 0 && i >= 0 {
			} else if j > 9 && i >= 0 {
			} else {
				if space[i][j] == 1 {
					if i == x && j == y{
					} else {
						neighbors += 1	
					}
				}
			}
		}
	}
	return neighbors
}

func updateSpace(space[][] int) [][]int{
	newSpace := make([][]int, 10)
	for k := 0; k < 10; k++ {
		newSpace[k] = make([]int, 10)
		for l := 0; l < 10; l++ {
			newSpace[k][l] = int(0)
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			currentNeighbors := 0 // init variable
			currentNeighbors = checkNeighbors(i,j,space)
			if space[i][j] == 1 { // if cell is alive
				if currentNeighbors < 2 {
					newSpace[i][j] = 0 // kill cell
					//fmt.Println("Cell ", "(", i, ",",j,") Killed due to no friends. He had", currentNeighbors, "friends.")
				} else if currentNeighbors == 2 || currentNeighbors == 3 {
					newSpace[i][j] = 1 
					//fmt.Println("Cell ", "(", i, ",",j,") kept alive.")
				} else if currentNeighbors >= 4 {
					newSpace[i][j] = 0 // kill the cell due to overcrowding
					//fmt.Println("Cell ", "(", i, ",",j,") Killed due to Overcrowding.")
				}
			} else if space[i][j] == 0 { // if cell is dead
				if currentNeighbors == 3 {
					newSpace[i][j] = 1 // bring cell to life
					//fmt.Println("Cell ", "(", i, ",",j,") brought to life.")
				}
			}
		}
	}
	return(newSpace)
}

func printSpace(space[][] int) {
	// prints board
	for k := 0; k < 10; k++ {
		fmt.Println(space[k])	
	}
}

func main() {
	space := make([][]int, 10)
	//initialize zeros in array
	for i := 0; i < 10; i++ {
		space[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			space[i][j] = int(0)
		}
	}
	//initialize starting alive cells
	space[5][5] = 1
	space[5][6] = 1
	space[6][5] = 1
	space[6][6] = 1
	space[2][3] = 1
	space[8][9] = 1
	space[4][5] = 1
	space[6][2] = 1
	space[2][4] = 1
	space[7][9] = 1
	space[7][8] = 1
	space[0][3] = 1
	space[2][7] = 1

	for r := 0; r < 100; r++ {
		fmt.Println("Iteration:", r, "-----------------------\b \b")
		printSpace(space)
		space = updateSpace(space)
	}
}