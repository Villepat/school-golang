package lemin

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	name       string
	neighbours []*room
	visited    bool
	x          string
	y          string
}

var listOfRooms []string
var rooms []room
var StartRoom *room
var EndRoom *room
var antNumbersFound bool

type fastPath struct {
	id          int
	length      int
	travel_time int
	ants        []ant
	path        []string
}

type ant struct {
	num  int
	path []string
}

var FastPaths []fastPath
var ants []ant

// Returns the sum of length of all arrays in float64
func lenAll(combo [][]string) float64 {
	var l float64
	l = 0
	for _, path := range combo {
		l += float64(len(path))
	}
	return l
}

// Equal tells whether a and b contain the same elements.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Assign a path to each ant in the colony
func DistributeAnts(FPC [][]string, Nr int) {
	for i := range FPC {
		FPC[i] = append(FPC[i], EndRoom.name)
		temp := []string{StartRoom.name}
		FPC[i] = append(temp, FPC[i]...)
	}
	for i := 1; i <= Nr; i++ {
		var a ant
		a.num = i
		ants = append(ants, a)
	}
	for i, p := range FPC {
		var fp fastPath
		fp.id = i
		fp.path = p
		fp.length = len(p)
		fp.travel_time = fp.length + len(fp.ants)
		FastPaths = append(FastPaths, fp)
	}
	for _, a := range ants {
		shortest := FastPaths[0]
		for _, p := range FastPaths {
			if p.length+len(p.ants) < shortest.length+len(shortest.ants) {
				shortest = p
			}
		}
		for i, p := range FastPaths {
			if p.id == shortest.id {
				FastPaths[i].ants = append(FastPaths[i].ants, a)
			}
		}
	}
}

var currentPath []string
var SimplePaths [][]string

// DFS is a recursive function that uses the depth-first search algorithm to
// find all paths from the start room to the end room
func DFS(current *room, end *room) {
	if current.visited {
		return
	}
	current.visited = true

	currentPath = append(currentPath, current.name)
	if current.name == end.name {
		temp := make([]string, len(currentPath))
		copy(temp, currentPath)
		SimplePaths = append(SimplePaths, temp)
		current.visited = false
		currentPath = currentPath[:(len(currentPath) - 1)]
		return
	}
	for _, adj := range current.neighbours {
		DFS(adj, end)
	}
	currentPath = currentPath[:(len(currentPath) - 1)]
	current.visited = false

}

// Non intersecting path combinations
var NIPC [][][]string

var sortedNIPC = make(map[float64][][]string)

// Find fastest combination of paths by checking the travel time of each path combination and comparing them
func FindFastest(NIPC [][][]string) [][]string {
	var average float64
	// this loop gets average length of all the path combinations and calculates the travel time of each path combination
	for _, combo := range NIPC {

		average = lenAll(combo) / float64(len(combo))
		average = average / float64(len(combo))
		sortedNIPC[average] = combo
	}

	averages := make([]float64, 0, len(sortedNIPC))
	for k := range sortedNIPC {
		averages = append(averages, k)
	}
	sort.Float64s(averages)
	fastest := sortedNIPC[averages[0]]
	return fastest
}

var preNIPC [][]string

// Find non-intersecting path combinations
func FNIPC(SimplePaths [][]string) {
	for i := range SimplePaths {
		SimplePaths[i] = SimplePaths[i][1 : len(SimplePaths[i])-1]
	}

	if len(SimplePaths) == 0 {
		fmt.Println("ERROR: invalid data format")
		exitCode := 0
		defer func() { os.Exit(exitCode) }()
	}

	for i := range SimplePaths {
		preNIPC = append(preNIPC, SimplePaths[i])

		for j := range SimplePaths {
			if i == j {
				continue
			}

			if !areIntersected(preNIPC, SimplePaths[j]) {
				preNIPC = append(preNIPC, SimplePaths[j])
			}
		}
		NIPC = append(NIPC, preNIPC)

		preNIPC = [][]string{}
	}
}

// Checks if two paths interset along the way
func areIntersected(paths [][]string, other_path []string) bool {
	for i := range paths {
		for j := range paths[i] {
			for k := range other_path {
				if paths[i][j] == other_path[k] {
					return true
				}
			}
		}
	}
	return false
}

