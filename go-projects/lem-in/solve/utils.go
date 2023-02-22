package lemin

import (
	"strconv"
	"strings"
)



// Checks if the input represents the number of ants in the colony
func isNrOfAnts(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false

		}

	}
	return true
}


func findStart(start string) bool {
	for i, r := range rooms {
		if r.name == start {
			StartRoom = &rooms[i]
			return true
		}
	}
	return false
}

func findEnd(end string) bool {
	for i, r := range rooms {
		if r.name == end {
			EndRoom = &rooms[i]
			return true
		}
	}
	return false
}

func isLink(line string) bool {
	if line[0] == '#' || line[0] == 'L' || line[0] == ' ' {
		return false
	}

	temp := strings.Split(line, "-")

	for _, r := range temp {
		ok := false
		for _, r2 := range rooms { //check if room exist

			if r == r2.name {
				ok = true
				break
			}
		}
		if !ok && len(temp) == 2 {
			//fmt.Println("Link with bad rooms:", line)
			return false
		}
	}
	if len(temp) == 2 {
		for i, room := range rooms {
			if room.name == temp[0] {
				for j, sroom := range rooms {
					if sroom.name == temp[1] {
						rooms[i].neighbours = append(rooms[i].neighbours, &rooms[j])
						break
					}
				}
				break
			}

		}
		return true
	}

	return false

}

func isRoom(line string) bool {
	temp := strings.Split(line, " ")
	if len(temp) != 3 {
		return false
	}
	_, err1 := strconv.Atoi(temp[1])
	_, err2 := strconv.Atoi(temp[2])
	if err1 != nil || err2 != nil {
		return false
	}
	if temp[0][0] == '#' || temp[0][0] == 'L' || temp[0][0] == ' ' {
		return false
	}
	listOfRooms = append(listOfRooms, temp[0])

	return true
}
