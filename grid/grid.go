package grid

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"strings"
	"world-x/city"
	"world-x/monster"
)

type Grid struct {

	Cities []city.City
	Monsters []monster.Monster
	Max_Monster_Moves int
	Game_Speed int

}

func (e Grid) Search_City(city_name string) (*city.City, bool) {

	if city_name != "" {

		for i := range e.Cities {
			if e.Cities[i].Name == city_name {
				return &e.Cities[i], true
				break

			}

		}

	}

	empty_city := city.City {}
	return &empty_city, false

}

func (e Grid) grid_health() int {

	var health int = 0

	for _, c := range e.Cities {
		if c.Alive() { health++ }
	}

	return health

}

func get_random_number(min int, max int, not_in []int) int {

	var r_num int = 1				// TODO: think of better default return value
	var r_num_found bool = false
	var can_create int
	if (max - min) <= 0 { return r_num }
		

	for k := range not_in {			// Return default if not_in contains entire range
		if (not_in[k] >= min) && (not_in[k] <= max) { can_create++ }
	}

	if can_create >= max { return r_num }	


	time.Sleep(10 * time.Millisecond)		// Add some time to increase randomness.


	for !(r_num_found) {			// Get random number that is not in exclude array not_in
		var is_unique bool = true
		rand.Seed(time.Now().UnixNano())
		r_num = min + rand.Intn(max - min)		

		for j := range not_in {
			if not_in[j] == r_num { is_unique = false }
		}

		if is_unique { r_num_found = true }

	}

	return r_num
}

func (e *Grid) Spawn_Monsters(amount int) {

	var m_pos int = 0
	//var monster_locations []int			// Not currently used. Could be used to put monsters in unique locations.
	var total_size int = len(e.Cities)	

	for i := 0; i < amount; i++ {		
		m_pos = get_random_number(1, total_size, []int{}) - 1
		//monster_locations = append(monster_locations, m_pos)

		m := monster.Monster {
			Name: strconv.Itoa(i),
			Health: 1,
			Max_Health: 1,
			Min_Health: 1,
			Current_City: e.Cities[m_pos],
			Is_Stuck: false,
		}

		e.Monsters = append(e.Monsters, m)
	}

}

func (e Grid) Get_Available_Monsters() []int {

	var available []int

	for i := range e.Monsters {
		if e.Monsters[i].Alive() && !(e.Monsters[i].Is_Stuck) { available = append(available, i) }
	}
	
	return available

}

// func (e Grid) Alive_Monster_Count() int {		// function not required
// 
// 	var count int = 0
// 
// 	for i := range e.Monsters {
// 		if e.Monsters[i].Alive() { count++ }
// 	}
// 
// 	return count
// 
// }

func (e Grid) Alive_City_Count() int {

	var count int = 0

	for i := range e.Cities {
		if e.Cities[i].Alive() { count++ }
	}

	return count

}

func (e Grid) Print_Remaining_Cities() {

	fmt.Println("\n----------------- Remaining City Data -----------------\n")

	for i := range e.Cities {
		if e.Cities[i].Health == 0 { continue }

		fmt.Printf("%s", e.Cities[i].Name)
		
		if e.Cities[i].North_City_Data != nil && e.Cities[i].North_City_Data.Health > 0 {
			fmt.Printf(" north=%s", e.Cities[i].North_City_Data.Name)
		}
		
		if e.Cities[i].South_City_Data != nil && e.Cities[i].South_City_Data.Health > 0 {
			fmt.Printf(" south=%s", e.Cities[i].South_City_Data.Name)
		}
		
		if e.Cities[i].East_City_Data != nil && e.Cities[i].East_City_Data.Health > 0 {
			fmt.Printf(" east=%s", e.Cities[i].East_City_Data.Name)
		}
		
		if e.Cities[i].West_City_Data != nil && e.Cities[i].West_City_Data.Health > 0 {
			fmt.Printf(" west=%s", e.Cities[i].West_City_Data.Name)
		}
		
		fmt.Println("")
	}

	fmt.Println("\n-------------------------------------------------------\n")

}

