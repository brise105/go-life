package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "math"
  "math/rand"
)

const ALIVE int = 1
const DEAD int = 0

var game Life
var input Vars

//structures
type Life struct {	
	NullGameTable bool

	ArrRowSize int
	ArrColumnSize int

	GameTable [][]int
	CurrentArray [][]int
	NextArray [][]int

	GenCount int
	Living int

	State int //flow of control for output 
}

type Vars struct {
	GameMode rune
	OutputMode rune

	OGRowSize float64
	OGColumnSize float64

	Density float64
}



//objects
func WelcomeFunction() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Conway's Game of Life!!")
	fmt.Print("Press Enter to continue...")
	reader.ReadString('\n')	

	fmt.Println("\nRULES:")
	fmt.Println("1--Any live cell with fewer than two live neighbours dies.")
	fmt.Println("2--Any live cell with more than three live neighbours dies.")
	fmt.Println("3--Any live cell with two or three live neighbours lives on to the next generation.")
	fmt.Println("4--Any dead cell with exactly three live neighbours will come to life in the next generation.")
	fmt.Print("Press Enter to continue...")
	reader.ReadString('\n')	

	validate1 := false
	for validate1 == false {
		reader.Reset(os.Stdin)
		fmt.Println("\nWhat game mode would you like to play in: classic, mirror, or doughnut? (ex: 'c', 'm', 'd')")
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		switch char {
			case 'c':
				input.GameMode = 'c'
				//fmt.Printf("%c\n", input) //debug statement
				validate1 = true
				break
			case 'm':
				input.GameMode = 'm'
				validate1 = true
				break
			case 'd':
				input.GameMode = 'd'
				validate1 = true
				break
			default:
				break
		}
	}

	validate2 := false
	for validate2 == false {
		reader.Reset(os.Stdin)
		fmt.Println("\nChoose your output style:")
		fmt.Println("(b) buffer output to console")
		fmt.Println("(w) wait for user Enter press")

		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		switch char {
			case 'b':
				input.OutputMode = 'b'
				validate2 = true
				break
			case 'w':
				input.OutputMode = 'w'
				validate2 = true
				break
			default:
				break
		}
	}

	validate3 := false
	for validate3 == false {
		reader.Reset(os.Stdin)
		fmt.Print("\nEnter the number of desired rows: ")
		height, err := reader.ReadString('\n') //read in string
		if err != nil {
			fmt.Println(err)
		}
		heightStr := height[0:len(height)-1] //trim string for newlines/returncarriers
		heightInt, err := strconv.ParseInt(heightStr, 10, 32) //parse to int
		if heightInt >  0 {
			h := int(heightInt) //type conversion to regular uint
			input.OGRowSize = float64(h)
			game.ArrRowSize = h+2
			validate3 = true
		}
	}

	validate4 := false
	for validate4 == false {
		reader.Reset(os.Stdin)
		fmt.Print("\nEnter the number of desired columns: ")
		width, err := reader.ReadString('\n') //read in string
		if err != nil {
			fmt.Println(err)
		}
		widthStr := width[0:len(width)-1] //trim string for newlines/returncarriers
		widthInt, err := strconv.ParseInt(widthStr, 10, 32) //parse to int
		if widthInt >  0 {
			w := int(widthInt) //type conversion to regular uint
			input.OGColumnSize = float64(w)
			game.ArrColumnSize = w+2
			validate4 = true
		}
	}

	validate5 := false
	for validate5 == false {
		reader.Reset(os.Stdin)
		fmt.Print("\nEnter a decimal value greater than 0 AND less than or equal to 1: ")
		dens, err := reader.ReadString('\n') //read in string
		if err != nil {
			fmt.Println(err)
		}
		densityStr := dens[0:len(dens)-1] //trim string for newlines/returncarriers
		densityFloat, err := strconv.ParseFloat(densityStr, 64) //parse to float
		if densityFloat >  0 && densityFloat <= 1 {
			d := float64(densityFloat) //type conversion to float32
			input.Density = d
			validate5 = true
		}
	}					
}

