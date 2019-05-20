package city

import (
	"os"
	"bufio"
	"regexp"
	"strings"	
)

type Loader struct {

	Has_Loaded bool
	Count int
	Cities []City
	
}

func (e Loader) Search_City(city_name string) (*City, bool) {

	if city_name != "" {
		for i := range e.Cities {
			if e.Cities[i].Name == city_name {
				return &e.Cities[i], true
				break
			}
			
		}

	}

	empty_city := City {}
	return &empty_city, false

}

func (e *Loader) Load(data_file string) bool {

	var file *os.File
	var err error

	switch data_file {
		case "small":
		file, err = os.Open("resource/world_map_small.txt")			// Assuming file names

		case "medium":
		file, err = os.Open("resource/world_map_medium.txt")		// Assuming file names

		default:
		return false
	}

	if err != nil { return false }

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txt_lines []string

	for scanner.Scan() {
		txt_lines = append(txt_lines, scanner.Text())
	}

	file.Close()

	for _, line := range txt_lines {
		if strings.TrimSpace(line) == "" { continue }

		name_rex, _ := regexp.Compile("^[A-Za-z\\-]+")				// File format is assumed.
		north_rex, _ := regexp.Compile("\\snorth=[A-Za-z\\-]+")
		east_rex, _ := regexp.Compile("\\seast=[A-Za-z\\-]+")
		south_rex, _ := regexp.Compile("\\ssouth=[A-Za-z\\-]+")
		west_rex, _ := regexp.Compile("\\swest=[A-Za-z\\-]+")

		c_name := strings.TrimSpace(name_rex.FindString(line))
		c_north := strings.TrimSpace(north_rex.FindString(line))
		c_east := strings.TrimSpace(east_rex.FindString(line))
		c_south := strings.TrimSpace(south_rex.FindString(line))
		c_west := strings.TrimSpace(west_rex.FindString(line))

		// e.g. c_north = "north=city"
		// If statement will change so c_north = "city"

		if c_north != "" {
			c_north = strings.Split(c_north, "=")[1]
		}
		if c_east != "" {
			c_east = strings.Split(c_east, "=")[1]
		}
		if c_south != "" {
			c_south = strings.Split(c_south, "=")[1]
		}
		if c_west != "" {
			c_west = strings.Split(c_west, "=")[1]
		}

		c := City {
			Name: c_name,
			Health: 1,
			Max_Health: 1,
			Min_Health: 1,
			North: c_north,
			East: c_east,
			South: c_south,
			West: c_west,
		}

		e.Cities = append(e.Cities, c)
		e.Count++
	}


	// Load linked data
	for i := range e.Cities {
		if e.Cities[i].Loaded_Linked_Data { continue }

		x_city, found := e.Search_City(e.Cities[i].North)

		if found {
			e.Cities[i].North_City_Data = x_city
		}
		
		x_city, found = e.Search_City(e.Cities[i].East)

		if found {
			e.Cities[i].East_City_Data = x_city
		}
		
		x_city, found = e.Search_City(e.Cities[i].South)

		if found {
			e.Cities[i].South_City_Data = x_city
		}

		x_city, found = e.Search_City(e.Cities[i].West)

		if found {
			e.Cities[i].West_City_Data = x_city
		}

		e.Cities[i].Loaded_Linked_Data = true
	}

	e.Has_Loaded = true
	return e.Has_Loaded

}