func (e Grid) Play() {

	for e.grid_health() > 0 {
		if e.Game_Speed > 1 { time.Sleep(time.Duration(e.Game_Speed) * time.Millisecond) }

		if len(e.Get_Available_Monsters()) == 0 || e.Alive_City_Count() < 1 {
			break
		}

		var number_to_move int = 0
		number_to_move = get_random_number(1, len(e.Get_Available_Monsters()), []int{})		// Select a random number of monsters to move this turn

		e.move_monsters(number_to_move)
		e.fight_check()
	}
	
	fmt.Println("********************************************************")
	fmt.Printf("********** Attack Ended. %d Cities survived. ***********\n", e.Alive_City_Count())
	fmt.Println("********************************************************")

}

func (e *Grid) fight_check() {

	// Create map of City Names and Number of Monsters
	// If more than 2 monsters in a city, get City object and destroy
	// Loop monsters and kill each monster in the City

	// TODO: Think of a better way to do this

	m := make(map[string]int)

	for n := range e.Monsters {
		if !(e.Monsters[n].Alive()) { continue }

		m[e.Monsters[n].Current_City.Name] = (m[e.Monsters[n].Current_City.Name] + 1)

		if m[e.Monsters[n].Current_City.Name] >= 2 {
			x_city, found := e.Search_City(e.Monsters[n].Current_City.Name)
			if !(found) { continue }

			x_city.Health = 0
			fmt.Printf("\n%s has been destroyed by the following monsters: \n", x_city.Name)			
			var monster_list string

			for i := range e.Monsters {
				if !(e.Monsters[i].Current_City.Name == x_city.Name) { continue }
				e.Monsters[i].Health = 0
				monster_list = monster_list + fmt.Sprintf(" %s,", e.Monsters[i].Name)
			}

			monster_list = strings.TrimRight(monster_list, ",")
			monster_list = strings.TrimSpace(monster_list)

			fmt.Printf(monster_list)
			fmt.Printf("\n\n")
		}
	}

	//fmt.Println(m)
}

func (e *Grid) move_monsters(number_to_move int) {

	var m_pos int = 0
	var monsters_moved []int

	for i := 0; i < number_to_move; i++ {
		var available_monster []int = e.Get_Available_Monsters()
		var r_num int = get_random_number(1, len(available_monster), monsters_moved) - 1	// Get a random available monster
		m_pos = available_monster[r_num]
		monsters_moved = append(monsters_moved, m_pos)
		
		if  e.Monsters[m_pos].Moves >= e.Max_Monster_Moves {
			e.Monsters[m_pos].Health = 0
			continue
		}

		e.Monsters[m_pos].Moves++
		var move_direction int = get_random_number(1, 4, []int{})	// Pick a random direction to move in 1=North, 2=East, 3=South, 4=West

		// Check if monster is stuck and can't move anywhere
		if !(e.Monsters[m_pos].Can_Move_North()) && !(e.Monsters[m_pos].Can_Move_East()) && !(e.Monsters[m_pos].Can_Move_South()) && !(e.Monsters[m_pos].Can_Move_West()) {
			e.Monsters[m_pos].Is_Stuck = true
			continue
		}

		switch move_direction {
			case 1:		// North
			if e.Monsters[m_pos].Can_Move_North() {					
				e.Monsters[m_pos].Current_City = *e.Monsters[m_pos].Current_City.North_City_Data
			}

			case 2:		// East
			if e.Monsters[m_pos].Can_Move_East() {
				e.Monsters[m_pos].Current_City = *e.Monsters[m_pos].Current_City.East_City_Data
			}

			case 3:		// South
			if e.Monsters[m_pos].Can_Move_South() {
				e.Monsters[m_pos].Current_City = *e.Monsters[m_pos].Current_City.South_City_Data
			}

			case 4:		// West
			if e.Monsters[m_pos].Can_Move_West() {
				e.Monsters[m_pos].Current_City = *e.Monsters[m_pos].Current_City.West_City_Data
			}
		}
	}

}
