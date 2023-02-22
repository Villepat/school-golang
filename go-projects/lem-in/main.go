package main

import l "lem-in/solve"

func main() {
	//Parse input map into usable variables
	l.ParseInput()
	//Go over all the rooms recursively to find all the possivle paths
	l.DFS(l.StartRoom, l.EndRoom)
	//Find non intersecting path combinations
	l.FNIPC(l.SimplePaths)
	//Fastest path combination
	FPC := l.FindFastest(l.NIPC)
	//Assign path to each ant
	l.DistributeAnts(FPC, l.Nr)
	//Move ants and print moves
	l.MandP(l.FastPaths, l.Nr)
}