func SetArrays() {
	game.GameTable = make([][]int, game.ArrRowSize)
	for i := 0; i < game.ArrRowSize; i++ {
		game.GameTable[i] = make([]int, game.ArrColumnSize)
	}

	game.CurrentArray = make([][]int, game.ArrRowSize)
	for j := 0; j < game.ArrRowSize; j++ {
		game.CurrentArray[j] = make([]int, game.ArrColumnSize)
	}

	game.NextArray = make([][]int, game.ArrRowSize)
	for k := 0; k < game.ArrRowSize; k++ {
		game.NextArray[k] = make([]int, game.ArrColumnSize)
	}
}

func FillArrayOG(arr [][]int) {

	Y := 0
	X := 0
	game.Living = int(math.Ceil((input.OGRowSize*input.OGColumnSize)*(input.Density)))

	for i := 1; i < game.ArrRowSize-1; i++ {
		for j := 1; j < game.ArrColumnSize-1; j++ {
			arr[i][j] = DEAD
		}
	}

	for game.Living > 0 {
		
		Y = rand.Intn(int(input.OGRowSize+1))
		X = rand.Intn(int(input.OGColumnSize+1))
		if arr[Y][X] != ALIVE {
			arr[Y][X] = ALIVE
			game.Living--
		}
	}	
}

func FillArrayBoundary(arr [][]int) {
	if input.GameMode == 'c' {
		for i := 0; i < game.ArrRowSize; i++ {
			for j := 0; j < game.ArrColumnSize; j++ {
				if i == 0 || i == game.ArrRowSize-1 {
					arr[i][j] = DEAD
				}
				if j == 0 || j == game.ArrColumnSize-1 {
					arr[i][j] = ALIVE
				}
			}
		}
	}
	if input.GameMode == 'm' || input.GameMode == 'd' {
		for i := 1; i < game.ArrRowSize-1; i++ {
			for j := 1; j < game.ArrColumnSize-1; j++ { 
				if i == 1 && j == 1 { //(1)top-left
					if input.GameMode == 'm' {
						arr[i-1][j-1] = arr[i][j]; //diagonal left top = current index
						arr[i-1][j] = arr[i][j]; //above = current index
						arr[i][j-1] = arr[i][j]; //leftside = current index						
					}
					if input.GameMode == 'd' {
						arr[i-1][j-1] = arr[game.ArrRowSize-2][game.ArrColumnSize-2]; //diagonal left top = top-right						
					}
				}
				if i == 1 && j > 1 && j < game.ArrColumnSize-2 { //(2)top-middle
					if input.GameMode == 'm' {
						arr[i-1][j] = arr[i][j]; //above = current index 
					}
					if input.GameMode == 'd' {
						arr[i-1][j-1] = arr[game.ArrRowSize-2][j-1]; //diagonal left top = diagonal left bottom  
						arr[i-1][j] = arr[game.ArrRowSize-2][j]; //above = bottom
						arr[i-1][j+1] = arr[game.ArrRowSize-2][j+1]; //diagonal right top = diagonal right bottom
					}
				}
				if i == 1 && j == game.ArrColumnSize-2 { //(3)top-middle
					if input.GameMode == 'm' {
						arr[i-1][j+1] = arr[i][j]; //diagonal right top = current index
						arr[i-1][j] = arr[i][j]; //above = current index
						arr[i][j+1] = arr[i][j]; //rightside = current index
					}
					if input.GameMode == 'd' {
						arr[i-1][j+1] = arr[game.ArrRowSize-2][1]; //diagonal right top = bottom-left
					}
				}
				if i > 1 && i < game.ArrRowSize-2 && j == 1 { //(4)top-middle
					if input.GameMode == 'm' {
						arr[i][j-1] = arr[i][j]; //leftside = current index 
					}
					if input.GameMode == 'd' {
						arr[i+1][j-1] = arr[i+1][game.ArrColumnSize-2]; //diagonal left top = diagonal right top 
						arr[i][j-1] = arr[i][game.ArrColumnSize-2]; //leftside = rightside
						arr[i-1][j-1] = arr[i-1][game.ArrColumnSize-2]; //diagonal left bottom = diagonal right bottom						
					}
				}	
				if i > 1 && i < game.ArrRowSize-2 && j == game.ArrColumnSize-2 { //(5)top-middle
					if input.GameMode == 'm' {
						arr[i][j+1] = arr[i][j]; //rightside = current index 
					}
					if input.GameMode == 'd' {
						arr[i-1][j+1] = arr[i-1][1]; //diagonal right top = diagonal left top 
						arr[i][j+1] = arr[i][1]; //rightside = leftside
						arr[i+1][j+1] = arr[i+1][1]; //diagonal right bottom = diagonal left bottom
					}
				}
				if i == game.ArrRowSize-2 && j == 1 { //(6)top-middle
					if input.GameMode == 'm' {
						arr[i+1][j-1] = arr[i][j]; //diagonal left bottom = current index
						arr[i][j-1] = arr[i][j]; //leftside = current index
						arr[i+1][j] = arr[i][j]; //below = current index						
					}
					if input.GameMode == 'd' {
						arr[i-1][j-1] = arr[game.ArrRowSize-2][game.ArrColumnSize-2]; //diagonal left bottom = bottom-right						
					}
				}
				if i == game.ArrRowSize-2 && j > 1 && j < game.ArrColumnSize { //(7)top-middle
					if input.GameMode == 'm' {
						arr[i+1][j] = arr[i][j]; //below = current index 
					}
					if input.GameMode == 'd' {
						arr[i+1][j-1] = arr[1][j-1]; //diagonal left bottom = diagonal right top  
						arr[i+1][j] = arr[1][j]; //below = above
						arr[i+1][j+1] = arr[1][j+1]; //diagonal right bottom = diagonal left top						
					}
				}
				if i == game.ArrRowSize-2 && j == game.ArrColumnSize-2 { //(8)top-middle
					if input.GameMode == 'm' {
						arr[i+1][j+1] = arr[i][j]; //diagonal right top = current index
						arr[i][j+1] = arr[i][j]; //rightside = current index
						arr[i+1][j] = arr[i][j]; //below = current index						
					}
					if input.GameMode == 'd' {
						arr[i+1][j+1] = arr[1][1]; //diagonal right bottom = top-left						
					}
				}																																	
			}
		}
	}
}

