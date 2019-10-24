package main

import (
	"fmt"
	"os"
	"strconv"
	"world-x/city"
	"world-x/grid"
)

func main() {

	setup_and_run()

}

func setup_and_run() bool {
	
	if len(os.Args) < 3 {
		fmt.Println("\nPlease run using: world-x <number_of_monsters> <small|medium>\n")
		return false
	}

	arg_monsters := os.Args[1]
	arg_file := os.Args[2]
	monsters_to_spawn, err := strconv.Atoi(arg_monsters)
	
	if  err != nil || monsters_to_spawn <= 0 || monsters_to_spawn > 12000 {		// Assume no more than 12000 monsters needed. Monsters take time to create.
		fmt.Println("\n<number_of_monsters> must be a number between 0 and 12000\n")
		return false
	}

	city_loader := city.Loader {}
	city_loader.Load(arg_file)		// File names assumed and selected by small/medium parameter

	if !(city_loader.Has_Loaded) {
		fmt.Println("\nThere was a problem loading City data\n")
		return false
	}

	fmt.Println("********************************************************")
	fmt.Println("*****************  Welcome to World X  *****************")
	fmt.Println("********************************************************")
	fmt.Printf("\n***** You have released %d Monsters, on %d Cities *****\n\n", monsters_to_spawn, city_loader.Count)
	fmt.Println("********************************************************")

	world := grid.Grid {
		Cities: city_loader.Cities,
		Game_Speed: 0,				// to match spec document set to 0. Enter > 1 to slow down. Time in milliseconds.
		Max_Monster_Moves: 10000,	// to match spec document set to 10000. 10 is good for testing.
	}

	world.Spawn_Monsters(monsters_to_spawn)
	world.Play()
	world.Print_Remaining_Cities()

	return true

}