// Moves the ants and prints their moves
func MandP(fps []fastPath, Nr int) {
	//remove duplicates in fps
	for fp := range fps {
		for i := 0; i < len(fps[fp].path)-1; i++ {
			if fps[fp].path[i] == fps[fp].path[i+1] {
				fps[fp].path = append(fps[fp].path[:i], fps[fp].path[i+1:]...)
			}
		}
	}
	for j, fp := range fps {
		for i := range fp.ants {
			fps[j].ants[i].path = fps[j].path[1:]
		}
	}
	moves := []string{}

	arrived := 0
	steps := 0
	for arrived < Nr {
		for _, fp := range fps {
			for i, a := range fp.ants {
				ok := true
				for _, m := range moves {
					if len(fp.ants[i].path) > 0 {
						if fp.ants[i].path[0] == m && m != EndRoom.name {
							ok = false
						}
					}

				}
				if ok {
					if len(fp.ants[i].path) > 0 {
						fmt.Printf("L%d-%s ", a.num, fp.ants[i].path[0])
						if fp.ants[i].path[0] == EndRoom.name {
							arrived++
						}
						moves = append(moves, fp.ants[i].path[0])
						fp.ants[i].path = fp.ants[i].path[1:]
					}
				}
			}
		}
		fmt.Println()
		moves = []string{}
		steps++
	}
	// fmt.Println(steps) if you want to see the number of steps
}

var Nr = 0

// Reads the data from the file passed as argument
func ParseInput() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run lem-in.go <filename.txt>")
		exitCode := 0
		defer func() { os.Exit(exitCode) }() // this exits without os.Exit status print which is required in this particular subject
	} else {
		reader, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		farm := string(reader)
		lines := strings.Split(farm, "\n")
		for i := 0; i < len(lines); i++ {
			if lines[i] == "" {
				if i < len(lines)-1 {
					lines = append(lines[:i], lines[i+1:]...)
					i--
				} else {
					lines = lines[:i]
					i--
				}
			}
		}
		start := ""
		end := ""

		// Parsing the text file into a graph (network of nodes)
		for i, line := range lines {
			switch {
			case line == "##start":

				// Find next valid room
				for j := i + 1; j < len(lines); j++ {
					if isRoom(lines[j]) {
						start = strings.Split(lines[j], " ")[0]
						break
					}
				}
				if start == "" {
					fmt.Println("ERROR: invalid data format")
					exitCode := 0
					defer func() { os.Exit(exitCode) }()
				}
			case line == "##end":
				// Find next valid room
				if i < len(lines)-3 { // end can not be at last line, must be X amout of links after end
					for j := i + 1; j < len(lines); j++ {
						if isRoom(lines[j]) {
							end = strings.Split(lines[j], " ")[0]
							break
						}
					}
					if end == "" {
						fmt.Println("ERROR: invalid data format")
						exitCode := 0
						defer func() { os.Exit(exitCode) }()
					}
				} else {
					fmt.Println("Invalid map format") // call error function?
					return
				}
			case isRoom(line):
				// add room to graph
				// also validate temp for correct room format? no invalid characters etc
				temp := strings.Split(line, " ")
				var r room
				if temp[0] != " " { //instead check if 1st element is letter(s) (except L which is reserved for ants) or numbers. check if 2nd and third element is numbers.
					r.name = temp[0]
					r.x = temp[1]
					r.y = temp[2]
					r.visited = false
					rooms = append(rooms, r)

				} else {
					//big error!
					fmt.Println("incorrect map format")
					exitCode := 0
					defer func() { os.Exit(exitCode) }()
				}

			// add connection between nodes
			case isLink(line):
				continue

			case isNrOfAnts(line):
				Nr, err = strconv.Atoi(line)

				if Nr < 1 { // if ant nr is 0 or less
					fmt.Println("ERROR: invalid data format")
					exitCode := 0
					defer func() { os.Exit(exitCode) }()

				}

				if antNumbersFound { // if there are more than 2 antnumber lines in the file -> error

					fmt.Println("ERROR: invalid data format")
					exitCode := 0
					defer func() { os.Exit(exitCode) }()
				}
				if err != nil {
					fmt.Println("Error fetching number of ants.")
					exitCode := 0
					defer func() { os.Exit(exitCode) }()
				}
				antNumbersFound = true

			default:
				continue
			}

		}
		findStart(start)
		findEnd(end)
		if !findEnd(end) || !findStart(start) || start == end {
			fmt.Println("ERROR: invalid data format")
			exitCode := 0
			defer func() { os.Exit(exitCode) }()
		}
	}
	for i, r := range rooms {
		for j, r2 := range rooms {
			if i == j {
				continue
			}
			if r.x == r2.x && r.y == r2.y { // if rooms have same coordinates
				fmt.Println("ERROR: invalid data format")
				exitCode := 0
				defer func() { os.Exit(exitCode) }()
			}
		}
	}
	for i, r := range rooms {
		for j, r2 := range rooms {
			if i == j {
				continue
			}
			if r.name == r2.name {
				fmt.Println("ERROR: invalid data format")
				exitCode := 0
				defer func() { os.Exit(exitCode) }()
			}
		}
	}
}