func CopyFunction(arr1 [][]int, arr2 [][]int) {
	for i := 0; i < game.ArrRowSize; i++ {
		for j := 0; j < game.ArrColumnSize; j++ {
			if arr1[i][j] == ALIVE {
				arr2[i][j] = ALIVE
			}
			if arr1[i][j] == DEAD {
				arr2[i][j] = DEAD
			}
		}
	}
}

func CountLiving(arr [][]int) {
	count := 0
	for i := 0; i < game.ArrRowSize; i++ {
		for j := 0; j < game.ArrColumnSize; j++ {
			if arr[i][j] == ALIVE {
				count += 1
			}	
		}
	}
	game.Living = count		
}

func StableGameBool(arr1 [][]int, arr2 [][]int) bool {
	for i := 0; i < game.ArrRowSize; i++ {
		for j := 0; j < game.ArrColumnSize; j++ {
			if arr1[i][j] != arr2[i][j] {
				game.NullGameTable = false
				return false
			} else {
				game.NullGameTable = true
			}
		}
	}
	return true
}

func PrintOGArray(arr [][]int) {
	fmt.Println("Generation: ", game.GenCount)
	for i := 1; i < game.ArrRowSize-1; i++ {
		for j := 1; j < game.ArrColumnSize-1; j++ {
			if arr[i][j] == ALIVE {
				fmt.Print("X")
			} else if arr[i][j] == DEAD {
				fmt.Print("-")
			}
		}
		fmt.Println("\n")
	}
	game.GenCount++
}

