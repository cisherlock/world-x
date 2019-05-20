package city

type City struct {

	Health int
	Max_Health int		// Added min and max health for potential future expansion
	Min_Health int
	Name string
	North string
	East string
	South string
	West string
	North_City_Data *City
	East_City_Data *City
	South_City_Data *City
	West_City_Data *City
	Loaded_Linked_Data bool

}

func (e City) Alive() bool {

	return (e.Health >= e.Min_Health)
	
}