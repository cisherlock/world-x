package monster

import (
	"world-x/city"
)

type Monster struct {

	Health int
	Max_Health int		// Added min and max health for potential future expansion
	Min_Health int
	Name string
	Current_City city.City
	Is_Stuck bool
	Moves int
	
}

func (e Monster) Alive() bool {

	fmt.Println("test branch")
	return (e.Health >= e.Min_Health)

	// test comment came after

}

func (e Monster) Can_Move_North() bool {

	if !(e.Alive()) || e.Is_Stuck { return false }
	if e.Current_City.North_City_Data == nil { return false }
	return e.Current_City.North_City_Data.Alive()

}

func (e Monster) Can_Move_East() bool {

	if !(e.Alive()) || e.Is_Stuck { return false }
	if e.Current_City.East_City_Data == nil { return false }
	return e.Current_City.East_City_Data.Alive()

}

func (e Monster) Can_Move_South() bool {

	if !(e.Alive()) || e.Is_Stuck { return false }
	if e.Current_City.South_City_Data == nil { return false }
	return e.Current_City.South_City_Data.Alive()

}

func (e Monster) Can_Move_West() bool {

	if !(e.Alive()) || e.Is_Stuck { return false }
	if e.Current_City.West_City_Data == nil { return false }
	return e.Current_City.West_City_Data.Alive()

}