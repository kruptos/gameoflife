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


	var userInputX int
	var userInputY int
	var userInputEnd string
	var userInputIterations int
	// User input for Starting Alive cells

	userInputEnd = "y"
	userInputX = 0
	userInputY = 0

	for userInputEnd != "n" {

		fmt.Println("Enter coords of starting alive cells (between 0 and 9)")

		for {
			fmt.Println("Enter X:")
				_, errX := fmt.Scanf("%d", &userInputX)
				if errX != nil {
					panic(errX)
				}
			if userInputX < 0 || userInputX > 9 {
				//values should not be in this range
				fmt.Println("Coordinate inputed is less than 0 or greater than 9. Try again.")
			} else {
				//values should be in this range
				// correct value entered breaking out of loop to head to Y value input
				break
			}					
		}		
		
		for {
			fmt.Println("Enter Y:")
				_, errY := fmt.Scanf("%d", &userInputY)
				if errY != nil {
					panic(errY)
				}
			if userInputY < 0 || userInputY > 9 {
				//values should not be in this range
				fmt.Println("Coordinate inputed is less than 0 or greater than 9. Try again.")
			} else {
				//values should be in this range
				// correct value entered breaking out of loop to head to Y value input
				break
			}					
		}
		
		space[userInputX][userInputY] = 1
		fmt.Println("Cell (",userInputX,",",userInputY,") is ALIVE!!!!")
		fmt.Println("Do you want to add more cells? (y or n):")

		_, err := fmt.Scanf("%s", &userInputEnd)
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("Successfully exited.\n","How many times should the simulation iterate?")
	fmt.Println("Enter number > 0:")

	_, err := fmt.Scanf("%d", &userInputIterations)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting simulation; iterating through", userInputIterations, "times.")

	for r := 0; r < userInputIterations; r++ {
		fmt.Println("Iteration:", r, "-----------------------\b \b")
		printSpace(space)
		space = updateSpace(space)
	}
}