func PrintCurrentArray(arr [][]int) {
	fmt.Println("Generation: ", game.GenCount)
	for i := 1; i < game.ArrRowSize-1; i++ {
		for j := 1; j < game.ArrColumnSize-1; j++ {
			if arr[i][j] == ALIVE {
				fmt.Print("X")
			} else if arr[i][j] == DEAD {
				fmt.Print("-")
			}
		}
		fmt.Println("\n")
	}
	game.GenCount++
}

func PrintStyled() {
	if input.OutputMode == 'b' {
		if game.GenCount < 1 {
			PrintOGArray(game.GameTable)
		} else {
			PrintCurrentArray(game.CurrentArray)
		}
	}
	if input.OutputMode == 'w' {
		reader := bufio.NewReader(os.Stdin)
		if game.GenCount < 1 {
			PrintOGArray(game.GameTable)
			fmt.Println("Press Enter to continue...")
			reader.ReadString('\n')			
		} else {
			PrintCurrentArray(game.CurrentArray)
			fmt.Println("Press Enter to continue...")
			reader.ReadString('\n')			
		}
	}
}

func AutomateLife() {
	for game.NullGameTable == false {
		for i := 1; i < game.ArrRowSize-1; i++ {
			for j := 1; j < game.ArrColumnSize-1; j++ {
                if (game.CurrentArray[i-1][j-1])+(game.CurrentArray[i-1][j])+
                    (game.CurrentArray[i-1][j+1])+(game.CurrentArray[i][j-1])+
                    (game.CurrentArray[i][j+1])+(game.CurrentArray[i+1][j-1])+
                    (game.CurrentArray[i+1][j])+(game.CurrentArray[i+1][j+1]) <= 1 {
                    game.NextArray[i][j] = DEAD
                }
                if (game.CurrentArray[i-1][j-1])+(game.CurrentArray[i-1][j])+
                    (game.CurrentArray[i-1][j+1])+(game.CurrentArray[i][j-1])+
                    (game.CurrentArray[i][j+1])+(game.CurrentArray[i+1][j-1])+
                    (game.CurrentArray[i+1][j])+(game.CurrentArray[i+1][j+1]) == 2 {
                    game.NextArray[i][j] = game.CurrentArray[i][j]
                }
                if (game.CurrentArray[i-1][j-1])+(game.CurrentArray[i-1][j])+
                    (game.CurrentArray[i-1][j+1])+(game.CurrentArray[i][j-1])+
                    (game.CurrentArray[i][j+1])+(game.CurrentArray[i+1][j-1])+
                    (game.CurrentArray[i+1][j])+(game.CurrentArray[i+1][j+1]) == 3 {
                    game.NextArray[i][j] = ALIVE
                }
                if (game.CurrentArray[i-1][j-1])+(game.CurrentArray[i-1][j])+
                    (game.CurrentArray[i-1][j+1])+(game.CurrentArray[i][j-1])+
                    (game.CurrentArray[i][j+1])+(game.CurrentArray[i+1][j-1])+
                    (game.CurrentArray[i+1][j])+(game.CurrentArray[i+1][j+1]) >= 4 {
                    game.NextArray[i][j] = DEAD
                }	                	                	                				
			}
		}
		FillArrayBoundary(game.NextArray)
		if StableGameBool(game.CurrentArray, game.NextArray) == true {
			fmt.Println("The generations have stablized. Ending program.")
			break
		}			
		CopyFunction(game.NextArray, game.CurrentArray)
		//CountLiving(game.CurrentArray)
		PrintStyled()	
	}
}

func PrepGame() {
	WelcomeFunction()
	SetArrays()
	FillArrayOG(game.GameTable)
	FillArrayBoundary(game.GameTable)
	CopyFunction(game.GameTable, game.CurrentArray)
	//CountLiving(game.CurrentArray)
	PrintStyled()
}

func main() {
	PrepGame()
	AutomateLife()
}
